package database

import (
	"github.com/ivanmeca/DesafioPismo/v2/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IEventRepository interface {
	CreateLogEvent(e *model.LogEventMessage) (*model.LogEventMessage, error)
	CreateMonitoringEvent(e *model.MonitoringEventMessage) (*model.MonitoringEventMessage, error)
	CreateUserOperationEvent(e *model.UserOperationEventMessage) (*model.UserOperationEventMessage, error)
}

type EventRepository interface {
	IEventRepository
}

func NewGormRepository(db *gorm.DB, logger zap.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

type Repository struct {
	logger zap.Logger
	db     *gorm.DB
}

func (r *Repository) CreateLogEvent(e *model.LogEventMessage) (*model.LogEventMessage, error) {

	r.logger.Debug("Create Log Event", zap.Any("event", e))
	err := r.db.Create(&e).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *Repository) CreateMonitoringEvent(e *model.MonitoringEventMessage) (*model.MonitoringEventMessage, error) {

	r.logger.Debug("Create Log Event", zap.Any("event", e))
	err := r.db.Create(&e).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *Repository) CreateUserOperationEvent(e *model.UserOperationEventMessage) (*model.UserOperationEventMessage, error) {

	r.logger.Debug("Create Log Event", zap.Any("event", e))
	err := r.db.Create(&e).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}
