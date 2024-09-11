package app

import (
	"context"
	"fmt"
	"github.com/ivanmeca/DesafioPismo/v2/config"
	"github.com/ivanmeca/DesafioPismo/v2/pkg/database"
	"go.uber.org/zap"
)

type IApplication interface {
	Init(config *config.Config) error
	Run(ctx context.Context) error
}

type app struct {
	appCtx         context.Context
	configuration  *config.Config
	dataRepository database.IEventRepository
	logger         zap.Logger
}

func NewApp() IApplication {
	return &app{}
}

func (a *app) Init(config *config.Config) error {
	a.configuration = config
	a.logger = zap.Logger{}

	db, err := database.StartDB(config.GetDB())
	if err != nil {
		return err
	}

	a.dataRepository = database.NewGormRepository(db, a.logger)
	database.RunMigrations(db)
	return nil
}

func (a *app) Run(ctx context.Context) error {

	fmt.Println("App is running")
	return nil
}
