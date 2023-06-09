package implementation

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"log"
)

func (bs *BasicService) GetListEmployees(ctx context.Context, filters, sorts map[string]string) ([]domain.Employee, error) {
	employees, err := bs.rep.SelectEmployees(ctx, filters, sorts)
	if err != nil {
		log.Println("GetListEmployees err:", err)
		return nil, global.InternalServerErr
	}

	return employees, nil
}

func (bs *BasicService) GetEmployeeByID(ctx context.Context, id uint64) (domain.Employee, error) {
	if id == 0 {
		return domain.Employee{}, global.BadRequestErr
	}

	employee, err := bs.rep.SelectEmployeeByID(ctx, id)
	if err != nil {
		log.Println("GetEmployeeByID err:", err)
		return employee, global.InternalServerErr
	}

	if employee.ID == 0 {
		return employee, global.DataNotFoundErr
	}

	return employee, nil
}

func (bs *BasicService) StoreEmployee(ctx context.Context, employee domain.Employee) (uint64, error) {
	// если id не задан - запрос на создание
	if employee.ID == 0 {
		// не прошли валидацию, отправляемся обратно
		if !employee.ValidateFields(true) {
			return 0, global.IncorrectValidFormErr
		}

		id, err := bs.rep.CreateEmployee(ctx, employee)
		if err != nil || id == 0 {
			log.Println("StoreEmployee create err:", err)
			return 0, global.InternalServerErr
		}

		return id, nil
	}

	// обновление
	if !employee.ValidateFields(false) {
		return 0, global.IncorrectValidFormErr
	}

	id, err := bs.rep.UpdateEmployee(ctx, employee)
	if err != nil || id == 0 {
		log.Println("StoreEmployee update err:", err)
		return 0, global.InternalServerErr
	}

	return id, nil
}
