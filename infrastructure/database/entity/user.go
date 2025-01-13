package entity

import (
	"projeto-final/core/domain"
)

type User struct {
	Id    int    `gorm:"type:int;autoIncrement;primaryKey"`
	Email string `gorm:"type:varchar(100);not null; unique"`
	Name  string `gorm:"type:varchar(100)"`
	Idade int    `gorm:"type:int;not null"`
}

func NewUserEntity(u *domain.User) (*User, error) {

	return &User{Id: u.Id, Email: u.Email, Idade: u.Idade}, nil
}

func (e User) ToDomain() *domain.User {
	return &domain.User{Id: e.Id, Email: e.Email, Idade: e.Idade}
}
