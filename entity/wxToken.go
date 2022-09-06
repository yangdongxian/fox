package entity

import "time"

type WxToken struct {
	//ID          string    `gorm:"primary_key,type:varchar(36)" json:"id"`
	ID          string    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	AccessToken string    `gorm:"not null;type:varchar(512)" json:"access_token"`
	ExpiresIn   int64     `gorm:"not null" json:"expires_in"`
	IsDelete    bool      `json:"is_delete"`
	CreateAt    time.Time `gorm:"index" json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
	DeleteAt    time.Time `json:"delete_at"`
}
