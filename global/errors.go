package global

import (
	"errors"
	"net/http"
)

var (
	BadRequestErr       = errors.New("запрос является некорретным")
	UnauthorizedErr     = errors.New("вы не авторизированы")
	ForbiddenErr        = errors.New("недостаточно прав доступа")
	NotFoundErr         = errors.New("запрашиваемый API не найден")
	MethodNotAllowedErr = errors.New("метод временно заблокирован")

	InternalServerErr = errors.New("системная ошибка, попробуйте ещё раз")
	NotImplementedErr = errors.New("запрашиваемый метод не поддерживается")

	IncorrectValidFormErr       = errors.New("поля заполнены некорректно")
	IncorrectLoginOrPassErr     = errors.New("неверный логин, или пароль")
	EmployeeIsDismissalErr      = errors.New("сотрудник уволен")
	EmployeeDateAppointmentsErr = errors.New("дата назначения ещё не наступила")
	IncorrectUpdateUserData     = errors.New("невозможно изменить текущего пользователя")
	AccessRightsObsoleteErr     = errors.New("права доступа устарели, перезайдите в сервис")
	IncorrectBodyRequestErr     = errors.New("тело запроса некорректно")
	DataNotFoundErr             = errors.New("запрашиваемые данные не найдены")
)

func GetStatusCodeByErr(err error) int {
	switch err {
	case BadRequestErr, IncorrectValidFormErr, IncorrectBodyRequestErr, DataNotFoundErr, IncorrectUpdateUserData:
		return http.StatusBadRequest
	case UnauthorizedErr, IncorrectLoginOrPassErr, EmployeeIsDismissalErr, EmployeeDateAppointmentsErr:
		return http.StatusUnauthorized
	case ForbiddenErr, AccessRightsObsoleteErr:
		return http.StatusForbidden
	case NotFoundErr:
		return http.StatusNotFound
	case MethodNotAllowedErr:
		return http.StatusMethodNotAllowed
	case InternalServerErr:
		return http.StatusInternalServerError
	case NotImplementedErr:
		return http.StatusNotImplemented
	}

	return 0
}
