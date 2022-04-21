package schema

import (
	"bytes"
	"fmt"
	"github.com/jf-tech/omniparser"
	"github.com/jf-tech/omniparser/transformctx"
	"os"
)

var (
	NULL = []byte{110, 117, 108, 108}
)

type Schemas struct {
	name    string
	schemas []omniparser.Schema
	ctx     *transformctx.Ctx
}

func NewSchemas(name string, schemas ...string) *Schemas {
	var s Schemas

	s.ctx = &transformctx.Ctx{}
	s.name = name
	for _, schema := range schemas {
		schemaReader, err := os.Open(schema)
		if err != nil {
			panic(err)
		}

		omniSchema, err := omniparser.NewSchema(schema, schemaReader)
		if err != nil {
			panic(err)
		}

		s.schemas = append(s.schemas, omniSchema)
	}

	return &s
}

func (schemas Schemas) Transform(bb []byte) ([]byte, error) {
	for _, schema := range schemas.schemas {
		transform, err := schema.NewTransform(schemas.name, bytes.NewReader(bb), schemas.ctx)
		if err != nil {
			continue
		}

		b, err := transform.Read()
		if err != nil || bytes.Equal(b, NULL) {
			continue
		}

		return b, err
	}

	return nil, fmt.Errorf("not supported schema")
}
