package services

import (
	"FP-RPL-ECommerce/dto"
	"FP-RPL-ECommerce/entity"
	"FP-RPL-ECommerce/repository"
	"FP-RPL-ECommerce/utils"
	"context"
	"fmt"

	"github.com/jinzhu/copier"
)

type custSvc struct {
	custRepo repository.CustRepo
}

type CustSvc interface {
	RegisterCust(ctx context.Context, cust dto.UserCreate) (entity.User, error)
	VerifyCust(ctx context.Context, email string, password string) (bool, error)
	FindCustByEmail(ctx context.Context, email string) (entity.User, error)
	FindCustByID(ctx context.Context, id uint64) (entity.User, error)
	GetAllCust(ctx context.Context) (entity.User, error)
	UpdateCust(ctx context.Context, custParam dto.UserUpdate, id uint64) (entity.User, error)
	DeleteCust(ctx context.Context, id uint64) (entity.User, error)
}

func NewCustSvc(repo repository.CustRepo) CustSvc {
	return &custSvc{
		custRepo: repo,
	}
}

func (svc *custSvc) RegisterCust(ctx context.Context, custParam dto.UserCreate) (entity.User, error) {
	var cust entity.User
	copier.Copy(&cust, &custParam)

	createdCust, err := svc.custRepo.RegisterCust(ctx, cust)
	if err != nil {
		return entity.User{}, err
	}
	return createdCust, nil
}

func (svc *custSvc) VerifyCust(ctx context.Context, email string, password string) (bool, error) {
	cust, err := svc.custRepo.CheckEmailCust(ctx, email)
	if err != nil {
		fmt.Println("email salah")
		return false, err
	}

	CheckedPwd, err := utils.ComparePassword(cust.Password, password)
	if err != nil {
		fmt.Println("password salah")
		return false, err
	}

	if cust.Email == email && CheckedPwd {
		return true, nil
	}

	return false, nil
}

func (svc *custSvc) FindCustByEmail(ctx context.Context, email string) (entity.User, error) {
	cust, err := svc.custRepo.CheckEmailCust(ctx, email)
	if err != nil {
		return entity.User{}, err
	}
	return cust, nil
}

func (svc *custSvc) FindCustByID(ctx context.Context, id uint64) (entity.User, error) {
	cust, err := svc.custRepo.CheckIDCust(ctx, id)
	if err != nil {
		return entity.User{}, err
	}
	return cust, nil
}

func (svc *custSvc) GetAllCust(ctx context.Context) (entity.User, error) {
	cust, err := svc.custRepo.GetAllCust(ctx)
	if err != nil {
		return entity.User{}, err
	}
	return cust, nil
}

func (svc *custSvc) UpdateCust(ctx context.Context, custParam dto.UserUpdate, id uint64) (entity.User, error) {
	var cust entity.User
	copier.Copy(&cust, &custParam)

	custParam.ID = id
	updatedCust, err := svc.custRepo.UpdateCust(ctx, cust, id)
	if err != nil {
		return entity.User{}, err
	}
	return updatedCust, nil
}

func (svc *custSvc) DeleteCust(ctx context.Context, id uint64) (entity.User, error) {
	cust, err := svc.custRepo.DeleteCust(ctx, id)
	if err != nil {
		return entity.User{}, err
	}
	return cust, nil
}
