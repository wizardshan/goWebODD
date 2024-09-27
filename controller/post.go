package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goWebODD/controller/request"
	"goWebODD/controller/response"
	"goWebODD/repository"
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
	fields, ok := c.GetQueryMap("Fields")
	fmt.Println(fields, ok)

	domPost := ctr.repo.FetchByID(c, req.ID.Value)
	return domPost.Mapper(fields), nil
}
