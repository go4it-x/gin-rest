package model

// User object
type User struct {
	Model
	Username string `gorm:"type:varchar(20);unique;not null;comment:'名称'" json:"username"`
	Password string `gorm:"type:varchar(80);not null;comment:'密码'" json:"password"`
}
