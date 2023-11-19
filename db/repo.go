package db

import (
	"gorm.io/gorm"
	"ms-go-initial/model"
)

type IRepository interface {
	SaveUser(userName string, status bool) error
	UpdateUser(userId uint, userName string, status bool) error
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &repository{db}
}

func (r repository) SaveUser(userName string, status bool) error {

	return r.DB.Create(&model.Users{
		UserName: userName,
		Status:   status,
	}).Error

}

func (r repository) UpdateUser(userId uint, userName string, status bool) error {

	return r.DB.Save(&model.Users{
		ID:       userId,
		UserName: userName,
		Status:   status,
	}).Error

	//return r.DB.Update(&model.Users{
	//	UserName: userName,
	//	Status:   status,
	//}).Error

}
