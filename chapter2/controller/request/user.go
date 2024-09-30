package request

import "goWebODD/chapter2/domain"

type UserLogin struct {
	Mobile domain.Mobile `binding:"required"`
}

type UserOne struct {
	Mobile domain.Mobile `binding:"required"`
}

type UserMany struct {
	Mobile domain.Mobile `binding:"omitempty"`
}
