package service

type LoginService interface {
	LoginUser(email string,password string)bool
}

type LoginInfomation struct {
	email string
	password string
}

func NewLoginInfomation() LoginService {
	return &LoginInfomation{
		email: "eastsheen@gmail.com",
		password: "eastsheen",
	}
}
func (info *LoginInfomation) LoginUser(email string,password string) bool{
	return info.email == email && info.password == password
}
