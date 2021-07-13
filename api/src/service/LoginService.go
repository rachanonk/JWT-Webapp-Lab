package service

type LoginService interface {
	LoginUser(email string, password string) bool
}
type loginInformation struct {
	email    string
	password string
}

func StaticLoginService() LoginService {

	return &loginInformation{
		email:    "skrrachanon@gmail.com",
		password: "p@ssw0rd",
	}
}
func (info *loginInformation) LoginUser(email string, password string) bool {
	return (info.email == email && info.password == password)
}
