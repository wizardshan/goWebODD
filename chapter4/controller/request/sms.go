package request

import "goWebODD/chapter4/domain"

type SmsCaptcha struct {
	Mobile domain.Mobile `binding:"required"`
}
