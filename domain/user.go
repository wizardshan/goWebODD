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

func (domUser *User) Name() string {
	return "User"
}

func (domUser *User) Mapper() *response.User {
	if domUser == nil {
		return nil
	}
	return &response.User{
		ID:        &domUser.ID.Value,
		Mobile:    &domUser.Mobile.Value,
		Age:       &domUser.Age.Value,
		Level:     &domUser.Level.Value,
		LevelDesc: &domUser.Level.Desc,
	}
}

func (domUsers Users) Mapper() response.Users {
	size := len(domUsers)
	respUsers := make(response.Users, size)
	for i := 0; i < size; i++ {
		respUsers[i] = domUsers[i].Mapper()
	}
	return respUsers
}
