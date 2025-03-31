package model

type Session struct {
	ID        string
	UserID    string
	IsDeleted bool
	Dates     Dates
}
