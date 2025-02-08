package model

type Todo struct {
	ID          string `bson:"_id,omitempty"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	UserID      string `bson:"userId"`
	IsDeleted   bool   `bson:"deleted"`
	IsCompleted bool   `bson:"completed"`
	Dates       *Dates `bson:"dates"`
}
