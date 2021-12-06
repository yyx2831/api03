package forms

type PassWordLoginForm struct {
	//Email    string `form:"email" json:"email" binding:"required,email"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号
	Password string `form:"password" json:"password" binding:"required,min=6,max=20"`
}
