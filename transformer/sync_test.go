package transformer

import (
	"encoding/json"
	"github.com/MinseokOh/data-warehouse/producer"
	"github.com/MinseokOh/data-warehouse/types"
	"github.com/MinseokOh/data-warehouse/types/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSyncGetQuery(t *testing.T) {
	sync := &Sync{}
	schemas := schema.NewSchemas(
		"Sdk",
		"../schema/json_1_schema.json",
	)

	var products []*types.Sdk
	prod := producer.NewProducer("../configs/producer_1_config.json")

	for i := 0; i < 10000; i++ {
		b, err := prod.GenRandom()
		assert.NoError(t, err)

		b, err = schemas.Transform(b)
		var Sdk types.Sdk
		err = json.Unmarshal(b, &Sdk)
		assert.NoError(t, err)

		products = append(products, &Sdk)
	}

	query := sync.GetQuery(products)
	assert.NotEqualValues(t, query, "")
}
