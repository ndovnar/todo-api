package dto

type Profile struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Dates     *Dates `json:"dates"`
}
