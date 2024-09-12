package model

import (
	"github.com/google/uuid"
)

const (
	LogEvent        = "Log"
	UserEvent       = "User"
	MonitoringEvent = "Monitoring"
)

type BaserEventHeader struct {
	ID            uuid.UUID `gorm:"type:uuid"`
	Producer      string    `json:"producer"`
	Client        string    `json:"client"`
	ReferenceName string    `json:"reference-name"`
}

func (h *BaserEventHeader) Init() {
	h.ID = uuid.New()
}

type BaseEventMessage struct {
	BaserEventHeader
}
