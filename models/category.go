package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Category struct {
	GormModel
	Type   string `json:"type" form:"type" valid:"required"`
	UserID uint
	User   *User
	Tasks  []Task `gorm:"constraint:onUpdate:CASCADE,OnDelete:SET NULL;" json:"tasks"`
}

type CategoryResponse struct {
	ID        uint   		 `json:"id"`
	Type      string 		 `json:"type"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
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