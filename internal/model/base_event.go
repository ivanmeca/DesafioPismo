package model

const (
	LogEvent        = "LogEvent"
	UserEvent       = "UserEvent"
	MonitoringEvent = "MonitoringEvent"
)

type BaserEventHeader struct {
	Producer      string `json:"producer"`
	Sender        string `json:"sender"`
	ReferenceName string `json:"reference-name"`
}

type BaseEventMessage struct {
	BaserEventHeader
	Payload interface{} `json:"data"`
}
