package post

import (
	"github.com/tidwall/gjson"
)

type Content struct {
	Value string `binding:"min=6,max=30"`
	Set   bool
	Mask  bool
}

func NewContent(v string) Content {
	return Content{Value: v, Set: true}
}

func (obj Content) IsPresent(f func(o Content)) {
	if obj.Set {
		f(obj)
	}
}

func (obj Content) IsMask(f func(o Content)) {
	if obj.Mask {
		f(obj)
	}
}

func (obj *Content) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		obj.Value = gjson.ParseBytes(data).String()
		obj.Set = true
		return nil
	}

	results := gjson.GetManyBytes(data, "Value", "Mask")
	if results[0].Exists() {
		obj.Value = results[0].String()
		obj.Set = true
	}
	if results[1].Exists() {
		obj.Mask = results[1].Bool()
	}
	return nil
}
