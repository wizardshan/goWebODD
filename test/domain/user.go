package domain

import (
	"fmt"
	"goWebODD/test/controller/response"
	user2 "goWebODD/test/domain/user"
)

type Users []User

type User struct {
	ID     ID
	Mobile Mobile
	Age    user2.Age
	Level  user2.Level
	//HashID   user.HashID
	//Password user.Password

}

func (domUser User) IsSet() bool {
	var zero User
	return zero != domUser
}

func (domUser User) IsMask() bool {
	return domUser.ID.Mask ||
		domUser.Mobile.Mask ||
		domUser.Age.Mask ||
		domUser.Level.Mask
}

func (domUser User) Mapper(queryUser User) *response.User {
	if !domUser.IsSet() {
		return nil
	}

	if !queryUser.IsMask() {
		return &response.User{
			ID:        &domUser.ID.Value,
			Mobile:    &domUser.Mobile.Value,
			Age:       &domUser.Age.Value,
			Level:     &domUser.Level.Value,
			LevelDesc: &domUser.Level.Desc,
		}
	}

	respUser := &response.User{}
	if queryUser.ID.Mask {
		respUser.ID = &domUser.ID.Value
	}

	if queryUser.Mobile.Mask {
		respUser.Mobile = &domUser.Mobile.Value
	}

	if queryUser.Age.Mask {
		respUser.Age = &domUser.Age.Value
	}
	fmt.Println(queryUser.Level.Mask)
	if queryUser.Level.Mask {
		respUser.Level = &domUser.Level.Value
		respUser.LevelDesc = &domUser.Level.Desc
	}

	return respUser
}

func (domUsers Users) Mapper(queryUser User) response.Users {
	size := len(domUsers)
	respUsers := make(response.Users, size)
	for i := 0; i < size; i++ {
		respUsers[i] = domUsers[i].Mapper(queryUser)
	}
	return respUsers
}
