package repository

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

type repositoryCustomers interface {
	SelectCustomers(ctx context.Context, fioFilter string, isArchive bool) ([]domain.Customer, error)
	SelectCustomerByID(ctx context.Context, id uint64) (domain.Customer, error)
	CreateCustomer(ctx context.Context, customer domain.Customer) (uint64, error)
	UpdateCustomer(ctx context.Context, customer domain.Customer) (uint64, error)
}
