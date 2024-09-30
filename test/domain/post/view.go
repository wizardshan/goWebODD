package post

import (
	"github.com/tidwall/gjson"
)

type View struct {
	Value int64 `binding:"min=1,max=150"`
	Set   bool
	Mask  bool
}

func NewView(v int64) View {
	return View{Value: v, Set: true}
}

func (obj View) IsPresent(f func(o View)) {
	if obj.Set {
		f(obj)
	}
}

func (obj View) IsMask(f func(o View)) {
	if obj.Mask {
		f(obj)
	}
}

func (obj *View) UnmarshalJSON(data []byte) error {
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
