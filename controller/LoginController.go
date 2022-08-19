package controller

import (
	"fmt"
	"fox/LoCred"
	"fox/service"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(c *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func LoginHandler(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (controller *loginController) Login(c *gin.Context) string {
	var credential LoCred.LoginCredentials
	err := c.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	fmt.Println("shouldBind:", credential)
	isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controller.jwtService.GenerateToken(credential.Email)
	}
	return ""
}
