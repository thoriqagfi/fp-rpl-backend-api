package services

import (
	"FP-RPL-ECommerce/dto"
	"FP-RPL-ECommerce/entity"
	"FP-RPL-ECommerce/repository"
	"FP-RPL-ECommerce/utils"
	"context"

	"github.com/jinzhu/copier"
)

type sellerSvc struct {
	sellerRepo repository.SellerRepo
}

type SellerSvc interface {
	RegisterSeller(ctx context.Context, sellerParam dto.UserCreate) (entity.User, error)
	VerifySeller(ctx context.Context, email string, password string) (bool, error)
	FindSellerByEmail(ctx context.Context, email string) (entity.User, error)
}

func NewSellerSvc(repo repository.SellerRepo) SellerSvc {
	return &sellerSvc{
		sellerRepo: repo,
	}
}

func (svc *sellerSvc) RegisterSeller(ctx context.Context, sellerParam dto.UserCreate) (entity.User, error) {
	var seller entity.User
	copier.Copy(&seller, &sellerParam)

	createdCust, err := svc.sellerRepo.Register(ctx, seller)
	if err != nil {
		return entity.User{}, err
	}
	return createdCust, nil
}

func (svc *sellerSvc) VerifySeller(ctx context.Context, email string, password string) (bool, error) {
	seller, err := svc.sellerRepo.CheckEmailSeller(ctx, email)
	if err != nil {
		return false, err
	}

	CheckedPwd, err := utils.ComparePassword(seller.Password, password)
	if err != nil {
		return false, err
	}

	if seller.Email == email && CheckedPwd {
		return true, nil
	}

	return false, nil
}

func (svc *sellerSvc) FindSellerByEmail(ctx context.Context, email string) (entity.User, error) {
	seller, err := svc.sellerRepo.CheckEmailSeller(ctx, email)
	if err != nil {
		return entity.User{}, err
	}
	return seller, nil
}
