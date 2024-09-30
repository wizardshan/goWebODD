package request

import "goWebODD/chapter5/domain"

type SmsCaptcha struct {
	Mobile domain.Mobile `binding:"required"`
}
