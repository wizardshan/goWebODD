package user

import (
	"github.com/tidwall/gjson"
)

const (
	LevelNormal   = 0
	LevelSilver   = 10
	LevelGold     = 20
	LevelPlatinum = 30
)

type Level struct {
	Value int64 `binding:"oneof=0 10 20 30"`
	Set   bool
	Mask  bool
	Desc  string
}

func NewLevel(v int64) Level {
	o := Level{Value: v, Set: true}
	o.Desc = o.GetDesc()
	return o
}

func (obj Level) IsPresent(f func(o Level)) {
	if obj.Set {
		f(obj)
	}
}

func (obj Level) IsMask(f func(o Level)) {
	if obj.Mask {
		f(obj)
	}
}

func (obj Level) GetDesc() string {
	desc, ok := obj.DescMapping()[obj.Value]
	if ok {
		return desc
	}
	return "未知"
}

func (obj Level) DescMapping() map[int64]string {
	return map[int64]string{
		LevelNormal:   "普通",
		LevelSilver:   "白银",
		LevelGold:     "黄金",
		LevelPlatinum: "铂金",
	}
}

func (obj *Level) UnmarshalJSON(data []byte) error {
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
