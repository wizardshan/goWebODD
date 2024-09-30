package controller

import (
	"github.com/gin-gonic/gin"
	"goWebODD/repository"
	"goWebODD/test/controller/request"
	"goWebODD/test/controller/response"
	"goWebODD/test/domain"
)

type Post struct {
	repo *repository.Post
}

func NewPost(repo *repository.Post) *Post {
	ctr := new(Post)
	ctr.repo = repo
	return ctr
}

func (ctr *Post) One(c *gin.Context) (response.Data, error) {
	req := new(request.PostOne)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}

	queryPost := domain.Post{
		ID:      req.ID,
		Title:   req.Title,
		Content: req.Content,
		User:    req.User,
	}

	domPost := ctr.repo.FetchByID(c, req.ID.Value)
	return domPost.Mapper(queryPost), nil
}
