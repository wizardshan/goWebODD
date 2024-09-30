package controller

import (
	"github.com/gin-gonic/gin"
	"goWebODD/repository"
	"goWebODD/test/controller/request"
	"goWebODD/test/controller/response"
	domain2 "goWebODD/test/domain"
	"strconv"
)

type User struct {
	repo *repository.User
}

func NewUser(repo *repository.User) *User {
	ctr := new(User)
	ctr.repo = repo
	return ctr
}

func (ctr *User) One(c *gin.Context) (response.Data, error) {
	req := new(request.UserOne)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}

	queryUser := domain2.User{
		ID: req.ID,
	}

	domUser := ctr.repo.FetchByID(c, req.ID.Value)
	return domUser.Mapper(queryUser), nil
}

func (ctr *User) OnePath(c *gin.Context) (response.Data, error) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	req := new(request.UserPath)
	req.ID = domain2.NewID(id)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}
	queryUser := domain2.User{
		ID: req.ID,
	}

	domUser := ctr.repo.FetchByID(c, req.ID.Value)
	return domUser.Mapper(queryUser), nil
}

func (ctr *User) Many(c *gin.Context) (response.Data, error) {
	req := new(request.UserMany)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}
	queryUser := domain2.User{
		Level:  req.Level,
		Age:    req.Age,
		Mobile: req.Mobile,
	}

	domUser := ctr.repo.FetchMany(c, queryUser)
	return domUser.Mapper(queryUser), nil
}

//func (ctr *User) Show(c *gin.Context) (response.Data, error) {
//	req := request.User{HashID: c.Param("hashID")}
//	domUser, err := req.Mapper()
//	if err != nil {
//		return nil, err
//	}
//
//	domUser = ctr.repo.FetchByID(c, domUser.HashID.ID)
//	return domUser.Mapper(), nil
//}
//
//func (ctr *User) One(c *gin.Context) (response.Data, error) {
//	req := new(request.User)
//	if err := c.ShouldBind(req); err != nil {
//		return nil, err
//	}
//
//	domUser, err := req.Mapper()
//	if err != nil {
//		return nil, err
//	}
//
//	domUser = ctr.repo.FetchByID(c, domUser.HashID.ID)
//	return domUser.Mapper(), nil
//}
//
//func (ctr *User) Login(c *gin.Context) (response.Data, error) {
//	req := new(request.UserLogin)
//	if err := c.ShouldBind(req); err != nil {
//		return nil, err
//	}
//
//	domUser, domCaptcha, err := req.Mapper()
//	if err != nil {
//		return nil, err
//	}
//	fmt.Println(domUser, domCaptcha)
//	return nil, nil
//}
//
//func (ctr *User) Register(c *gin.Context) (response.Data, error) {
//	req := new(request.UserRegister)
//	if err := c.ShouldBind(req); err != nil {
//		return nil, err
//	}
//
//	domUser, err := req.Mapper()
//	if err != nil {
//		return nil, err
//	}
//
//	domUser = ctr.repo.FetchByID(c, domUser.HashID.ID)
//	return domUser.Mapper(), nil
//}
