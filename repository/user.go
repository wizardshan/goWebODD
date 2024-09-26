package repository

import (
	"context"
	"goWebODD/domain"
	"goWebODD/repository/ent"
	"goWebODD/repository/ent/user"
)

type User struct {
	db *ent.Client
}

func NewUser(db *ent.Client) *User {
	repo := new(User)
	repo.db = db
	return repo
}

func (repo *User) FetchByID(ctx context.Context, id int64) *domain.User {
	return repo.db.User.Query().Where(user.ID(id)).FirstX(ctx).Mapper()
}

func (repo *User) FetchByHashID(ctx context.Context, hashID string) *domain.User {
	return repo.db.User.Query().Where(user.HashID(hashID)).FirstX(ctx).Mapper()
}

func (repo *User) FetchMany(ctx context.Context, domUser *domain.User) domain.Users {
	query := repo.db.User.Query()
	if domUser.Level.Set {
		query.Where(user.Level(domUser.Level.Value))
	}
	if domUser.Age.Set {
		query.Where(user.Age(domUser.Age.Value))
	}
	if domUser.Mobile.Set {
		query.Where(user.Mobile(domUser.Mobile.Value))
	}
	var entUsers ent.Users = query.AllX(ctx)
	return entUsers.Mapper()
}
