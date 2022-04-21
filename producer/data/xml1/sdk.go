package xml1

import (
	"encoding/xml"
	"github.com/MinseokOh/data-warehouse/util"
	"math/rand"
)

type Sdk struct {
	Id   string `xml:"id"`
	Name string `xml:"name"`
	V    int    `xml:"v"`
}

func (p Sdk) GenRandom() ([]byte, error) {
	p.Id = util.RandString(rand.Intn(20))
	p.Name = util.RandString(rand.Intn(20))
	p.V = rand.Intn(10000000)
	return xml.Marshal(p)
}
