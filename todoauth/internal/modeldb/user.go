package modeldb

type User struct {
	ID        string `bson:"_id,omitempty"`
	Email     string `bson:"email"`
	Password  string `bson:"password"`
	IsDeleted bool   `bson:"deleted"`
	Dates     *Dates `bson:"dates"`
}
