package ent

import (
	domain2 "goWebODD/test/domain"
	post2 "goWebODD/test/domain/post"
)

func (entPost *Post) Mapper() domain2.Post {
	var domPost domain2.Post
	if entPost == nil {
		return domPost
	}

	domPost.ID = domain2.NewID(entPost.ID)
	domPost.Title = post2.NewTitle(entPost.Title)
	domPost.Content = post2.NewContent(entPost.Content)
	domPost.View = post2.NewView(entPost.View)

	domPost.User = entPost.Edges.User.Mapper()

	return domPost
}

func (entPosts Posts) Mapper() domain2.Posts {
	size := len(entPosts)
	domPosts := make(domain2.Posts, size)
	for i := 0; i < size; i++ {
		domPosts[i] = entPosts[i].Mapper()
	}
	return domPosts
}
