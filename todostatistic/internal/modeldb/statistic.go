package modeldb

type TodoStatistic struct {
	ID     string `bson:"_id,omitempty"`
	UserID string `bson:"userId"`
	TodoID string `bson:"todoId"`
	Month  string `bson:"month"`
	Year   string `bson:"year"`
}

type UserTodoStatistic struct {
	Month int64
	Year  int64
	Total int64
}
