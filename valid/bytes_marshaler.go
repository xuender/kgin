package valid

import (
	"encoding/json"

	"github.com/samber/lo"
)

type BytesMarshaler []byte

func NewBytesMarshaler(marshaler json.Marshaler) BytesMarshaler {
	data := lo.Must1(marshaler.MarshalJSON())

	return BytesMarshaler(data)
}

func (p BytesMarshaler) MarshalJSON() ([]byte, error) {
	return p, nil
}
