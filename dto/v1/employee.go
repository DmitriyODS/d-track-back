package v1

type Employee struct {
	ID                 uint64      `json:"id"`
	FIO                string      `json:"fio"`
	Login              string      `json:"login"`
	Password           string      `json:"password"`
	PhoneNumber        string      `json:"phone_number"`
	EmailAddress       string      `json:"email_address"`
	AddressOfResidence string      `json:"address_of_residence"`
	Position           SelectList  `json:"position"`
	LevelAccess        LevelAccess `json:"level_access"`
	FreedomType        SelectList  `json:"freedom_type"`
	DateAppointments   int64       `json:"date_appointments"`
	DateOfDismissal    int64       `json:"date_of_dismissal"`
}
