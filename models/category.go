package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Category struct {
	GormModel
	Type       	string `json:"type" form:"type" valid:"required"`
	UserID      uint
	User        *User
	Tasks []Task	 			 `gorm:"constraint:onUpdate:CASCADE,OnDelete:SET NULL;" json:"tasks"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (c *Category) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}