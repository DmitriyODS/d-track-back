package store

import (
	"context"
	"database/sql"
	"fmt"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

const (
	selectCustomersQuery = `
SELECT id,
       fio,
       phone,
       email,
       address,
       date_created
FROM user_data.customers AS c
`
	selectCustomerByIDQuery = `
SELECT id,
       fio,
       phone,
       email,
       address,
       date_created
FROM user_data.customers
WHERE id = $1;
`
	createCustomerQuery = `
INSERT INTO user_data.customers(fio,
                                phone,
                                email,
                                address,
                                date_created)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;
`
	updateCustomerQuery = `
UPDATE user_data.customers
SET fio=$1,
    phone=$2,
    email=$3,
    address=$4
WHERE id = $5
RETURNING id;
`
)

func selectCustomerPlaceholder(customer *domain.Customer) []interface{} {
	return []interface{}{
		&customer.ID,
		&customer.FIO,
		&customer.Phone,
		&customer.Email,
		&customer.Address,
		&customer.DateCreated,
	}
}

func createCustomerPlaceholder(customer domain.Customer) []interface{} {
	return []interface{}{
		customer.FIO,
		customer.Phone,
		customer.Email,
		customer.Address,
		customer.DateCreated,
	}
}

func updateCustomerPlaceholder(customer domain.Customer) []interface{} {
	return []interface{}{
		customer.FIO,
		customer.Phone,
		customer.Email,
		customer.Address,
		customer.ID,
	}
}

func (s *Store) SelectCustomers(ctx context.Context, fioFilter string, _ bool, claimID uint64) ([]domain.Customer, error) {
	sqlWithFilters := selectCustomersQuery

	if fioFilter != "" {
		sqlWithFilters = fmt.Sprintf("%s WHERE c.fio ILIKE '%s'", sqlWithFilters, "%"+fioFilter+"%")
	}

	if claimID != 0 {
		if fioFilter != "" {
			sqlWithFilters = fmt.Sprintf("%s AND id IN (SELECT customer_id FROM user_data.claims AS cl WHERE cl.id=%d)", sqlWithFilters, claimID)
		} else {
			sqlWithFilters = fmt.Sprintf("%s WHERE id IN (SELECT customer_id FROM user_data.claims AS cl WHERE cl.id=%d)", sqlWithFilters, claimID)
		}
	}

	sqlWithFilters = fmt.Sprintf("%s ORDER by date_created DESC", sqlWithFilters)

	rows, err := s.Query(ctx, sqlWithFilters)
	if err == sql.ErrNoRows {
		return []domain.Customer{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customersLst := make([]domain.Customer, 0)

	var customer domain.Customer
	for rows.Next() {
		if err = rows.Scan(selectCustomerPlaceholder(&customer)...); err != nil {
			return nil, err
		}

		customersLst = append(customersLst, customer)
	}

	return customersLst, rows.Err()
}

func (s *Store) SelectCustomerByID(ctx context.Context, id uint64) (domain.Customer, error) {
	customer := domain.NewCustomer(0)

	row, err := s.QueryRow(ctx, selectCustomerByIDQuery, id)
	if err != nil {
		return customer, err
	}

	err = row.Scan(selectCustomerPlaceholder(&customer)...)
	if err == sql.ErrNoRows {
		return customer, nil
	}

	return customer, err
}

func (s *Store) CreateCustomer(ctx context.Context, customer domain.Customer) (uint64, error) {
	row, err := s.QueryRow(ctx, createCustomerQuery, createCustomerPlaceholder(customer)...)
	if err != nil {
		return 0, err
	}

	var resId uint64
	err = row.Scan(&resId)
	if err == sql.ErrNoRows {
		return 0, nil
	}

	return resId, err
}

func (s *Store) UpdateCustomer(ctx context.Context, customer domain.Customer) (uint64, error) {
	var row *sql.Row
	var err error

	row, err = s.QueryRow(ctx, updateCustomerQuery, updateCustomerPlaceholder(customer)...)
	if err != nil {
		return 0, err
	}

	var resId uint64
	err = row.Scan(&resId)
	if err == sql.ErrNoRows {
		return 0, nil
	}

	return resId, err
}
