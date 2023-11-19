package model

import (
	"github.com/jinzhu/gorm"
)

type Contract struct {
	ID                  uint `gorm:"primarykey"`
	ApplicationId       int64
	CreditAccountNumber string `gorm:"unique;not null"`
	Status              bool
	Message             string
	gorm.Model          // TimeStamps

}

type Users struct {
	ID       uint   `gorm:"primarykey"`
	UserName string `gorm:"unique;not null"`
	Status   bool
}
