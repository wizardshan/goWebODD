package user

import "encoding/json"

type Age struct {
	Value int  `binding:"min=1,max=150"`
	Set   bool `binding:"required"`
}

func NewAge(v int) Age {
	return Age{Value: v, Set: true}
}

func (o *Age) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &o.Value); err != nil {
		return err
	}
	o.Set = true
	return nil
}
