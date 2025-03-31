package model

type Todo struct {
	ID          string
	Title       string
	Description string
	UserID      string
	IsDeleted   bool
	IsCompleted bool
	Dates       *Dates
}
