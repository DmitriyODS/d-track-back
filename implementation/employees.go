package implementation

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"log"
)

func (bs *BasicService) GetListEmployees(ctx context.Context, fioFilter string, isArchive bool) ([]domain.Employee, error) {
	employees, err := bs.rep.SelectEmployees(ctx, fioFilter, isArchive)
	if err != nil {
		log.Println("GetListEmployees err:", err)
		return nil, global.InternalServerErr
	}

	return employees, nil
}

func (bs *BasicService) GetEmployeeByID(ctx context.Context, id uint64) (domain.Employee, error) {
	if id == 0 {
		return domain.NewEmployee(0), global.BadRequestErr
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

	// если id администратора, то нельзя менять уровень доступа, роль и увольнять сотрудника
	if employee.ID == global.EmployeeAdminID {
		if employee.FreedomType.ID == global.EmployeeFreedomTypeFired {
			return 0, global.NotFiredAdminUser
		}

		if employee.LevelAccess.ID != global.EmployeeLevelAccessAdmin {
			return 0, global.NotEditLevelAccessAdminUser
		}

		if employee.Position.ID != global.EmployeePositionAdmin {
			return 0, global.NotEditRoleAdminUser
		}
	}

	// получаем данные текущего пользователя
	claims, ok := ctx.Value(global.JwtClaimsCtxKey).(*global.JwtClaims)
	if !ok {
		return 0, global.InternalServerErr
	}

	// если id сотрудника совпадает с id текущего пользователя, то нельзя менять уровень доступа, роль и увольнять сотрудника
	if employee.ID == claims.UserID {
		// получаем данные сотрудника до изменения
		oldEmployee, err := bs.rep.SelectEmployeeByID(ctx, employee.ID)
		if err != nil {
			log.Println("StoreEmployee select old employee err:", err)
			return 0, global.InternalServerErr
		}

		if employee.FreedomType.ID == global.EmployeeFreedomTypeFired {
			return 0, global.NotFiredYourselfUser
		}

		if employee.LevelAccess.ID != oldEmployee.LevelAccess.ID {
			return 0, global.NotEditLevelAccessYourselfUser
		}

		if employee.Position.ID != oldEmployee.Position.ID {
			return 0, global.NotEditRoleYourselfUser
		}
	}

	// проверка корректнности заполнения полей
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
