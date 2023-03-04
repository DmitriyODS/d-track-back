package v1

// RequestList - структура запроса списка чего - то от клиента
type RequestList struct {
	Filters map[string]string `json:"filters,omitempty"`
	Sorts   map[string]string `json:"sorts,omitempty"`
}

type RequestByID struct {
	ID uint64 `json:"id" uri:"id" binding:"required"`
}
