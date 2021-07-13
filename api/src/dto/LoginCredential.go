package dto

//Login credential
type LoginCredentials struct {
	Email     string `form:"email"`
	Password  string `form:"password"`
	Privilege string `form:"privilege"`
}
