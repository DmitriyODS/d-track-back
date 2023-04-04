package repository

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

type repositoryEmployees interface {
	SelectEmployees(ctx context.Context, fioFilter string, isArchive bool) ([]domain.Employee, error)
	SelectEmployeeByID(ctx context.Context, id uint64) (domain.Employee, error)
	CreateEmployee(ctx context.Context, employee domain.Employee) (uint64, error)
	UpdateEmployee(ctx context.Context, employee domain.Employee) (uint64, error)
}
