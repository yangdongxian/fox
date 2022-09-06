package repository

import (
	"fox/common"
	"fox/entity"
	"gorm.io/gorm"
	"time"
)

type WxTokenRepository interface {
	InsertWxToken(wxToken entity.WxToken) entity.WxToken
	FindByAccessToken() entity.WxToken
}

type wxTokenConnection struct {
	connection *gorm.DB
}

func NewWxTokenConnection(db *gorm.DB) WxTokenRepository {
	return &wxTokenConnection{
		connection: db,
	}
}

func (db *wxTokenConnection) InsertWxToken(wxToken entity.WxToken) entity.WxToken {
	db.connection.Save(&wxToken)
	return wxToken
}

func (db *wxTokenConnection) FindByAccessToken() entity.WxToken {
	var token entity.WxToken
	now := time.Now().Add(-time.Second * common.GapTime)
	db.connection.Where("create_at > ?", now).Order("create_at desc").Find(&token)
	return token
}
