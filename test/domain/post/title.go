package post

import (
	"github.com/tidwall/gjson"
)

type Title struct {
	Value string `binding:"min=6,max=30"`
	Set   bool
	Mask  bool
}

func NewTitle(v string) Title {
	return Title{Value: v, Set: true}
}

func (obj Title) IsPresent(f func(o Title)) {
	if obj.Set {
		f(obj)
	}
}

func (obj Title) IsMask(f func(o Title)) {
	if obj.Mask {
		f(obj)
	}
}

func (o *Title) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.Value = gjson.ParseBytes(data).String()
		o.Set = true
		return nil
	}

	results := gjson.GetManyBytes(data, "Value", "Mask")
	if results[0].Exists() {
		o.Value = results[0].String()
		o.Set = true
	}

	if results[1].Exists() {
		o.Mask = results[1].Bool()
	}
	return nil
}
