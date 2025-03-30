package dto

type UserTodoStatisticResponse struct {
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
	Total int64 `json:"total"`
}
