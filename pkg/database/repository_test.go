package database

import (
	"github.com/ivanmeca/DesafioPismo/v2/config"
	"github.com/ivanmeca/DesafioPismo/v2/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
	"gorm.io/gorm"
	"testing"
)

func TestRepository_CreateLogEvent(t *testing.T) {

	db := getTestDB(t)
	logger := zaptest.NewLogger(t)
	repo := NewGormRepository(db, *logger)
	RunMigrations(db)

	tData := model.LogEventMessage{}

	tData.Init()
	tData.Producer = "testProducer"
	tData.Client = "testSender"
	tData.ReferenceName = "testReference"
	tData.Payload = model.JSONB{
		"test":  "value1",
		"test2": "value2",
	}

	event, err := repo.CreateLogEvent(&tData)
	require.NoError(t, err)
	assert.NotEmpty(t, event)
	require.Equal(t, "testProducer", event.Producer)
	require.Equal(t, "testSender", event.Client)
	require.Equal(t, "testReference", event.ReferenceName)
}

func TestRepository_CreateMonitoringEvent(t *testing.T) {

	db := getTestDB(t)
	logger := zaptest.NewLogger(t)
	repo := NewGormRepository(db, *logger)
	RunMigrations(db)

	tData := model.MonitoringEventMessage{}

	tData.Init()
	tData.Producer = "testProducer"
	tData.Client = "testSender"
	tData.ReferenceName = "testReference"
	tData.Payload = model.JSONB{
		"metric1": "value1",
	}

	event, err := repo.CreateMonitoringEvent(&tData)
	require.NoError(t, err)
	assert.NotEmpty(t, event)
	require.Equal(t, "testProducer", event.Producer)
	require.Equal(t, "testSender", event.Client)
	require.Equal(t, "testReference", event.ReferenceName)

}

func TestRepository_CreateUserOperationEvent(t *testing.T) {

	db := getTestDB(t)
	logger := zaptest.NewLogger(t)
	repo := NewGormRepository(db, *logger)
	RunMigrations(db)

	tData := model.UserOperationEventMessage{}

	tData.Init()
	tData.Producer = "testProducer"
	tData.Client = "testSender"
	tData.ReferenceName = "testReference"
	tData.Payload = model.JSONB{
		"op":    "creation",
		"test2": "value2",
	}

	event, err := repo.CreateUserOperationEvent(&tData)
	require.NoError(t, err)
	assert.NotEmpty(t, event)
	require.Equal(t, "testProducer", event.Producer)
	require.Equal(t, "testSender", event.Client)
	require.Equal(t, "testReference", event.ReferenceName)

}

func getTestDB(t *testing.T) *gorm.DB {

	cfg := config.DBConfig{
		Host:     "localhost",
		Port:     "25432",
		User:     "admin",
		Name:     "pismo",
		Sslmode:  "disable",
		Password: "123456",
	}

	db, err := StartDB(cfg)
	require.NoError(t, err)

	return db
}
