package schema

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSchemas(t *testing.T) {
	assert.NotPanics(t, func() {
		NewSchemas(
			"test",
			"../../schema/json_1_schema.json",
			"../../schema/json_2_schema.json",
			"../../schema/xml_1_schema.json",
			"../../schema/xml_2_schema.json",
		)
	})
}
