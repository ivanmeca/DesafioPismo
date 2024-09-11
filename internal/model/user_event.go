package model

type UserOperationEventPayload struct {
	UserID    string `json:"user-id"`
	Operation string `json:"operation"`
}

type UserOperationEventMessage struct {
	BaserEventHeader
	Payload JSONB `gorm:"type:jsonb",json:"data"`
}
