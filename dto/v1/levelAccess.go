package v1

type LevelAccess struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Access byte   `json:"access"`
}
