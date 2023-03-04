package repository

// Repository - интерфейс взаимодействия с глобальным хранилищем.
// С помощью него мы можем получать, обновлять, вставлять и удалять записи.
// Сюда встроены интерфейсы каждого репозитория доменной сущности описанные в других файлах.
type Repository interface {
	repositoryEmployees
	repositorySelectLists
	auth
}
