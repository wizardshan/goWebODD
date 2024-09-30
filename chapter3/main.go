package main

import (
	"github.com/gin-gonic/gin"
	"goWebODD/chapter3/controller"
)

func main() {

	engine := gin.New()

	ctrUser := controller.NewUser()
	engine.GET("/user/login", controller.Wrapper(ctrUser.Login))
	engine.GET("/user", controller.Wrapper(ctrUser.One))
	engine.GET("/users", controller.Wrapper(ctrUser.Many))

	ctrSms := controller.NewSms()
	engine.GET("/sms/captcha", controller.Wrapper(ctrSms.Captcha))

	engine.Run()
}
