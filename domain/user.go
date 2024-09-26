package domain

import (
	"goWebODD/controller/response"
	"goWebODD/domain/user"
)

type Users []*User

type User struct {
	ID     ID
	Mobile Mobile
	Age    user.Age
	Level  user.Level
	//HashID   user.HashID
	//Password user.Password
}

func (user *User) Mapper() *response.User {
	if user == nil {
		return nil
	}
	return &response.User{
		ID:        user.ID.Value,
		Mobile:    user.Mobile.Value,
		Age:       user.Age.Value,
		Level:     user.Level.Value,
		LevelDesc: user.Level.Desc,
	}
}

func (users Users) Mapper() response.Users {
	size := len(users)
	respUsers := make(response.Users, size)
	for i := 0; i < size; i++ {
		respUsers[i] = users[i].Mapper()
	}
	return respUsers
}
