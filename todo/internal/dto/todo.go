package dto

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	UserID      string `json:"userId"`
	Description string `json:"description"`
	IsCompleted bool   `json:"completed"`
	Dates       *Dates `json:"dates"`
}
