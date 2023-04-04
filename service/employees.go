package service

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

type employees interface {
	GetListEmployees(ctx context.Context, fioFilter string, isArchive bool) ([]domain.Employee, error)
	GetEmployeeByID(ctx context.Context, id uint64) (domain.Employee, error)
	StoreEmployee(ctx context.Context, employee domain.Employee) (uint64, error)
}
