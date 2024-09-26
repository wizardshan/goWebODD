package domain

import (
	"encoding/json"
)

type ID struct {
	Value int64 `binding:"min=1"`
	Set   bool
}

func NewID(v int64) ID {
	return ID{Value: v, Set: true}
}

func (o *ID) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &o.Value); err != nil {
		return err
	}
	o.Set = true
	return nil
}
