package entity

import (
	"projeto-final/core/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id    string `gorm:"type:char(36);primaryKey"`
	Email string `gorm:"type:char(80);not null; unique"`
	Idade int    `gorm:"type:int;not null"`
}

func (p *User) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()
	p.Id = uuid
	return nil
}

func NewUserEntity(u *domain.User) (*User, error) {

	return &User{Id: u.Id, Email: u.Email, Idade: u.Idade}, nil
}

func (e User) ToDomain() *domain.User {
	return &domain.User{Id: e.Id, Email: e.Email, Idade: e.Idade}
}
