package domain

type Auth struct {
	UserID   uint64
	Login    string
	Password string
	JWT      string
}

func (a Auth) ValidateFields() bool {
	if a.Login == "" {
		return false
	}

	if a.Password == "" {
		return false
	}

	return true
}
