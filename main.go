package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"goWebODD/controller"
	"goWebODD/repository"
	"goWebODD/repository/ent"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/bbs?charset=utf8mb4&parseTime=true"
	db, err := ent.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	engine := gin.New()
	repoUser := repository.NewUser(db)
	ctrUser := controller.NewUser(repoUser)
	engine.GET("/user", controller.Wrapper(ctrUser.One))
	engine.GET("/user/:id", controller.Wrapper(ctrUser.OnePath))
	engine.GET("/users", controller.Wrapper(ctrUser.Many))

	ctrSms := controller.NewSms()
	engine.GET("/sms/captcha", controller.Wrapper(ctrSms.Captcha))

	engine.Run()
}
