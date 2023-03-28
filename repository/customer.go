package repository

import (
	"FP-RPL-ECommerce/entity"
	"context"
	"log"

	"gorm.io/gorm"
)

type custRepo struct {
	db *gorm.DB
}

type CustRepo interface {
	RegisterCust(ctx context.Context, cust entity.User) (entity.User, error)
	CheckEmailCust(ctx context.Context, email string) (cust entity.User, err error)
}

func NewCustRepo(db *gorm.DB) CustRepo {
	return &custRepo{
		db: db,
	}
}

func (repo *custRepo) RegisterCust(ctx context.Context, cust entity.User) (entity.User, error) {
	var err error
	tx := repo.db.Create(&cust)
	if tx.Error != nil {
		log.Println(err)
		return entity.User{}, err
	}
	return cust, nil
}

func (repo *custRepo) CheckEmailCust(ctx context.Context, email string) (cust entity.User, err error) {
	tx := repo.db.Where("email = ?", email).Take(&cust)
	if tx.Error != nil {
		return entity.User{}, err
	}
	return cust, nil
}
