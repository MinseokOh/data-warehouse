package producer

import "github.com/MinseokOh/data-warehouse/util/log"

type Config struct {
	Type     string `json:"type"`
	Duration int    `json:"duration"`
	Kafka    struct {
		Brokers   []string `json:"brokers"`
		Topic     string   `json:"topic"`
		Partition int      `json:"partition"`
	} `json:"kafka"`
	Logger log.Config `json:"log"`
}
