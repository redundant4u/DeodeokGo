package queue

type Event interface {
	EventName() string
}
