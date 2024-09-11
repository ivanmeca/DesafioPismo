package model

type LogEventPayload struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

type LogEventMessage struct {
	BaseEventMessage
	Payload LogEventPayload `json:"data"`
}
