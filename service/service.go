package service

// Service - интерфейс взаимодействия с методами сервиса
// Встраивает в себя интерфейсы каждой секции из пакета: service
type Service interface {
	employees
	selectLists
	auth
}
