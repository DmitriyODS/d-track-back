package implementation

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

func (bs *BasicService) GetListCustomers(ctx context.Context, fioFilter string, isArchive bool) ([]domain.Customer, error) {
	return nil, nil
}

func (bs *BasicService) GetCustomerByID(ctx context.Context, id uint64) (domain.Customer, error) {
	return domain.Customer{}, nil
}

func (bs *BasicService) StoreCustomer(ctx context.Context, customer domain.Customer) (uint64, error) {
	return 0, nil
}
