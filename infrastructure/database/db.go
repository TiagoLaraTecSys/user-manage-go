package database

import (
	"context"
	"projeto-final/adapter/database"
	"projeto-final/core/domain"
	"projeto-final/core/erros"
	"projeto-final/infrastructure/database/entity"
	"projeto-final/infrastructure/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQLConnection struct {
	db *gorm.DB
}

var _ database.DbUser = (*SQLConnection)(nil)

func NewSQLConnection(dns string) (*SQLConnection, error) {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		return &SQLConnection{}, err
	}

	db.AutoMigrate(&entity.User{})
	return &SQLConnection{db: db}, nil
}

func (s *SQLConnection) Add(ctx *context.Context, user *domain.User) (domain.User, error) {
	u, err := entity.NewUserEntity(user)
	if err != nil {
		return domain.User{}, err
	}
	result := s.db.Create(u)

	if result.Error != nil {
		return domain.User{}, result.Error
	}
	user.Id = u.Id
	return *user, nil
}

func (s *SQLConnection) GetById(ctx *context.Context, Id string) (domain.User, error) {
	var user entity.User
	result := s.db.Where("id=?", Id).Find(&user)

	logger.Info("Erro", result.Error)

	if result.Error != nil {
		return domain.User{}, result.Error
	}

	if user.Id == "" {
		return domain.User{}, erros.NewNotFoundErr("User", Id)
	}
	logger.Info("User", user)

	return *user.ToDomain(), nil
}

func (s *SQLConnection) GetUsers(ctx *context.Context) ([]domain.User, error) {
	var users []entity.User

	result := s.db.Find(&users)

	if result.Error != nil {
		logger.Info("AAAAAAA")
		return []domain.User{}, result.Error
	}

	uList, err := entityToUserList(&users)
	if err != nil {
		return []domain.User{}, err
	}

	return uList, nil
}

func entityToUserList(list *[]entity.User) ([]domain.User, error) {
	var domainList []domain.User

	for _, u := range *list {
		u := u.ToDomain()
		domainList = append(domainList, *u)
	}
	return domainList, nil
}
