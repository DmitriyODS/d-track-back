package v1

type RequestCustomerListFilters struct {
	IsArchive bool   `form:"is_archive"`
	FioFilter string `form:"fio_filter"`
}

type Customer struct {
	ID          uint64 `json:"id"`
	FIO         string `json:"fio"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	DateCreated int64  `json:"date_created"`
}
