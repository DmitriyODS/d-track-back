package v1

type Auth struct {
	UserID   uint64 `json:"user_id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	JWT      string `json:"jwt"`
}
