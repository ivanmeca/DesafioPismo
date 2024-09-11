package database

import (
	"github.com/ivanmeca/DesafioPismo/v2/internal/model"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(model.LogEventMessage{})
	db.AutoMigrate(model.MonitoringEventMessage{})
	db.AutoMigrate(model.UserOperationEventMessage{})
}
