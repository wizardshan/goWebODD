package ent

import (
	"goWebODD/domain"
	"goWebODD/domain/post"
)

func (entPost *Post) Mapper() *domain.Post {
	if entPost == nil {
		return nil
	}

	return &domain.Post{
		ID:      domain.NewID(entPost.ID),
		Title:   post.NewTitle(entPost.Title),
		Content: post.NewContent(entPost.Content),
		View:    post.NewView(entPost.View),
		User:    entPost.Edges.User.Mapper(),
	}
}

func (entPosts Posts) Mapper() domain.Posts {
	size := len(entPosts)
	domPosts := make(domain.Posts, size)
	for i := 0; i < size; i++ {
		domPosts[i] = entPosts[i].Mapper()
	}
	return domPosts
}
