package schema

import (
	"encoding/json"
	"github.com/MinseokOh/data-warehouse/types"
	"github.com/jf-tech/omniparser"
	"github.com/jf-tech/omniparser/transformctx"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// TestSchema
func RunTestSchema(t *testing.T, schemaFile, inputFile string, object interface{}) {
	schemaReader, err := os.Open(schemaFile)
	assert.NoError(t, err)

	inputReader, err := os.Open(inputFile)
	assert.NoError(t, err)

	defer func() {
		err := schemaReader.Close()
		assert.NoError(t, err)

		err = inputReader.Close()
		assert.NoError(t, err)
	}()

	schema, err := omniparser.NewSchema(schemaFile, schemaReader)
	assert.NoError(t, err)

	ctx := &transformctx.Ctx{}
	transform, err := schema.NewTransform(inputFile, inputReader, ctx)
	assert.NoError(t, err)

	b, err := transform.Read()
	assert.NoError(t, err)

	err = json.Unmarshal(b, &object)
	assert.NoError(t, err)
}

func TestJson1(t *testing.T) {
	var Sdk types.Sdk
	RunTestSchema(t, "./json_1_schema.json", "./test/json_1_test.json", &Sdk)

	assert.Equal(t, Sdk.Id, "1")
	assert.Equal(t, Sdk.Name, "golang")
}

func TestJson2(t *testing.T) {
	var Sdk types.Sdk
	RunTestSchema(t, "./json_2_schema.json", "./test/json_2_test.json", &Sdk)

	assert.Equal(t, Sdk.Id, "12")
	assert.Equal(t, Sdk.Name, "golang")
	assert.Equal(t, Sdk.Version, "100,000")
}

func TestXml1(t *testing.T) {
	var Sdk types.Sdk
	RunTestSchema(t, "./xml_1_schema.json", "./test/xml_1_test.xml", &Sdk)

	assert.Equal(t, Sdk.Id, "165")
	assert.Equal(t, Sdk.Name, "golang")
	assert.Equal(t, Sdk.Version, "15")
}

func TestXml2(t *testing.T) {
	var Sdk types.Sdk
	RunTestSchema(t, "./xml_2_schema.json", "./test/xml_2_test.xml", &Sdk)

	assert.Equal(t, Sdk.Id, "177")
	assert.Equal(t, Sdk.Name, "golang")
	assert.Equal(t, Sdk.Version, "11000")
}

func TestCsv1(t *testing.T) {
	var Sdk types.Sdk
	RunTestSchema(t, "./csv_1_schema.json", "./test/csv_1_test.csv", &Sdk)

	assert.Equal(t, Sdk.Id, "go-lang")
	assert.Equal(t, Sdk.Name, "version")
	assert.Equal(t, Sdk.Version, "1000")
}
