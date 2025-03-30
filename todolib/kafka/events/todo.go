package events

type TodoEventType string

const (
	TodoCreatedEventType TodoEventType = "todo_created"
)

type TodoEvent struct {
	EventType TodoEventType `json:"eventType"`
	UserID    string        `json:"userId"`
	TodoID    string        `json:"todoId"`
}
