package app

import (
	"context"
	"github.com/ivanmeca/DesafioPismo/v2/config"
	"github.com/ivanmeca/DesafioPismo/v2/pkg/database"
	queue "github.com/ivanmeca/DesafioPismo/v2/pkg/queues"
	"go.uber.org/zap"
)

type IApplication interface {
	Init(config *config.Config) error
	Run(ctx context.Context) error
}

type app struct {
	appCtx          context.Context
	configuration   *config.Config
	dataRepository  database.IEventRepository
	queueRepository *queue.Repository
	eventQueue      *queue.Queue
	ep              *EventProcessor
	logger          *zap.SugaredLogger
}

func NewApp() IApplication {
	return &app{}
}

func (a *app) Init(config *config.Config) error {
	a.configuration = config
	logger, _ := zap.NewProduction()
	a.logger = logger.Sugar()
	a.logger.Info("Initializing application")

	db, err := database.StartDB(config.GetDB())
	if err != nil {
		return err
	}

	a.dataRepository = database.NewGormRepository(db, *a.logger)
	database.RunMigrations(db)

	a.ep = NewEventProcessor(a.dataRepository, *a.logger)
	a.queueRepository, err = queue.NewQueueRepository(a.configuration.GetQueue())
	if err != nil {
		return err
	}

	params := queue.NewQueueParams("Events")
	q, err := a.queueRepository.QueueDeclare(params, true)
	if err != nil {
		return err
	}

	a.eventQueue = q
	return nil
}

func (a *app) Run(ctx context.Context) error {

	err := a.eventQueue.StartConsume(func(queueName string, msg []byte) bool {
		a.logger.Info("New message received")
		return a.ep.HandleMessage(msg)
	})
	if err != nil {
		return err
	}
	a.logger.Info("App is running")
	for {
		<-ctx.Done()
		return nil
	}
}
