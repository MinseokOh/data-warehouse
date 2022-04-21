package xml2

import (
	"encoding/xml"
	"github.com/MinseokOh/data-warehouse/util"
	"math/rand"
)

type Sdk struct {
	SdkID      string `xml:"sdkID"`
	SdkName    string `xml:"sdkName"`
	SdkVersion int    `xml:"sdkVersion"`
}

func (p Sdk) GenRandom() ([]byte, error) {
	p.SdkID = util.RandString(rand.Intn(20))
	p.SdkName = util.RandString(rand.Intn(20))
	p.SdkVersion = rand.Intn(10000000)
	return xml.Marshal(p)
}
