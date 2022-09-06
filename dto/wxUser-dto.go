package dto

import "time"

type WxUserCreateDTO struct {
	ID         string    `json:"id"`
	OpenId     string    `json:"open_id" binding:"required"`
	UnionId    string    `json:"union_id" binding:"required"`
	SessionKey string    `json:"session_key" binding:"required"`
	IsDelete   bool      `json:"is_delete"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
	DeleteAt   time.Time `json:"delete_at"`
}
type WxUserInsertDTO struct {
	OpenId     string `json:"open_id" binding:"required"`
	UnionId    string `json:"union_id" binding:"required"`
	SessionKey string `json:"session_key" binding:"required"`
}

type WxUserUpdateDTO struct {
	SessionKey string    `json:"session_key" binding:"required"`
	IsDelete   bool      `json:"is_delete"`
	UpdateAt   time.Time `json:"update_at"`
}
