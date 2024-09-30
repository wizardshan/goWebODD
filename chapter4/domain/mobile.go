package domain

import (
	"errors"
	"github.com/tidwall/gjson"
	"regexp"
)

type Mobile struct {
	Value string `binding:"number,mobile"`
	Set   bool
}

func NewMobile(v string) Mobile {
	return Mobile{Value: v, Set: true}
}

func (o Mobile) IsValid() bool {
	return regexp.MustCompile(`^(1[3-9][0-9]\d{8})$`).MatchString(o.Value)
}

func (o *Mobile) UnmarshalJSON(data []byte) error {
	result := gjson.GetBytes(data, "Value")
	if !result.Exists() {
		return errors.New("`Value` key not found")
	}
	o.Value = result.String()
	o.Set = true
	return nil
}
