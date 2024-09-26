package ent

import (
	"goWebODD/domain"
	"goWebODD/domain/user"
)

func (entUser *User) Mapper() *domain.User {
	if entUser == nil {
		return nil
	}

	return &domain.User{
		ID:     domain.NewID(entUser.ID),
		Mobile: domain.NewMobile(entUser.Mobile),
		Age:    user.NewAge(entUser.Age),
		Level:  user.NewLevel(entUser.Level),
	}
}

func (entUsers Users) Mapper() domain.Users {
	size := len(entUsers)
	domUsers := make(domain.Users, size)
	for i := 0; i < size; i++ {
		domUsers[i] = entUsers[i].Mapper()
	}
	return domUsers
}
