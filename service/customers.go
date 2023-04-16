package service

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

type customers interface {
	GetListCustomers(ctx context.Context, fioFilter string, isArchive bool, claimID uint64) ([]domain.Customer, error)
	GetCustomerByID(ctx context.Context, id uint64) (domain.Customer, error)
	StoreCustomer(ctx context.Context, customer domain.Customer) (uint64, error)
}
