package user

import (
	"github.com/tidwall/gjson"
)

type Age struct {
	Value int64 `binding:"min=1,max=150"`
	Set   bool
	Mask  bool
}

func NewAge(v int64) Age {
	return Age{Value: v, Set: true}
}

func (obj Age) IsPresent(f func(o Age)) {
	if obj.Set {
		f(obj)
	}
}

func (obj Age) IsMask(f func(o Age)) {
	if obj.Mask {
		f(obj)
	}
}

func (obj *Age) UnmarshalJSON(data []byte) error {
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
