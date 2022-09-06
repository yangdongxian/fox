package entity

//User user infomation
type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(256)" json:"-"`
	Email    string `gorm:"uniqueIndex;type:varchar(256)" json:"-"`
	Password string `gorm:"<-;->;not null" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`
}
