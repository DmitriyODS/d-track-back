package v1

type RequestByID struct {
	ID uint64 `json:"id" uri:"id" binding:"required"`
}
