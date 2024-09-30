package controller

import (
	"github.com/gin-gonic/gin"
	"goWebODD/chapter5/controller/request"
	"goWebODD/chapter5/controller/response"
)

type Sms struct {
}

func NewSms() *Sms {
	ctr := new(Sms)
	return ctr
}

func (ctr *Sms) Captcha(c *gin.Context) (response.Data, error) {
	req := new(request.SmsCaptcha)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}

	return req, nil
}
