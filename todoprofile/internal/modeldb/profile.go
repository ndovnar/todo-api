package modeldb

type Profile struct {
	ID        string `bson:"_id,omitempty"`
	UserID    string `bson:"userId"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	IsDeleted bool   `bson:"deleted"`
	Dates     *Dates `bson:"dates"`
}
