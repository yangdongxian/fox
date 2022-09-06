package controller

import (
	"fox/dto"
	"fox/helper"
	"fox/service"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

type WxUserController interface {
	Insert(context *gin.Context)
}

type wxUserController struct {
	wxUserService service.WxUserService
}

func NewWxUserController(wxUserService service.WxUserService) WxUserController {
	return &wxUserController{
		wxUserService: wxUserService,
	}
}

func (user *wxUserController) Insert(context *gin.Context) {
	var wxUserCreateDTO dto.WxUserCreateDTO
	errDTO := context.ShouldBind(&wxUserCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.Empt{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	wxUserCreateDTO.ID = uuid.NewV4().String()
	//wxUserCreateDTO.UnionId = "ovy1F6u7sw25-dsfJte8PswfBG1I"
	//wxUserCreateDTO.OpenId = "owLNV5O_2fRtSAR4D7_BtJvRkV1Y"
	//wxUserCreateDTO.SessionKey = "ITnoUENUKZEAXywoJ92lEg=="
	//wxUserCreateDTO.UnionId = wxUserInsertDTO.UnionId
	//wxUserCreateDTO.OpenId = wxUserInsertDTO.OpenId
	//wxUserCreateDTO.SessionKey = wxUserInsertDTO.SessionKey
	wxUserCreateDTO.IsDelete = false
	wxUserCreateDTO.CreateAt = time.Now()
	wxUserCreateDTO.UpdateAt = time.Now()
	wxUserCreateDTO.DeleteAt = time.Now()

	res := user.wxUserService.CreateWxUser(wxUserCreateDTO)
	response := helper.BuildResponse(true, "OK", res)
	context.JSON(http.StatusOK, response)
}
