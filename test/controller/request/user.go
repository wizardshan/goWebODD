package request

import (
	domain2 "goWebODD/test/domain"
	user2 "goWebODD/test/domain/user"
)

type UserOne struct {
	ID domain2.ID
}

type UserPath struct {
	ID domain2.ID
}

type UserMany struct {
	Level  user2.Level    `binding:"omitempty"`
	Age    user2.Age      `binding:"omitempty"`
	Mobile domain2.Mobile `binding:"omitempty"`
}
