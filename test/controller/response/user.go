package response

type Users []*User

type User struct {
	ID        *int64  `json:",omitempty"`
	Mobile    *string `json:",omitempty"`
	Age       *int64  `json:",omitempty"`
	Level     *int64  `json:",omitempty"`
	LevelDesc *string `json:",omitempty"`
}
