package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
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
	Payload interface{} `json:"data"`
}

// JSONB Interface for JSONB Field of yourTableName Table
type JSONB map[string]interface{}

// Value Marshal
func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
