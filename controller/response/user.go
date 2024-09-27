package response

type Users []*User

type User struct {
	ID        *int64
	Mobile    *string
	Age       *int64
	Level     *int64
	LevelDesc *string
}
