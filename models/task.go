package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Task struct {
	GormModel
	Title       string `gorm:"not null" json:"title" form:"title"`
	Status      bool   `gorm:"not null" json:"status"`
	Description string `gorm:"not null" json:"description" form:"description"`
	UserID      uint
	CategoryID  uint 	 `json:"category_id" form:"category_id"`
	User        *User
	Category    *Category
}

type TaskResponse struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Status      bool       `json:"status"`
	Description string     `json:"description"`
	UserID      uint			 `json:"user_id"`
	CategoryID  uint 	 		 `json:"category_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
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