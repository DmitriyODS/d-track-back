package store

import (
	"context"
	"database/sql"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

const (
	selectSelectListEmployeesQuery = `
SELECT e.id,
       format('%s - %s', e.fio, p.name) AS "value"
FROM user_data.employees AS e
         INNER JOIN user_data.positions AS p on p.id = e.position_id
`
	selectSelectListCustomersQuery = `
SELECT c.id,
       c.fio
FROM user_data.customers AS c
ORDER BY c.fio;
`
	selectSelectListPositionsQuery = `
SELECT id, "name"
FROM user_data.positions
`
	selectLevelAccessesQuery = `
SELECT id, "name", "access"
FROM user_data.level_accesses
`
	selectSelectListFreedomTypesQuery = `
SELECT id, "name"
FROM user_data.freedom_types
`
	selectSelectListServicesQuery = `
SELECT id, "name"
FROM user_data.services;
`
	selectSelectListTaskStatesQuery = `
SELECT id, "name"
FROM user_data.task_states;
`
	selectSelectListClaimStatesQuery = `
SELECT id, "name"
FROM user_data.claim_states;
`
)

func selectSelectListPlaceholder(selectList *domain.SelectList) []interface{} {
	return []interface{}{
		&selectList.ID,
		&selectList.Value,
	}
}

func selectLevelAccessPlaceholder(la *domain.LevelAccess) []interface{} {
	return []interface{}{
		&la.ID,
		&la.Name,
		&la.Access,
	}
}

func (s *Store) selectSelectList(ctx context.Context, query string) ([]domain.SelectList, error) {
	rows, err := s.Query(ctx, query)
	if err == sql.ErrNoRows {
		return []domain.SelectList{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	selectLists := make([]domain.SelectList, 0)

	var selectList domain.SelectList
	for rows.Next() {
		if err = rows.Scan(selectSelectListPlaceholder(&selectList)...); err != nil {
			return nil, err
		}

		selectLists = append(selectLists, selectList)
	}

	return selectLists, rows.Err()
}

func (s *Store) SelectSelectListEmployees(ctx context.Context) ([]domain.SelectList, error) {
	return s.selectSelectList(ctx, selectSelectListEmployeesQuery)
}

func (s *Store) SelectSelectListCustomers(ctx context.Context) ([]domain.SelectList, error) {
	return s.selectSelectList(ctx, selectSelectListCustomersQuery)
}

func (s *Store) SelectSelectListPositions(ctx context.Context) ([]domain.SelectList, error) {
	return s.selectSelectList(ctx, selectSelectListPositionsQuery)
}

func (s *Store) SelectSelectListLevelAccesses(ctx context.Context) ([]domain.LevelAccess, error) {
	rows, err := s.Query(ctx, selectLevelAccessesQuery)
	if err == sql.ErrNoRows {
		return []domain.LevelAccess{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	levelAccesses := make([]domain.LevelAccess, 0)

	var levelAccess domain.LevelAccess
	for rows.Next() {
		if err = rows.Scan(selectLevelAccessPlaceholder(&levelAccess)...); err != nil {
			return nil, err
		}

		levelAccesses = append(levelAccesses, levelAccess)
	}

	return levelAccesses, rows.Err()
}

func (s *Store) SelectSelectListFreedomTypes(ctx context.Context) ([]domain.SelectList, error) {
	return s.selectSelectList(ctx, selectSelectListFreedomTypesQuery)
}

func (s *Store) SelectSelectListClaimStates(ctx context.Context) ([]domain.SelectList, error) {
	return s.selectSelectList(ctx, selectSelectListClaimStatesQuery)
}

func (s *Store) SelectSelectListTaskStates(ctx context.Context) ([]domain.SelectList, error) {
	return s.selectSelectList(ctx, selectSelectListTaskStatesQuery)
}

func (s *Store) SelectSelectListServices(ctx context.Context) ([]domain.SelectList, error) {
	return s.selectSelectList(ctx, selectSelectListServicesQuery)
}
