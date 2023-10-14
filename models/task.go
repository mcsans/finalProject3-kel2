package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Task struct {
	GormModel
	Title       string `gorm:"not null" json:"title" form:"title" valid:"required"`
	Description string `gorm:"not null" json:"description" form:"description" valid:"required"`
	Status      string `gorm:"not null" json:"status" form:"status" valid:"required,matches(1|0)"`
	UserID      uint
	CategoryID  uint	 `json:"category_id" form:"category_id" valid:"required"`
	User        *User
	Category    *Category
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(t)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (t *Task) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(t)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}