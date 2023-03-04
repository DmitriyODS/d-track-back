package middleware

import (
	"gitlab.com/ddda/d-track/d-track-back/service"
)

// Logger - простой логгер основанный на log.Print
// для учебного проекта - норм, а так нужно писать нормальный
// логгер с уровнями и мета-инфой
type Logger struct {
	next service.Service
}

// Auth - аутентификатор, проверяет уровни доступа к сервису
// упрощён, тупо чекает JWT и привелегии из БД
type Auth struct {
	next service.Service
}

func AddAuthMiddleware(next service.Service) *Auth {
	return &Auth{next: next}
}

func AddLoggerMiddleware(next service.Service) *Logger {
	return &Logger{next: next}
}
