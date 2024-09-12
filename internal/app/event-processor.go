package app

import (
	"encoding/json"
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
		ep.logger.Info("New log message received")
		return ep.HandleLogEvent(message)
	case model.MonitoringEvent:
		ep.logger.Info("New monitoring message received")
		return ep.HandleMonitoringEvent(message)
	case model.UserEvent:
		ep.logger.Info("New user message received")
		return ep.HandleUserEvent(message)
	default:
		ep.logger.Errorf("Unhandled event: %s", event.ReferenceName)
		return false
	}
}

func (ep *EventProcessor) HandleMonitoringEvent(message []byte) bool {
	data := model.MonitoringEventMessage{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		ep.logger.Errorf("Error unmarshalling message: %s", err)
		return false
	}
	_, err = ep.dataRepository.CreateMonitoringEvent(&data)
	if err != nil {
		ep.logger.Errorf("Error saving event: %s", err)
		return false
	}
	return true
}

func (ep *EventProcessor) HandleLogEvent(message []byte) bool {
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
	return true
}

func (ep *EventProcessor) HandleUserEvent(message []byte) bool {
	data := model.UserOperationEventMessage{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		ep.logger.Errorf("Error unmarshalling message: %s", err)
		return false
	}
	_, err = ep.dataRepository.CreateUserOperationEvent(&data)
	if err != nil {
		ep.logger.Errorf("Error saving event: %s", err)
		return false
	}
	return true
}
