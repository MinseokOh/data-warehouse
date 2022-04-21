package producer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProducer(t *testing.T) {
	assert.NotPanics(t, func() {
		NewProducer("../configs/producer_1_config.json")
	})

	assert.NotPanics(t, func() {
		NewProducer("../configs/producer_2_config.json")
	})

	assert.NotPanics(t, func() {
		NewProducer("../configs/producer_3_config.json")
	})

	assert.NotPanics(t, func() {
		NewProducer("../configs/producer_4_config.json")
	})

	assert.NotPanics(t, func() {
		NewProducer("../configs/producer_5_config.json")
	})
}
