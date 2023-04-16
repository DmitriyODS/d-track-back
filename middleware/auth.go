package middleware

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"log"
)

func (a *Auth) checkLevelAccess(ctx context.Context, section global.Section, method global.Method) (context.Context, error) {
	// проверяем, а не открыта ли уже сессия
	if v := ctx.Value(global.CurSessionCtxKey); v != nil {
		return ctx, nil
	}

	// получаем мета-инфу о текущем пользователе
	claims, ok := ctx.Value(global.JwtClaimsCtxKey).(*global.JwtClaims)
	if !ok {
		return ctx, global.UnauthorizedErr
	}

	// проверим, вдруг привелегии обновились
	check, err := a.next.CheckLevelAccessByEmployeeID(ctx, claims.UserID, []byte{claims.LevelAccess})
	if err != nil {
		log.Println("GetListEmployees auth CheckLevelAccessByEmployeeID err:", err)
		return ctx, global.InternalServerErr
	}
	if !check {
		return ctx, global.AccessRightsObsoleteErr
	}

	// проверяем может ли пользователь осуществить действие по привелегиям
	if err = global.CheckLevelAccessWithClaims(claims.LevelAccess, section, method); err != nil {
		return ctx, global.ForbiddenErr
	}

	return context.WithValue(ctx, global.CurSessionCtxKey, true), nil
}

func (a *Auth) GetListEmployees(ctx context.Context, fioFilter string, isArchive bool) (resLst []domain.Employee, err error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionEmployees, global.MethodView)
	if err != nil {
		return nil, err
	}

	return a.next.GetListEmployees(authCtx, fioFilter, isArchive)
}

func (a *Auth) GetEmployeeByID(ctx context.Context, id uint64) (res domain.Employee, err error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionEmployees, global.MethodView)
	if err != nil {
		return domain.NewEmployee(0), err
	}

	return a.next.GetEmployeeByID(authCtx, id)
}

func (a *Auth) StoreEmployee(ctx context.Context, employee domain.Employee) (id uint64, err error) {
	curMethod := global.Method(global.MethodEdit)
	if employee.FreedomType.ID == global.EmployeeFreedomTypeFired {
		curMethod = global.MethodCreate
	}

	authCtx, err := a.checkLevelAccess(ctx, global.SectionEmployees, curMethod)
	if err != nil {
		return 0, err
	}

	return a.next.StoreEmployee(authCtx, employee)
}

// AuthenticationByLogin - не защищаем, зачем защищать защиту? o_O
func (a *Auth) AuthenticationByLogin(ctx context.Context, auth domain.Auth) (res domain.Auth, err error) {
	return a.next.AuthenticationByLogin(ctx, auth)
}

// CheckLevelAccessByEmployeeID - нужен лишь для системных операций, нет смысла чекать
func (a *Auth) CheckLevelAccessByEmployeeID(ctx context.Context, id uint64, levelAccess []byte) (bool, error) {
	return a.next.CheckLevelAccessByEmployeeID(ctx, id, levelAccess)
}

func (a *Auth) GetSelectListEmployees(ctx context.Context) (resLst []domain.SelectList, err error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionSelectLists, global.MethodView)
	if err != nil {
		return nil, err
	}

	return a.next.GetSelectListEmployees(authCtx)
}

func (a *Auth) GetSelectListPosition(ctx context.Context) (resLst []domain.SelectList, err error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionSelectLists, global.MethodView)
	if err != nil {
		return nil, err
	}

	return a.next.GetSelectListPosition(authCtx)
}

func (a *Auth) GetSelectListLevelAccesses(ctx context.Context) (resLst []domain.LevelAccess, err error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionSelectLists, global.MethodView)
	if err != nil {
		return nil, err
	}

	return a.next.GetSelectListLevelAccesses(authCtx)
}

