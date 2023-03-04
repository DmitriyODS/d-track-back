package store

import (
	"context"
	"database/sql"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

const (
	selectEmployeeByLoginPassQuery = `
SELECT e.id,
       e.fio,
       e."login",
       e.phone,
       e.email,
       e.address_of_residence,
       e.position_id,
       p.name    AS position_name,
       e.level_access_id,
       la.name   AS level_access_name,
       la.access AS level_access,
       e.freedom_type_id,
       ft.name   AS freedom_type_name,
       e.date_appointments,
       e.date_of_dismissal
FROM user_data.employees AS e
         INNER JOIN user_data.positions AS p on p.id = e.position_id
         INNER JOIN user_data.level_accesses AS la on la.id = e.level_access_id
         INNER JOIN user_data.freedom_types AS ft on ft.id = e.freedom_type_id
WHERE e."login" = $1
  AND e.hash_pass = crypt($2, e.hash_pass);
`
	selectCheckLevelAccessByEmployeeIDQuery = `
SELECT la.access = $2 AS check_level_access
FROM user_data.employees AS e
         INNER JOIN user_data.level_accesses AS la on la.id = e.level_access_id
WHERE e.id = $1;
`
)

func (s *Store) SelectUserByLoginPass(ctx context.Context, auth domain.Auth) (domain.Employee, error) {
	employee := domain.NewEmployee(0)

	row, err := s.QueryRow(ctx, selectEmployeeByLoginPassQuery, auth.Login, auth.Password)
	if err != nil {
		return employee, err
	}

	err = row.Scan(selectEmployeePlaceholder(&employee)...)
	if err == sql.ErrNoRows {
		return employee, nil
	}

	return employee, err
}

func (s *Store) CheckLevelAccessByEmployeeID(ctx context.Context, id uint64, levelAccess []byte) (bool, error) {
	row, err := s.QueryRow(ctx, selectCheckLevelAccessByEmployeeIDQuery, id, levelAccess)
	if err != nil {
		return false, err
	}

	var isCheck bool
	err = row.Scan(&isCheck)
	if err != nil {
		return false, err
	}

	return isCheck, nil
}
