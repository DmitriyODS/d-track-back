package middleware

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
	"log"
)

func (l *Logger) GetListEmployees(ctx context.Context, filters, sorts map[string]string) (resLst []domain.Employee, err error) {
	log.Println("Start GetListEmployees")
	defer func() {
		if err != nil {
			log.Println("End GetListEmployees with err:", err)
			return
		}

		log.Println("End GetListEmployees")
	}()

	return l.next.GetListEmployees(ctx, filters, sorts)
}

func (l *Logger) GetEmployeeByID(ctx context.Context, id uint64) (res domain.Employee, err error) {
	log.Printf("Start GetEmployeeByID with id=%d\n", id)
	defer func() {
		if err != nil {
			log.Println("End GetEmployeeByID with err:", err)
			return
		}

		log.Println("End GetEmployeeByID")
	}()

	return l.next.GetEmployeeByID(ctx, id)
}

func (l *Logger) StoreEmployee(ctx context.Context, employee domain.Employee) (id uint64, err error) {
	log.Println("Start StoreEmployee with employee=", employee)
	defer func() {
		if err != nil {
			log.Println("End StoreEmployee with err:", err)
			return
		}

		log.Printf("End StoreEmployee store id=%d\n", id)
	}()

	return l.next.StoreEmployee(ctx, employee)
}

func (l *Logger) AuthenticationByLogin(ctx context.Context, auth domain.Auth) (res domain.Auth, err error) {
	log.Printf("Start AuthenticationByLogin with login=%s\n", auth.Login)
	defer func() {
		if err != nil {
			log.Println("End AuthenticationByLogin with err:", err)
			return
		}

		log.Printf("End AuthenticationByLogin with user id=%d\n", res.UserID)
	}()

	return l.next.AuthenticationByLogin(ctx, auth)
}

// CheckLevelAccessByEmployeeID - не логгируем, системный вызов
func (l *Logger) CheckLevelAccessByEmployeeID(ctx context.Context, id uint64, levelAccess []byte) (bool, error) {
	return l.next.CheckLevelAccessByEmployeeID(ctx, id, levelAccess)
}

func (l *Logger) GetSelectListEmployees(ctx context.Context, filters, sorts map[string]string) (resLst []domain.SelectList, err error) {
	log.Println("Start GetSelectListEmployees")
	defer func() {
		if err != nil {
			log.Println("End GetSelectListEmployees with err:", err)
			return
		}

		log.Println("End GetSelectListEmployees")
	}()

	return l.next.GetSelectListEmployees(ctx, filters, sorts)
}

func (l *Logger) GetSelectListPosition(ctx context.Context, filters, sorts map[string]string) (resLst []domain.SelectList, err error) {
	log.Println("Start GetSelectListPosition")
	defer func() {
		if err != nil {
			log.Println("End GetSelectListPosition with err:", err)
			return
		}

		log.Println("End GetSelectListPosition")
	}()

	return l.next.GetSelectListPosition(ctx, filters, sorts)
}

func (l *Logger) GetSelectListLevelAccesses(ctx context.Context, filters, sorts map[string]string) (resLst []domain.LevelAccess, err error) {
	log.Println("Start GetSelectListLevelAccesses")
	defer func() {
		if err != nil {
			log.Println("End GetSelectListLevelAccesses with err:", err)
			return
		}

		log.Println("End GetSelectListLevelAccesses")
	}()

	return l.next.GetSelectListLevelAccesses(ctx, filters, sorts)
}

func (l *Logger) GetSelectListFreedomType(ctx context.Context, filters, sorts map[string]string) (resLst []domain.SelectList, err error) {
	log.Println("Start GetSelectListFreedomType")
	defer func() {
		if err != nil {
			log.Println("End GetSelectListFreedomType with err:", err)
			return
		}

		log.Println("End GetSelectListFreedomType")
	}()

	return l.next.GetSelectListFreedomType(ctx, filters, sorts)
}
