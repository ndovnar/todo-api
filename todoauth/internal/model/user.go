package model

type User struct {
	ID        string
	Email     string
	Password  string
	IsDeleted bool
	Dates     *Dates
}
