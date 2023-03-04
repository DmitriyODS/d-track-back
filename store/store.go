package store

import (
	"gitlab.com/ddda/d-track/d-track-back/global"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Store - реализация сервиса работы с БД
type Store struct {
	global.DBController
}

func NewStore(strConnectDB string) *Store {
	return &Store{DBController: global.NewDBController(strConnectDB)}
}
