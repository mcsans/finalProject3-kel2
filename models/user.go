package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/mcsans/finalProject3-kel2/helpers"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName   string     `gorm:"not null" json:"full_name" form:"full_name" valid:"required"`
	Email      string     `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required,email"`
	Password   string     `gorm:"not null" json:"password" form:"password" valid:"required,minstringlength(6)"`
	Role       string     `gorm:"not null" json:"role" form:"role" valid:"required,matches(admin|member)"`
	Categories []Category `gorm:"constraint:onUpdate:CASCADE,OnDelete:SET NULL;" json:"categories"`
	Tasks      []Task     `gorm:"constraint:onUpdate:CASCADE,OnDelete:SET NULL;" json:"tasks"`
}

type UserResponse struct {
	ID        uint       `json:"id"`
	FullName  string     `json:"full_name"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}