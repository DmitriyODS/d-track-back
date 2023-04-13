package middleware

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
	"log"
)

func (l *Logger) GetListEmployees(ctx context.Context, fioFilter string, isArchive bool) (resLst []domain.Employee, err error) {
	log.Println("Start GetListEmployees")
	defer func() {
		if err != nil {
			log.Println("End GetListEmployees with err:", err)
			return
		}

		log.Println("End GetListEmployees")
	}()

	return l.next.GetListEmployees(ctx, fioFilter, isArchive)
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

func (l *Logger) GetSelectListEmployees(ctx context.Context) (resLst []domain.SelectList, err error) {
	log.Println("Start GetSelectListEmployees")
	defer func() {
		if err != nil {
			log.Println("End GetSelectListEmployees with err:", err)
			return
		}

		log.Println("End GetSelectListEmployees")
	}()

	return l.next.GetSelectListEmployees(ctx)
}

func (l *Logger) GetSelectListPosition(ctx context.Context) (resLst []domain.SelectList, err error) {
	log.Println("Start GetSelectListPosition")
	defer func() {
		if err != nil {
			log.Println("End GetSelectListPosition with err:", err)
			return
		}

		log.Println("End GetSelectListPosition")
	}()

	return l.next.GetSelectListPosition(ctx)
}

func (l *Logger) GetSelectListLevelAccesses(ctx context.Context) (resLst []domain.LevelAccess, err error) {
	log.Println("Start GetSelectListLevelAccesses")
	defer func() {
		if err != nil {
			log.Println("End GetSelectListLevelAccesses with err:", err)
			return
		}

		log.Println("End GetSelectListLevelAccesses")
	}()

	return l.next.GetSelectListLevelAccesses(ctx)
}

func (l *Logger) GetSelectListFreedomType(ctx context.Context) (resLst []domain.SelectList, err error) {
	log.Println("Start GetSelectListFreedomType")
	defer func() {
		if err != nil {
			log.Println("End GetSelectListFreedomType with err:", err)
			return
		}

		log.Println("End GetSelectListFreedomType")
	}()

	return l.next.GetSelectListFreedomType(ctx)
}

func (l *Logger) GetListClaims(ctx context.Context, fioFilter string, isArchive bool) (claims []domain.Claim, err error) {
	log.Println("Start GetListClaims")
	defer func() {
		if err != nil {
			log.Println("End GetListClaims with err:", err)
			return
		}

		log.Println("End GetListClaims")
	}()

	return l.next.GetListClaims(ctx, fioFilter, isArchive)
}

func (l *Logger) GetClaimByID(ctx context.Context, id uint64) (claim domain.Claim, err error) {
	log.Printf("Start GetClaimByID with id=%d\n", id)
	defer func() {
		if err != nil {
			log.Println("End GetClaimByID with err:", err)
			return
		}

		log.Println("End GetClaimByID")
	}()

	return l.next.GetClaimByID(ctx, id)
}

func (l *Logger) StoreClaim(ctx context.Context, claim domain.Claim) (id uint64, err error) {
	log.Println("Start StoreClaim with employee=", claim)
	defer func() {
		if err != nil {
			log.Println("End StoreClaim with err:", err)
			return
		}

		log.Printf("End StoreClaim store id=%d\n", id)
	}()

	return l.next.StoreClaim(ctx, claim)
}
