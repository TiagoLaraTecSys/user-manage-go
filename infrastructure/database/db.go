package database

import (
	"context"
	"projeto-final/adapter/database"
	"projeto-final/core/domain"
	"projeto-final/core/erros"
	"projeto-final/core/usecase/input"
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

func (s *SQLConnection) GetUsers(ctx *context.Context, i *input.PaginationInput) (domain.Data, error) {
	var users []entity.User

	limit := -1
	offset := -1
	if i.Limit != 0 {
		limit = i.Limit
	}

	if i.Page != 0 {
		offset = (i.Page - 1) * limit
	}

	result := s.db.Offset(offset).Limit(limit).Find(&users)
	var totalItems int64
	s.db.Find(&users).Count(&totalItems)
	logger.Info("Total de itens %d", totalItems)
	totalPages := (int(totalItems) + limit - 1) / limit
	if result.Error != nil {
		return domain.Data{}, result.Error
	}

	uList, err := entityToUserList(&users)
	if err != nil {
		return domain.Data{}, err
	}
	p := &domain.Pagination{Limit: limit, Page: i.Page, TotalPages: totalPages}
	d := domain.Data{Users: uList, Page: *p}

	return d, nil
}

func entityToUserList(list *[]entity.User) ([]domain.User, error) {
	var domainList []domain.User

	for _, u := range *list {
		u := u.ToDomain()
		domainList = append(domainList, *u)
	}
	return domainList, nil
}
