package transformer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MinseokOh/data-warehouse/types"
	"github.com/MinseokOh/data-warehouse/types/schema"
	"github.com/MinseokOh/data-warehouse/util"
	"github.com/MinseokOh/data-warehouse/util/log"
	"github.com/Shopify/sarama"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"sync"
	"time"
)

const (
	KafkaTopic = "Sdk"
)

type Transformer struct {
	consumerGroup sarama.ConsumerGroup
	sync          *Sync
	config        Config
	schemas       *schema.Schemas
	logger        *log.Logger
	consumedCount int64
	consumedBytes int64
	ready         chan bool
	wg            *sync.WaitGroup
}

func newTransformer() *Transformer {
	return &Transformer{
		logger:        log.NewLogger("transformer"),
		consumedCount: 0,
		ready:         make(chan bool),
		wg:            &sync.WaitGroup{},
	}
}

func NewTransformer(config string) *Transformer {
	transformer := newTransformer()
	transformer.loadConfig(config)
	transformer.logger = log.NewLoggerConfig(
		"transformer",
		transformer.config.Logger,
	)
	transformer.schemas = schema.NewSchemas(
		transformer.config.Schema.Name,
		transformer.config.Schema.Files...,
	)
	transformer.sync = NewSync(transformer.config)

	var err error
	saramaConfig := sarama.NewConfig()
	saramaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	transformer.consumerGroup, err = sarama.NewConsumerGroup(
		transformer.config.Kafka.Brokers,
		transformer.config.Kafka.GroupId,
		saramaConfig,
	)

	if err != nil {
		panic(err)
	}
	return transformer
}

func (t *Transformer) loadConfig(config string) {
	f, err := os.Open(config)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&t.config)
	if err != nil {
		panic(err)
	}
}

func (t *Transformer) Run() {
	t.logger.Info("run transformer")

	go t.sync.Run()

	go func() {
		for {
			err := t.consumerGroup.Consume(context.Background(), t.config.Kafka.Topics, t)
			if err != nil {
				panic(err)
			}

			t.ready = make(chan bool)
		}
	}()
	<-t.ready
	t.logger.Info("transformer ready")

	for range time.Tick(time.Second * 5) {
		t.logger.Info(
			fmt.Sprintf("consumed : %d | synced\t: %d | size : %s",
				t.consumedCount,
				t.sync.GetSyncCount(),
				util.ByteCountSI(t.consumedBytes),
			),
		)
	}
}

func (t *Transformer) Setup(session sarama.ConsumerGroupSession) error {
	t.logger.Info("consumer group setup")
	for topic, partitions := range session.Claims() {
		for _, partition := range partitions {
			t.logger.Info(fmt.Sprintf("consume topic: %s, partition: %d", topic, partition))
		}
	}
	close(t.ready)
	return nil
}

func (t *Transformer) Cleanup(sarama.ConsumerGroupSession) error {
	t.logger.Info("consumer group cleanup")
	return nil
}

func (t *Transformer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	t.logger.Info("consume claim")

	for message := range claim.Messages() {
		transformed, err := t.schemas.Transform(message.Value)
		if err != nil {
			t.logger.Error(err)
			continue
		}

		var Sdk types.Sdk
		if err = json.Unmarshal(transformed, &Sdk); err != nil {
			t.logger.Error(err)
			continue
		}

		t.consumedCount += 1
		t.consumedBytes += int64(len(message.Value))

		t.logger.Debug(t.consumedCount, message.Offset, "r", string(message.Value))
		t.logger.Debug(t.consumedCount, message.Offset, "t", string(transformed))

		session.MarkMessage(message, "")

		t.sync.Insert(&Sdk)
	}
	return nil
}
