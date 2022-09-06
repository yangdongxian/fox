package dto

import "time"

type WxTokenCreateDTO struct {
	ID          string    `json:"id"`
	AccessToken string    `json:"access_token" binding:"required"`
	ExpiresIn   int64     `json:"expires_in" binding:"required"`
	IsDelete    bool      `json:"is_delete"`
	CreateAt    time.Time `json:"create_at" binding:"required"`
	UpdateAt    time.Time `json:"update_at"`
	DeleteAt    time.Time `json:"delete_at"`
}
