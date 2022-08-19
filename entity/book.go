package entity

//Book book
type Book struct {
	ID uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title string `gorm:"type:varchar(256)" json:"title"`
	Description string `gorm:"type:varchar(256)" json:"description"`
	UserID uint64 `gorm:""`
	User 	User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE" json:"user"`
}