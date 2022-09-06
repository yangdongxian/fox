package entity

import "time"

type WxUser struct {
	ID         string    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	OpenId     string    `gorm:"unique;not null;type:varchar(512)" json:"open_id"`
	UnionId    string    `gorm:"unique;not null;type:varchar(512)" json:"union_id"`
	SessionKey string    `gorm:"not null;type:varchar(512)" json:"session_key"`
	IsDelete   bool      `json:"is_delete"`
	CreateAt   time.Time `gorm:"index" json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
	DeleteAt   time.Time `json:"delete_at"`
}
