package controller

import (
	"fox/helper"
	"fox/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WxTokenController interface {
	GetRedisAccessToken(c *gin.Context)
	GetDbAccessToken(context *gin.Context)
}

type wxTokenController struct {
	tokenService service.WxTokenService
}

func NewWxTokenController(tokenService service.WxTokenService) WxTokenController {
	return &wxTokenController{
		tokenService: tokenService,
	}
}

func (t *wxTokenController) GetRedisAccessToken(context *gin.Context) {
	accessToken := t.tokenService.GetRedisAccessToken()
	if accessToken == "" {
		helper.BuildErrorResponse("accessToken is invalide error", "", gin.H{})
	}
	response := helper.BuildResponse(true, "OK", gin.H{"access_token": accessToken})
	context.JSON(http.StatusOK, response)
}
func (t *wxTokenController) GetDbAccessToken(context *gin.Context) {
	userToken := t.tokenService.FindByAccessToken()
	if userToken.AccessToken == "" {
		helper.BuildErrorResponse("accessToken is invalide error", "", gin.H{})
	}
	response := helper.BuildResponse(true, "OK", userToken)
	context.JSON(http.StatusOK, response)

}
