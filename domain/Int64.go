package domain

import (
	"encoding/json"
	"github.com/tidwall/gjson"
)

type Int64 struct {
	Value int64
	Set   bool
	Mask  bool
}

func (o Int64) IsSet() bool { return o.Set }

func (o *Int64) Reset() {
	var v int64
	o.Value = v
	o.Set = false
}

func (o *Int64) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

func (o Int64) Get() (v int64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

func (o Int64) Or(d int64) int64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

func (o *Int64) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &o.Value); err == nil {
		o.Set = true
		return nil
	}

	results := gjson.GetManyBytes(data, "Value", "Set")
	o.Value = results[0].Int()
	o.Set = results[1].Bool()
	return nil
}

func (o Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Value)
}
