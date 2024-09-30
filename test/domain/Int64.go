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
	if data[0] == '{' {
		results := gjson.GetManyBytes(data, "Value", "Mask")
		if results[0].Exists() {
			o.Value = results[0].Int()
			o.Set = true
		}
		if results[1].Exists() {
			o.Mask = true
		}
		return nil
	}
	if err := json.Unmarshal(data, &o.Value); err == nil {
		o.Set = true
		return nil
	}

	return nil
}

func (o Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Value)
}
