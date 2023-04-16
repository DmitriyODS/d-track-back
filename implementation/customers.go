package implementation

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"log"
)

func (bs *BasicService) GetListCustomers(ctx context.Context, fioFilter string, isArchive bool, claimID uint64) ([]domain.Customer, error) {
	customers, err := bs.rep.SelectCustomers(ctx, fioFilter, isArchive, claimID)
	if err != nil {
		log.Println("GetListCustomers err:", err)
		return nil, global.InternalServerErr
	}

	return customers, nil
}

func (bs *BasicService) GetCustomerByID(ctx context.Context, id uint64) (domain.Customer, error) {
	if id == 0 {
		return domain.NewCustomer(0), global.BadRequestErr
	}

	customer, err := bs.rep.SelectCustomerByID(ctx, id)
	if err != nil {
		log.Println("GetCustomerByID err:", err)
		return customer, global.InternalServerErr
	}

	if customer.ID == 0 {
		return customer, global.DataNotFoundErr
	}

	return customer, nil
}

func (bs *BasicService) StoreCustomer(ctx context.Context, customer domain.Customer) (uint64, error) {
	// если id не задан - запрос на создание
	if customer.ID == 0 {
		// не прошли валидацию, отправляемся обратно
		if !customer.ValidateFields(true) {
			return 0, global.IncorrectValidFormErr
		}

		id, err := bs.rep.CreateCustomer(ctx, customer)
		if err != nil || id == 0 {
			log.Println("StoreCustomer create err:", err)
			return 0, global.InternalServerErr
		}

		return id, nil
	}

	// проверка корректнности заполнения полей
	if !customer.ValidateFields(false) {
		return 0, global.IncorrectValidFormErr
	}

	id, err := bs.rep.UpdateCustomer(ctx, customer)
	if err != nil || id == 0 {
		log.Println("StoreCustomer update err:", err)
		return 0, global.InternalServerErr
	}

	return id, nil
}
