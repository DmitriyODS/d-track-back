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
	check, err := a.next.CheckLevelAccessByEmployeeID(ctx, claims.UserID, claims.LevelAccess)
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

func (a *Auth) GetListEmployees(ctx context.Context, filters, sorts map[string]string) (resLst []domain.Employee, err error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionEmployees, global.MethodView)
	if err != nil {
		return nil, err
	}

	return a.next.GetListEmployees(authCtx, filters, sorts)
}

func (a *Auth) GetEmployeeByID(ctx context.Context, id uint64) (res domain.Employee, err error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionEmployees, global.MethodView)
	if err != nil {
		return domain.Employee{}, err
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

func (a *Auth) GetSelectListEmployees(ctx context.Context, filters, sorts map[string]string) (resLst []domain.SelectList, err error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionSelectLists, global.MethodView)
	if err != nil {
		return nil, err
	}

	return a.next.GetSelectListEmployees(authCtx, filters, sorts)
}

func (a *Auth) GetSelectListPosition(ctx context.Context, filters, sorts map[string]string) (resLst []domain.SelectList, err error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionSelectLists, global.MethodView)
	if err != nil {
		return nil, err
	}

	return a.next.GetSelectListPosition(authCtx, filters, sorts)
}

func (a *Auth) GetSelectListLevelAccesses(ctx context.Context, filters, sorts map[string]string) (resLst []domain.LevelAccess, err error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionSelectLists, global.MethodView)
	if err != nil {
		return nil, err
	}

	return a.next.GetSelectListLevelAccesses(authCtx, filters, sorts)
}

func (a *Auth) GetSelectListFreedomType(ctx context.Context, filters, sorts map[string]string) (resLst []domain.SelectList, err error) {
	authCtx, err := a.checkLevelAccess(ctx, global.SectionSelectLists, global.MethodView)
	if err != nil {
		return nil, err
	}

	return a.next.GetSelectListFreedomType(authCtx, filters, sorts)
}
