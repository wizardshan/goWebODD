package request

type MobileFieldV2 struct {
	Mobile string `binding:"required,number,mobile"`
}

type MobileField struct {
	Mobile string `binding:"number,mobile"`
}
