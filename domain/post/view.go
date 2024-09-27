package post

import (
	"github.com/tidwall/gjson"
)

type View struct {
	Value int64 `binding:"min=1,max=150"`
	Set   bool  `binding:"required"`
	Mask  bool
}

func NewView(v int64) View {
	return View{Value: v, Set: true}
}

func (o *View) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.Value = gjson.ParseBytes(data).Int()
		o.Set = true
		return nil
	}

	results := gjson.GetManyBytes(data, "Value", "Mask")
	if results[0].Exists() {
		o.Value = results[0].Int()
		o.Set = true
	}
	if results[1].Exists() {
		o.Mask = results[1].Bool()
	}
	return nil
}

func (o *View) Name() string {
	return "View"
}
