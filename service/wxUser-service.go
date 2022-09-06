package service

import (
	"fox/dto"
	"fox/entity"
	"fox/repository"
	"github.com/mashingan/smapping"
	"log"
)

type WxUserService interface {
	CreateWxUser(user dto.WxUserCreateDTO) entity.WxUser
	FindByOpenId(openId string) entity.WxUser
}

type wxUserService struct {
	wxUserRep repository.WxUserRepository
}

func NewWwxUserService(wxUser repository.WxUserRepository) WxUserService {
	return &wxUserService{wxUserRep: wxUser}
}
func (w *wxUserService) CreateWxUser(user dto.WxUserCreateDTO) entity.WxUser {
	wxUser := entity.WxUser{}
	err := smapping.FillStruct(&wxUser, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	insertWxUser := w.wxUserRep.CreateWxUser(wxUser)
	return insertWxUser
}

func (w *wxUserService) FindByOpenId(openId string) entity.WxUser {
	return w.wxUserRep.FindByOpenId(openId)
}
