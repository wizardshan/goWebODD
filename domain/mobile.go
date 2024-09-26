package domain

import (
	"encoding/json"
	"regexp"
)

type Mobile struct {
	Value string `binding:"mobile"`
	Set   bool
}

func NewMobile(v string) Mobile {
	return Mobile{Value: v, Set: true}
}

func (o Mobile) IsValid() bool {
	return regexp.MustCompile(`^(1[3-9][0-9]\d{8})$`).MatchString(o.Value)
}

func (o *Mobile) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &o.Value); err != nil {
		return err
	}
	o.Set = true
	return nil
}
