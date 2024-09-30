package request

import "goWebODD/chapter2/domain"

type SmsCaptcha struct {
	Mobile domain.Mobile `binding:"required"`
}
