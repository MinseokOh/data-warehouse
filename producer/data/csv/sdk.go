package csv

import (
	"fmt"
	"github.com/MinseokOh/data-warehouse/util"
	"math/rand"
)

type Sdk struct{}

func (p Sdk) GenRandom() ([]byte, error) {
	raw := fmt.Sprintf("%s %s %d", util.RandString(20), util.RandString(20), rand.Intn(10000000))
	return []byte(raw), nil
}
