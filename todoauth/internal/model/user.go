package model

type User struct {
	ID        string `json:"id,omitempty" bson:"_id,omitempty"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"-" bson:"password"`
	IsDeleted bool   `json:"deleted,omitempty" bson:"deleted"`
	Dates     Dates  `json:"dates" bson:"dates"`
}