func (a *Auth) GetSelectListFreedomType(ctx context.Context) (resLst []domain.SelectList, err error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionSelectLists, global.MethodView)
	if err != nil {
		return nil, err
	}

	return a.next.GetSelectListFreedomType(authCtx)
}

func (a *Auth) GetSelectListCustomers(ctx context.Context) ([]domain.SelectList, error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionSelectLists, global.MethodView)
	if err != nil {
		return nil, err
	}

	return a.next.GetSelectListCustomers(authCtx)
}

func (a *Auth) GetSelectListServices(ctx context.Context) ([]domain.SelectList, error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionSelectLists, global.MethodView)
	if err != nil {
		return nil, err
	}

	return a.next.GetSelectListServices(authCtx)
}

func (a *Auth) GetSelectListClaimStates(ctx context.Context) ([]domain.SelectList, error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionSelectLists, global.MethodView)
	if err != nil {
		return nil, err
	}

	return a.next.GetSelectListClaimStates(authCtx)
}

func (a *Auth) GetSelectListTaskStates(ctx context.Context) ([]domain.SelectList, error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionSelectLists, global.MethodView)
	if err != nil {
		return nil, err
	}

	return a.next.GetSelectListTaskStates(authCtx)
}

func (a *Auth) GetListClaims(ctx context.Context, fioFilter string, isArchive bool, customerID uint64) ([]domain.Claim, error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionClaims, global.MethodView)
	if err != nil {
		return nil, err
	}

	return a.next.GetListClaims(authCtx, fioFilter, isArchive, customerID)
}

func (a *Auth) GetClaimByID(ctx context.Context, id uint64) (domain.Claim, error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionClaims, global.MethodView)
	if err != nil {
		return domain.NewClaim(0), err
	}

	return a.next.GetClaimByID(authCtx, id)
}

func (a *Auth) StoreClaim(ctx context.Context, claim domain.Claim) (uint64, error) {
	curMethod := global.Method(global.MethodEdit)

	authCtx, err := a.checkLevelAccess(ctx, global.SectionClaims, curMethod)
	if err != nil {
		return 0, err
	}

	return a.next.StoreClaim(authCtx, claim)
}

func (a *Auth) GetListTasks(ctx context.Context, numberFilter string, isArchive bool) ([]domain.Task, error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionTasks, global.MethodView)
	if err != nil {
		return nil, err
	}

	return a.next.GetListTasks(authCtx, numberFilter, isArchive)
}

func (a *Auth) GetTaskByID(ctx context.Context, id uint64) (domain.Task, error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionTasks, global.MethodView)
	if err != nil {
		return domain.NewTask(0), err
	}

	return a.next.GetTaskByID(authCtx, id)
}

func (a *Auth) StoreTask(ctx context.Context, task domain.Task) (uint64, error) {
	curMethod := global.Method(global.MethodEdit)

	authCtx, err := a.checkLevelAccess(ctx, global.SectionTasks, curMethod)
	if err != nil {
		return 0, err
	}

	return a.next.StoreTask(authCtx, task)
}

func (a *Auth) GetListCustomers(ctx context.Context, fioFilter string, isArchive bool, claimID uint64) ([]domain.Customer, error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionCustomers, global.MethodView)
	if err != nil {
		return nil, err
	}

	return a.next.GetListCustomers(authCtx, fioFilter, isArchive, claimID)
}

func (a *Auth) GetCustomerByID(ctx context.Context, id uint64) (domain.Customer, error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionCustomers, global.MethodView)
	if err != nil {
		return domain.NewCustomer(0), err
	}

	return a.next.GetCustomerByID(authCtx, id)
}

func (a *Auth) StoreCustomer(ctx context.Context, customer domain.Customer) (uint64, error) {
	curMethod := global.Method(global.MethodEdit)

	authCtx, err := a.checkLevelAccess(ctx, global.SectionCustomers, curMethod)
	if err != nil {
		return 0, err
	}

	return a.next.StoreCustomer(authCtx, customer)
}
