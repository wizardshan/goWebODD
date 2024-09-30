package controller

import (
	"github.com/gin-gonic/gin"
	"goWebODD/chapter1/controller/request"
	"goWebODD/chapter1/controller/response"
)

type User struct {
}

func NewUser() *User {
	ctr := new(User)

	return ctr
}

func (ctr *User) LoginV1(c *gin.Context) (response.Data, error) {
	req := new(request.UserLoginV1)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (ctr *User) LoginV2(c *gin.Context) (response.Data, error) {
	req := new(request.UserLoginV2)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (ctr *User) Login(c *gin.Context) (response.Data, error) {
	req := new(request.UserLogin)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (ctr *User) One(c *gin.Context) (response.Data, error) {
	req := new(request.UserOne)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}

	return req, nil
}

func (ctr *User) Many(c *gin.Context) (response.Data, error) {
	req := new(request.UserMany)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}

	return req, nil
}
