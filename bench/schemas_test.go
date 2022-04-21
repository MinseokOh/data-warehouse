package bench

import (
	"encoding/json"
	"github.com/MinseokOh/data-warehouse/producer/data/xml1"
	"github.com/MinseokOh/data-warehouse/producer/data/xml2"
	"github.com/MinseokOh/data-warehouse/types"
	"github.com/MinseokOh/data-warehouse/types/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkSchemas(b *testing.B) {
	schemas := schema.NewSchemas("bench",
		"../schema/json_1_schema.json",
		"../schema/json_2_schema.json",
		"../schema/csv_1_schema.json",
		"../schema/xml_1_schema.json",
		"../schema/xml_2_schema.json",
	)

	dataType := xml2.Sdk{}
	r, err := dataType.GenRandom()
	assert.NoError(b, err)

	var Sdk types.Sdk

	for i := 0; i < b.N; i++ {
		transformed, err := schemas.Transform(r)
		assert.NoError(b, err)
		err = json.Unmarshal(transformed, &Sdk)
		assert.NoError(b, err)
	}
}

func Benchmark20Schemas(b *testing.B) {
	schemas := schema.NewSchemas("bench",
		"../schema/json_1_schema.json",
		"../schema/json_1_schema.json",
		"../schema/json_1_schema.json",
		"../schema/json_1_schema.json",
		"../schema/json_1_schema.json",
		"../schema/json_1_schema.json",
		"../schema/json_1_schema.json",
		"../schema/json_1_schema.json",
		"../schema/json_1_schema.json",
		"../schema/json_1_schema.json",
		"../schema/json_1_schema.json",
		"../schema/json_1_schema.json",
		"../schema/json_1_schema.json",
		"../schema/json_1_schema.json",
		"../schema/json_1_schema.json",
		"../schema/json_1_schema.json",
		"../schema/json_1_schema.json",
		"../schema/json_1_schema.json",
		"../schema/json_1_schema.json",
		"../schema/xml_1_schema.json",
	)

	dataType := xml1.Sdk{}
	r, err := dataType.GenRandom()
	assert.NoError(b, err)

	assert.NoError(b, err)

	var Sdk types.Sdk

	for i := 0; i < b.N; i++ {
		transformed, err := schemas.Transform(r)
		assert.NoError(b, err)
		err = json.Unmarshal(transformed, &Sdk)
		assert.NoError(b, err)
	}
}
