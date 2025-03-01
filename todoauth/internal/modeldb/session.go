package modeldb

type Session struct {
	ID        string `bson:"_id,omitempty"`
	UserID    string `bson:"userId"`
	IsDeleted bool   `bson:"deleted"`
	Dates     *Dates `bson:"dates"`
}
