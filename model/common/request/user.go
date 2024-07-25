package request

type UserLogin struct {
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
}
