package data

type Type interface {
	GenRandom() ([]byte, error)
}
