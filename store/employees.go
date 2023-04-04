package store

import (
	"context"
	"database/sql"
	"fmt"
	"gitlab.com/ddda/d-track/d-track-back/domain"
	"gitlab.com/ddda/d-track/d-track-back/global"
)

const (
	selectEmployeesQuery = `
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
`
	selectEmployeeByIDQuery = `
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
WHERE e.id = $1
`
	createEmployeeQuery = `
INSERT INTO user_data.employees(fio,
                                "login",
                                hash_pass,
                                phone,
                                email,
                                address_of_residence,
                                position_id,
                                level_access_id,
                                freedom_type_id,
                                date_appointments,
                                date_of_dismissal)
VALUES ($1, $2, crypt($3, gen_salt('bf')), $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING id;
`
	updateEmployeeQuery = `
UPDATE user_data.employees
SET fio=$1,
    "login"=$2,
    phone=$3,
    email=$4,
    address_of_residence=$5,
    position_id=$6,
    level_access_id=$7,
    freedom_type_id=$8,
    date_of_dismissal=$9
WHERE id = $10
RETURNING id;
`
	updateEmployeeWithPassQuery = `
UPDATE user_data.employees
SET fio=$1,
    "login"=$2,
    hash_pass=crypt($3, gen_salt('bf')),
    phone=$4,
    email=$5,
    address_of_residence=$6,
    position_id=$7,
    level_access_id=$8,
    freedom_type_id=$9,
    date_of_dismissal=$10
WHERE id = $11
RETURNING id;
`
)

func selectEmployeePlaceholder(employee *domain.Employee) []interface{} {
	return []interface{}{
		&employee.ID,
		&employee.FIO,
		&employee.Login,
		&employee.PhoneNumber,
		&employee.EmailAddress,
		&employee.AddressOfResidence,
		&employee.Position.ID,
		&employee.Position.Value,
		&employee.LevelAccess.ID,
		&employee.LevelAccess.Name,
		&employee.LevelAccess.Access,
		&employee.FreedomType.ID,
		&employee.FreedomType.Value,
		&employee.DateAppointments,
		&employee.DateOfDismissal,
	}
}

func createEmployeePlaceholder(employee domain.Employee) []interface{} {
	return []interface{}{
		employee.FIO,
		employee.Login,
		employee.Password,
		employee.PhoneNumber,
		employee.EmailAddress,
		employee.AddressOfResidence,
		employee.Position.ID,
		employee.LevelAccess.ID,
		employee.FreedomType.ID,
		employee.DateAppointments,
		employee.DateOfDismissal,
	}
}
func updateEmployeePlaceholder(employee domain.Employee) []interface{} {
	return []interface{}{
		employee.FIO,
		employee.Login,
		employee.PhoneNumber,
		employee.EmailAddress,
		employee.AddressOfResidence,
		employee.Position.ID,
		employee.LevelAccess.ID,
		employee.FreedomType.ID,
		employee.DateOfDismissal,
		employee.ID,
	}
}

func updateWithPassEmployeePlaceholder(employee domain.Employee) []interface{} {
	return []interface{}{
		employee.FIO,
		employee.Login,
		employee.Password,
		employee.PhoneNumber,
		employee.EmailAddress,
		employee.AddressOfResidence,
		employee.Position.ID,
		employee.LevelAccess.ID,
		employee.FreedomType.ID,
		employee.DateOfDismissal,
		employee.ID,
	}
}

func (s *Store) SelectEmployees(ctx context.Context, fioFilter string, isArchive bool) ([]domain.Employee, error) {
	sqlWithFilters := ""
	if isArchive {
		sqlWithFilters = fmt.Sprintf("%s WHERE e.freedom_type_id=%d", selectEmployeesQuery,
			global.EmployeeFreedomTypeFired)
	} else {
		sqlWithFilters = fmt.Sprintf("%s WHERE e.freedom_type_id!=%d", selectEmployeesQuery,
			global.EmployeeFreedomTypeFired)
	}

	if fioFilter != "" {
		sqlWithFilters = fmt.Sprintf("%s AND e.fio ILIKE '%s'", sqlWithFilters, "%"+fioFilter+"%")
	}

	sqlWithFilters = fmt.Sprintf("%s ORDER by date_appointments DESC", sqlWithFilters)

	rows, err := s.Query(ctx, sqlWithFilters)
	if err == sql.ErrNoRows {
		return []domain.Employee{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employeesLst := make([]domain.Employee, 0)

	var employee domain.Employee
	for rows.Next() {
		if err = rows.Scan(selectEmployeePlaceholder(&employee)...); err != nil {
			return nil, err
		}

		employeesLst = append(employeesLst, employee)
	}

	return employeesLst, rows.Err()
}

func (s *Store) SelectEmployeeByID(ctx context.Context, id uint64) (domain.Employee, error) {
	employee := domain.NewEmployee(0)

	row, err := s.QueryRow(ctx, selectEmployeeByIDQuery, id)
	if err != nil {
		return employee, err
	}

	err = row.Scan(selectEmployeePlaceholder(&employee)...)
	if err == sql.ErrNoRows {
		return employee, nil
	}

	return employee, err
}

func (s *Store) CreateEmployee(ctx context.Context, employee domain.Employee) (uint64, error) {
	row, err := s.QueryRow(ctx, createEmployeeQuery, createEmployeePlaceholder(employee)...)
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

func (s *Store) UpdateEmployee(ctx context.Context, employee domain.Employee) (uint64, error) {
	var row *sql.Row
	var err error

	if employee.Password == "" {
		row, err = s.QueryRow(ctx, updateEmployeeQuery, updateEmployeePlaceholder(employee)...)
	} else {
		row, err = s.QueryRow(ctx, updateEmployeeWithPassQuery, updateWithPassEmployeePlaceholder(employee)...)
	}

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
