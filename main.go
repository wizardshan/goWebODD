package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"goWebODD/repository"
	"goWebODD/repository/ent"
	controller2 "goWebODD/test/controller"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/bbs?charset=utf8mb4&parseTime=true"
	db, err := ent.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	engine := gin.New()
	repoUser := repository.NewUser(db)
	ctrUser := controller2.NewUser(repoUser)
	engine.GET("/user", controller2.Wrapper(ctrUser.One))
	engine.GET("/user/:id", controller2.Wrapper(ctrUser.OnePath))
	engine.GET("/users", controller2.Wrapper(ctrUser.Many))

	ctrSms := controller2.NewSms()
	engine.GET("/sms/captcha", controller2.Wrapper(ctrSms.Captcha))

	repoPost := repository.NewPost(db)
	ctrPost := controller2.NewPost(repoPost)
	engine.GET("/post", controller2.Wrapper(ctrPost.One))

	engine.Run()
}
