package domain

import (
	"github.com/tidwall/gjson"
)

type ID struct {
	Value int64 `binding:"min=1"`
	Set   bool
	Mask  bool
}

func NewID(v int64) ID {
	return ID{Value: v, Set: true}
}

func (obj ID) IsPresent(f func(o ID)) {
	if obj.Set {
		f(obj)
	}
}

func (obj ID) IsMask(f func(o ID)) {
	if obj.Mask {
		f(obj)
	}
}

func (obj *ID) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		obj.Value = gjson.ParseBytes(data).Int()
		obj.Set = true
		return nil
	}

	results := gjson.GetManyBytes(data, "Value", "Mask")
	if results[0].Exists() {
		obj.Value = results[0].Int()
		obj.Set = true
	}
	if results[1].Exists() {
		obj.Mask = results[1].Bool()
	}
	return nil
}
