package domain

import (
	"goWebODD/test/controller/response"
	post2 "goWebODD/test/domain/post"
)

type Posts []Post

type Post struct {
	ID      ID
	Title   post2.Title
	Content post2.Content
	View    post2.View
	User    User
}

func (domPost Post) IsSet() bool {
	var zero Post
	return zero == domPost
}

func (domPost Post) IsMask() bool {
	return domPost.ID.Mask ||
		domPost.Title.Mask ||
		domPost.Content.Mask ||
		domPost.View.Mask ||
		domPost.User.IsMask()
}

func (domPost *Post) Mapper(queryPost Post) *response.Post {
	if domPost == nil {
		return nil
	}

	if !queryPost.IsMask() {
		return &response.Post{
			ID:      &domPost.ID.Value,
			Title:   &domPost.Title.Value,
			Content: &domPost.Content.Value,
			View:    &domPost.View.Value,
			User:    domPost.User.Mapper(queryPost.User),
		}
	}

	respPost := &response.Post{}
	if queryPost.ID.Mask {
		respPost.ID = &domPost.ID.Value
	}

	if queryPost.Title.Mask {
		respPost.Title = &domPost.Title.Value
	}

	if queryPost.Content.Mask {
		respPost.Content = &domPost.Content.Value
	}

	if queryPost.View.Mask {
		respPost.View = &domPost.View.Value
	}

	respPost.User = domPost.User.Mapper(queryPost.User)
	return respPost
}

func (domPosts Posts) Mapper(queryPost Post) response.Posts {
	size := len(domPosts)
	respPosts := make(response.Posts, size)
	for i := 0; i < size; i++ {
		respPosts[i] = domPosts[i].Mapper(queryPost)
	}
	return respPosts
}
