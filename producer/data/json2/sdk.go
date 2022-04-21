package json2

import (
	"encoding/json"
	"github.com/MinseokOh/data-warehouse/util"
	"math/rand"
)

type Sdk struct {
	ID   string `json:"ID"`
	Name string `json:"name"`
	Ver  int    `json:"ver"`
}

func (p Sdk) GenRandom() ([]byte, error) {
	p.ID = util.RandString(rand.Intn(20))
	p.Name = util.RandString(rand.Intn(20))
	p.Ver = rand.Intn(10000000)
	return json.Marshal(p)
}
