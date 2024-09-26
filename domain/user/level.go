package user

import "encoding/json"

const (
	LevelNormal   = 0
	LevelSilver   = 10
	LevelGold     = 20
	LevelPlatinum = 30
)

type Level struct {
	Value int  `binding:"oneof=0 10 20 30"`
	Set   bool `binding:"required"`
	Desc  string
}

func NewLevel(v int) Level {
	o := Level{Value: v, Set: true}
	o.Desc = o.GetDesc()
	return o
}

func (o Level) GetDesc() string {
	desc, ok := o.DescMapping()[o.Value]
	if ok {
		return desc
	}
	return "未知"
}

func (o Level) DescMapping() map[int]string {
	return map[int]string{
		LevelNormal:   "普通",
		LevelSilver:   "白银",
		LevelGold:     "黄金",
		LevelPlatinum: "铂金",
	}
}

func (o *Level) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &o.Value); err != nil {
		return err
	}
	o.Set = true
	o.Desc = o.GetDesc()
	return nil
}
