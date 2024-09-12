package model

type MonitoringEventPayload struct {
	ObjectID string `json:"object-id"`
	TraceId  string `json:"traceId"`
	Message  string `json:"message"`
}

type MonitoringEventMessage struct {
	BaserEventHeader
	MonitoringEventPayload
}
