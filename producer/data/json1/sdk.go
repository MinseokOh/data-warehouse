package json1

import (
	"encoding/json"
	"github.com/MinseokOh/data-warehouse/util"
	"math/rand"
)

type Sdk struct {
	SdkNo   string `json:"sdkNo"`
	SdkName string `json:"sdkName"`
}

func (p Sdk) GenRandom() ([]byte, error) {
	p.SdkNo = util.RandString(rand.Intn(20))
	p.SdkName = util.RandString(rand.Intn(20))
	return json.Marshal(p)
}
