package producer

import (
	"encoding/json"
	"fmt"
	"github.com/MinseokOh/data-warehouse/producer/data"
	"github.com/MinseokOh/data-warehouse/producer/data/csv"
	"github.com/MinseokOh/data-warehouse/producer/data/json1"
	"github.com/MinseokOh/data-warehouse/producer/data/json2"
	"github.com/MinseokOh/data-warehouse/producer/data/xml1"
	"github.com/MinseokOh/data-warehouse/producer/data/xml2"
	"github.com/MinseokOh/data-warehouse/util"
	"github.com/MinseokOh/data-warehouse/util/log"
	"github.com/Shopify/sarama"
	"os"
	"time"
)

type Producer struct {
	producer      sarama.SyncProducer
	producedBytes int64
	duration      time.Duration
	config        Config
	logger        *log.Logger
	dataType      data.Type
}

func NewProducer(config string) *Producer {
	var producer Producer
	producer.loadConfig(config)

	switch producer.config.Type {
	case "xml1":
		producer.dataType = xml1.Sdk{}
		break
	case "xml2":
		producer.dataType = xml2.Sdk{}
		break
	case "json1":
		producer.dataType = json1.Sdk{}
		break
	case "json2":
		producer.dataType = json2.Sdk{}
		break
	case "csv":
		producer.dataType = csv.Sdk{}
		break
	default:
		panic("invalid data type, support (xml1|xml2|json1|json2|csv)")
	}
	producer.producedBytes = 0

	var err error
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Partitioner = sarama.NewRandomPartitioner
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Return.Successes = true
	producer.producer, err = sarama.NewSyncProducer(producer.config.Kafka.Brokers, saramaConfig)

	if err != nil {
		panic(err)
	}
	producer.logger = log.NewLoggerConfig(fmt.Sprintf("producer-%s", producer.config.Type), producer.config.Logger)

	return &producer
}

func (p *Producer) loadConfig(config string) {
	f, err := os.Open(config)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&p.config)
	if err != nil {
		panic(err)
	}
}

func (p Producer) Run() {
	go p.runProducer()

	for range time.Tick(time.Second * 5) {
		p.logger.Info("produced:", util.ByteCountSI(p.producedBytes))
	}
}

func (p Producer) GenRandom() ([]byte, error) {
	return p.dataType.GenRandom()
}

func (p *Producer) runProducer() {
	p.logger.Info("run producers")

	for range time.Tick(time.Nanosecond * time.Duration(p.config.Duration)) {
		bytes, err := p.GenRandom()
		if err != nil {
			p.logger.Error(err)
			panic(err)
		}

		p.logger.Debug(string(bytes))

		_, _, err = p.producer.SendMessage(&sarama.ProducerMessage{
			Topic:     p.config.Kafka.Topic,
			Partition: int32(p.config.Kafka.Partition),
			Value:     sarama.ByteEncoder(bytes),
		})

		if err != nil {
			p.logger.Error(err)
			panic(err)
		}

		p.producedBytes += int64(len(bytes))
	}
}
