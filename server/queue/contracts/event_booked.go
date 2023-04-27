package contracts

type EventBookedEvent struct {
	EventID  string `json:"event_id"`
	MemberID string `json:"member_id"`
}

func (c *EventBookedEvent) EventName() string {
	return "eventBooked"
}
