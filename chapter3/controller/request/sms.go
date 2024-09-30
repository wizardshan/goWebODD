package request

import "goWebODD/chapter3/domain"

type SmsCaptcha struct {
	Mobile domain.Mobile `binding:"required"`
}
