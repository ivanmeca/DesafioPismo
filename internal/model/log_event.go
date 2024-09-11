package model

type LogEventPayload struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

type LogEventMessage struct {
	BaserEventHeader
	Payload JSONB `gorm:"type:jsonb",json:"data"`
}
