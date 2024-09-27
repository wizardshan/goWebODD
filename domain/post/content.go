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

func (o *Content) UnmarshalJSON(data []byte) error {
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

func (o *Content) Name() string {
	return "Content"
}
