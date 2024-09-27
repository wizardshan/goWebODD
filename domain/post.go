package domain

import (
	"goWebODD/controller/response"
	"goWebODD/domain/post"
	"strings"
)

type Posts []*Post

type Post struct {
	ID      ID
	Title   post.Title
	Content post.Content
	View    post.View
	User    *User
}

func (domPost *Post) Name() string {
	return "Post"
}

func (domPost *Post) Fields() []string {
	return []string{
		domPost.ID.Name(),
		domPost.Title.Name(),
		domPost.Content.Name(),
		domPost.View.Name(),
		domPost.User.Name(),
	}
}

func (domPost *Post) Mapper(fields map[string]string) *response.Post {
	if domPost == nil {
		return nil
	}

	if len(fields) == 0 {
		return &response.Post{
			ID:      &domPost.ID.Value,
			Title:   &domPost.Title.Value,
			Content: &domPost.Content.Value,
			View:    &domPost.View.Value,
			User:    domPost.User.Mapper(),
		}
	}

	respPost := &response.Post{}

	if fs, ok := fields[domPost.Name()]; ok {
		if strings.Contains(fs, domPost.ID.Name()) {
			respPost.ID = &domPost.ID.Value
		}
		if strings.Contains(fs, domPost.Title.Name()) {
			respPost.Title = &domPost.Title.Value
		}
		if strings.Contains(fs, domPost.Content.Name()) {
			respPost.Content = &domPost.Content.Value
		}
		if strings.Contains(fs, domPost.View.Name()) {
			respPost.View = &domPost.View.Value
		}
	}

	return respPost
}

func (domPosts Posts) Mapper(fields map[string]string) response.Posts {
	size := len(domPosts)
	respPosts := make(response.Posts, size)
	for i := 0; i < size; i++ {
		respPosts[i] = domPosts[i].Mapper(fields)
	}
	return respPosts
}
