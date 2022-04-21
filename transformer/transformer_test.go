package transformer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTransformer(t *testing.T) {
	assert.Panics(t, func() {
		NewTransformer("../configs/producer_1_config.json")
	})

	//assert.NotPanics(t, func() {
	//	NewTransformer("../configs/transformer_config.json")
	//})
}