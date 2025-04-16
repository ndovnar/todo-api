package model

// rename to Todo statistic
type TodoStatistic struct {
	ID     string
	UserID string
	TodoID string
	Month  string
	Year   string
}

// rename to User statistic
type UserTodoStatistic struct {
	Month int64
	Year  int64
	Total int64
}
