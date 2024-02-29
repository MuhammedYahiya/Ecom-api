package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `json:"username" validate:"required,min=8,max=24" `
	Email       string `json:"email" validate:"email,required"`
	Phone       string `json:"phone" validate:"required,len=10"`
	Password    string `json:"password" validate:"required,min=8,max=16"`
	Otp         string `json:"otp"`
	IsVerified  bool   `json:"isverified" gorm:"default:false"`
	IsAdmin     bool   `json:"isadmin" gorm:"default:false"`
	Dateofbirth string `json:"dateofbirth"`
	Gender      string `json:"gender"`
}
