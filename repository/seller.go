package repository

import (
	"FP-RPL-ECommerce/entity"
	"context"
	"log"

	"gorm.io/gorm"
)

type sellerRepo struct {
	db *gorm.DB
}

type SellerRepo interface {
	Register(ctx context.Context, seller entity.User) (entity.User, error)
	CheckEmailSeller(ctx context.Context, email string) (seller entity.User, err error)
}

func NewSellerRepo(db *gorm.DB) SellerRepo {
	return &sellerRepo{
		db: db,
	}
}

func (repo *sellerRepo) Register(ctx context.Context, seller entity.User) (entity.User, error) {
	var err error
	tx := repo.db.WithContext(ctx).Debug().Create(&seller)
	if tx.Error != nil {
		log.Println(err)
		return entity.User{}, err
	}
	return seller, nil
}

func (repo *sellerRepo) CheckEmailSeller(ctx context.Context, email string) (seller entity.User, err error) {
	tx := repo.db.Where("email = ?", email).Take(&seller)
	if tx.Error != nil {
		return entity.User{}, err
	}
	return seller, nil
}
