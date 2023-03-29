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
	CheckIDCust(ctx context.Context, id uint64) (cust entity.User, err error)
	GetAllCust(ctx context.Context) (cust entity.User, err error)
	UpdateCust(ctx context.Context, cust entity.User, id uint64) (entity.User, error)
	DeleteCust(ctx context.Context, id uint64) (cust entity.User, err error)
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

func (repo *custRepo) CheckIDCust(ctx context.Context, id uint64) (cust entity.User, err error) {
	tx := repo.db.Where("id = ?", id).Take(&cust)
	if tx.Error != nil {
		return entity.User{}, err
	}
	return cust, nil
}

func (repo *custRepo) GetAllCust(ctx context.Context) (cust entity.User, err error) {
	tx := repo.db.Where("role = 'customer'").Take(&cust)
	if tx.Error != nil {
		return entity.User{}, err
	}
	return cust, nil
}

func (repo *custRepo) UpdateCust(ctx context.Context, cust entity.User, id uint64) (entity.User, error) {
	var err error
	tx := repo.db.Where("id = ?", id).Updates(&cust)
	if tx.Error != nil {
		return entity.User{}, err
	}
	return cust, nil
}

// admin bisanya
func (repo *custRepo) DeleteCust(ctx context.Context, id uint64) (cust entity.User, err error) {
	tx := repo.db.Delete(&cust, &id)
	if tx.Error != nil {
		return entity.User{}, err
	}
	return cust, nil
}
