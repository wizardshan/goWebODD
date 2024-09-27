package request

import (
	"goWebODD/domain"
	"goWebODD/domain/user"
)

type UserOne struct {
	ID domain.ID
}

type UserPath struct {
	ID domain.ID
}

type UserMany struct {
	Level  user.Level    `binding:"omitempty"`
	Age    user.Age      `binding:"omitempty"`
	Mobile domain.Mobile `binding:"omitempty"`
}
