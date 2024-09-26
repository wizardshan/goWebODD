package response

type Users []*User

type User struct {
	ID        int64
	Mobile    string
	Age       int
	Level     int
	LevelDesc string
}
