package transformer

import (
	"github.com/MinseokOh/data-warehouse/types/db"
	"github.com/MinseokOh/data-warehouse/util/log"
)

type Config struct {
	Kafka struct {
		Brokers []string `json:"brokers"`
		GroupId string   `json:"group_id"`
		Topics  []string `json:"topics"`
	} `json:"kafka"`
	Schema struct {
		Name  string   `json:"name"`
		Files []string `json:"files"`
	} `json:"schema"`
	Logger log.Config `json:"log"`
	DB     db.Config  `json:"db"`
}
