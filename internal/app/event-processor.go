package app

import (
	"encoding/json"
	"fmt"
	"github.com/ivanmeca/DesafioPismo/v2/internal/model"
	"github.com/ivanmeca/DesafioPismo/v2/pkg/database"
	"go.uber.org/zap"
)

func NewEventProcessor(r database.IEventRepository, logger zap.SugaredLogger) *EventProcessor {
	return &EventProcessor{
		dataRepository: r,
		logger:         logger,
	}
}

type EventProcessor struct {
	logger         zap.SugaredLogger
	dataRepository database.IEventRepository
}

func (ep *EventProcessor) HandleMessage(message []byte) bool {

	event := model.BaseEventMessage{}

	err := json.Unmarshal(message, &event)
	if err != nil {
		ep.logger.Errorf("Error unmarshalling message: %s", err)
		return false
	}

	switch event.ReferenceName {
	case model.LogEvent:
		data := model.LogEventMessage{}
		err := json.Unmarshal(message, &data)
		if err != nil {
			ep.logger.Errorf("Error unmarshalling message: %s", err)
			return false
		}
		_, err = ep.dataRepository.CreateLogEvent(&data)
		if err != nil {
			ep.logger.Errorf("Error saving event: %s", err)
			return false
		}
	}

	fmt.Println(event.ReferenceName)
	fmt.Println(string(message))

	return true
}
