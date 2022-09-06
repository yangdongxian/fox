package repository

import (
	"fmt"
	"fox/entity"
	"gorm.io/gorm"
)

type WxUserRepository interface {
	CreateWxUser(wxUser entity.WxUser) entity.WxUser
	UpdateWxUser(wxUser entity.WxUser) entity.WxUser
	FindByOpenId(openId string) entity.WxUser
}

type wxUserRepository struct {
	connection *gorm.DB
}

func NewWxUserRepository(db *gorm.DB) WxUserRepository {
	return &wxUserRepository{
		connection: db,
	}
}
func (db *wxUserRepository) CreateWxUser(wxUser entity.WxUser) entity.WxUser {
	var temWxUser = entity.WxUser{}
	res := db.connection.Where("open_id", wxUser.OpenId).Or("union_id", wxUser.UnionId).Find(&temWxUser)
	//fmt.Printf("RowsAffected:%v temWxUser:%#v \n", res.RowsAffected, temWxUser)
	if res.RowsAffected <= 0 {
		db.connection.Save(wxUser)
	} else {
		return temWxUser
	}
	return wxUser
}
func (db *wxUserRepository) UpdateWxUser(wxUser entity.WxUser) entity.WxUser {
	var temWxUser = entity.WxUser{}
	res := db.connection.Where("open_id", wxUser.OpenId).Or("union_id", wxUser.UnionId).Find(&temWxUser)
	fmt.Printf("RowsAffected:%v temWxUser:%#v \n", res.RowsAffected, temWxUser)
	if res.RowsAffected <= 0 {
		db.connection.Save(wxUser)
	} else {
		return temWxUser
	}
	return wxUser
}
func (db *wxUserRepository) FindByOpenId(openId string) entity.WxUser {
	var wxUser entity.WxUser
	db.connection.Find("open_id = ?", openId).Take(&wxUser)
	return wxUser
}
