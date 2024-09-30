package ent

import (
	domain2 "goWebODD/test/domain"
	user2 "goWebODD/test/domain/user"
)

func (entUser *User) Mapper() domain2.User {
	var domUser domain2.User
	if entUser == nil {
		return domUser
	}
	domUser.ID = domain2.NewID(entUser.ID)
	domUser.Mobile = domain2.NewMobile(entUser.Mobile)
	domUser.Age = user2.NewAge(entUser.Age)
	domUser.Level = user2.NewLevel(entUser.Level)
	return domUser
}

func (entUsers Users) Mapper() domain2.Users {
	size := len(entUsers)
	domUsers := make(domain2.Users, size)
	for i := 0; i < size; i++ {
		domUsers[i] = entUsers[i].Mapper()
	}
	return domUsers
}
