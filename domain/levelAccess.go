package domain

type LevelAccess struct {
	ID     uint64
	Name   string
	Access []byte
}
