package request

type UserLoginV1 struct {
	Mobile string `binding:"required,number,mobile"`
}

type UserLoginV2 struct {
	MobileFieldV2
}

type UserLogin struct {
	MobileField `binding:"required"`
}

type UserOne struct {
	MobileField `binding:"required"`
}

type UserMany struct {
	MobileField `binding:"omitempty"`
}
