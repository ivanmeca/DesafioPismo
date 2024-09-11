package app

import (
	"context"
	"fmt"
	"github.com/ivanmeca/DesafioPismo/v2/config"
	"gorm.io/gorm"
)

type IApplication interface {
	Init(ctx context.Context, config *config.Config, db *gorm.DB) error
	Run(ctx context.Context) error
}

type app struct {
	appCtx         context.Context
	configuration  *config.Config
	dataRepository database.IRepository
}

func NewApp() IApplication {
	return &app{}
}

func (a *app) Init(ctx context.Context, config *config.Config, db *gorm.DB) error {
	a.appCtx = ctx
	a.configuration = config
	a.dataRepository = database.NewRepository(db)
	//migrations.RunMigrations(database.GetDatabase())
	c := controllers.NewController(a.dataRepository)
	jwtService := services.NewJWTService(a.dataRepository)
	a.router = routes.SetupRouter(jwtService, c)
	return nil
}

func (a *app) Run(ctx context.Context) error {

	fmt.Println("App is running")
	return nil
}
