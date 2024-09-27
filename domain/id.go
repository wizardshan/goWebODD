package domain

import (
	"github.com/tidwall/gjson"
)

type ID struct {
	Value int64 `binding:"min=1"`
	Set   bool
}

func NewID(v int64) ID {
	return ID{Value: v, Set: true}
}

func (o *ID) UnmarshalJSON(data []byte) error {
	o.Value = gjson.ParseBytes(data).Int()
	o.Set = true
	return nil
}

func (o *ID) Name() string {
	return "ID"
}
