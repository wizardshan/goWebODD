package repository

import (
	"context"
	"goWebODD/repository/ent"
	"goWebODD/repository/ent/post"
	"goWebODD/test/domain"
)

type Post struct {
	db *ent.Client
}

func NewPost(db *ent.Client) *Post {
	repo := new(Post)
	repo.db = db
	return repo
}

func (repo *Post) FetchByID(ctx context.Context, id int64) domain.Post {
	return repo.db.Post.Query().WithUser().Where(post.ID(id)).FirstX(ctx).Mapper()
}

func (repo *Post) FetchMany(ctx context.Context, domPost domain.Post) domain.Posts {
	var entPosts ent.Posts = repo.db.Post.Query().AllX(ctx)
	return entPosts.Mapper()
}
