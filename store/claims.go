package store

import (
	"context"
	"database/sql"
	"fmt"
	"gitlab.com/ddda/d-track/d-track-back/domain"
	"gitlab.com/ddda/d-track/d-track-back/global"
)

const (
	selectClaimsQuery = `
SELECT claim.id,
       claim.number,
       claim.date_created,
       claim.date_completed,
       claim.date_estimated_completion,
       customer.id,
       customer.fio,
       claim.subject,
       service.id,
       service.name,
       claim.description,
       cs.id,
       cs.name,
       employee.id,
       employee.fio
FROM user_data.claims AS claim
         INNER JOIN user_data.customers AS customer on customer.id = claim.customer_id
         INNER JOIN user_data.employees AS employee on employee.id = claim.executor_id
         INNER JOIN user_data.services AS service on service.id = claim.service_type_id
         INNER JOIN user_data.claim_states AS cs on cs.id = claim.state_id
`
	selectClaimByIDQuery = `
SELECT claim.id,
       claim.number,
       claim.date_created,
       claim.date_completed,
       claim.date_estimated_completion,
       customer.id,
       customer.fio,
       claim.subject,
       service.id,
       service.name,
       claim.description,
       cs.id,
       cs.name,
       employee.id,
       employee.fio
FROM user_data.claims AS claim
         INNER JOIN user_data.customers AS customer on customer.id = claim.customer_id
         INNER JOIN user_data.employees AS employee on employee.id = claim.executor_id
         INNER JOIN user_data.services AS service on service.id = claim.service_type_id
         INNER JOIN user_data.claim_states AS cs on cs.id = claim.state_id
WHERE claim.id = $1
`
	createClaimQuery = `
INSERT INTO user_data.claims(number,
                             customer_id,
                             subject,
                             description,
                             date_created,
                             date_estimated_completion,
                             service_type_id,
                             state_id,
                             executor_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id;
`
	updateClaimQuery = `
UPDATE user_data.claims
SET date_completed=$1,
    date_estimated_completion=$2,
    description=$3,
    state_id=$4,
    executor_id=$5
WHERE id = $6
RETURNING id;
`
)

func selectClaimPlaceholder(claim *domain.Claim) []interface{} {
	return []interface{}{
		&claim.ID,
		&claim.Number,
		&claim.DateCreated,
		&claim.DateCompleted,
		&claim.DateEstimatedCompletion,
		&claim.Customer.ID,
		&claim.Customer.FIO,
		&claim.Subject,
		&claim.ServiceType.ID,
		&claim.ServiceType.Value,
		&claim.Description,
		&claim.Status.ID,
		&claim.Status.Value,
		&claim.Executor.ID,
		&claim.Executor.FIO,
	}
}

func createClaimPlaceholder(claim domain.Claim) []interface{} {
	return []interface{}{
		claim.Number,
		claim.Customer.ID,
		claim.Subject,
		claim.Description,
		claim.DateCreated,
		claim.DateEstimatedCompletion,
		claim.ServiceType.ID,
		claim.Status.ID,
		claim.Executor.ID,
	}
}
func updateClaimPlaceholder(claim domain.Claim) []interface{} {
	return []interface{}{
		claim.DateCompleted,
		claim.DateEstimatedCompletion,
		claim.Description,
		claim.Status.ID,
		claim.Executor.ID,
		claim.ID,
	}
}

func (s *Store) SelectClaims(ctx context.Context, numberFilter string, isArchive bool) ([]domain.Claim, error) {
	sqlWithFilters := ""
	if isArchive {
		sqlWithFilters = fmt.Sprintf("%s WHERE claim.state_id=%d", selectClaimsQuery, global.ClaimStateClose)
	} else {
		sqlWithFilters = fmt.Sprintf("%s WHERE claim.state_id!=%d", selectClaimsQuery, global.ClaimStateClose)
	}

	if numberFilter != "" {
		sqlWithFilters = fmt.Sprintf("%s AND claim.number ILIKE '%s'", sqlWithFilters, "%"+numberFilter+"%")
	}

	sqlWithFilters = fmt.Sprintf("%s ORDER by claim.date_created DESC", sqlWithFilters)

	rows, err := s.Query(ctx, sqlWithFilters)
	if err == sql.ErrNoRows {
		return []domain.Claim{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	claimsLst := make([]domain.Claim, 0)

	var claim domain.Claim
	for rows.Next() {
		if err = rows.Scan(selectClaimPlaceholder(&claim)...); err != nil {
			return nil, err
		}

		claimsLst = append(claimsLst, claim)
	}

	return claimsLst, rows.Err()
}

func (s *Store) SelectClaimByID(ctx context.Context, id uint64) (domain.Claim, error) {
	claim := domain.NewClaim(0)

	row, err := s.QueryRow(ctx, selectClaimByIDQuery, id)
	if err != nil {
		return claim, err
	}

	err = row.Scan(selectClaimPlaceholder(&claim)...)
	if err == sql.ErrNoRows {
		return claim, nil
	}

	return claim, err
}

func (s *Store) CreateClaim(ctx context.Context, claim domain.Claim) (uint64, error) {
	row, err := s.QueryRow(ctx, createClaimQuery, createClaimPlaceholder(claim)...)
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

func (s *Store) UpdateClaim(ctx context.Context, claim domain.Claim) (uint64, error) {
	var row *sql.Row
	var err error

	row, err = s.QueryRow(ctx, updateClaimQuery, updateClaimPlaceholder(claim)...)
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
