package controller

import (
	"fox/dto"
	"fox/entity"
	"fox/helper"
	"fox/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authControler struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authControler{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authControler) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.Empt{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK", v)
		//fmt.Printf("response:%#v\n", v)

		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("please check again your credential", "Invalid Credential", helper.Empt{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authControler) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.Empt{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.Empt{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createUser.ID, 10))
		createUser.Token = token
		response := helper.BuildResponse(true, "OK", createUser)
		ctx.JSON(http.StatusOK, response)

	}
}
