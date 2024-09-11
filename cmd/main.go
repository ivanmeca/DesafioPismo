package main

import (
	"context"
	"github.com/ivanmeca/DesafioPismo/v2/config"
	"github.com/ivanmeca/DesafioPismo/v2/pkg/database"
	"os"
	"os/signal"
	"sort"
)

func runApplication(cli *cli.Context) error {
	c := context.Background()
	ctx, cancel := context.WithCancel(c)
	defer cancel()

	appMan := application.NewApp()

	cfg, errLoad := config.Load()
	if errLoad != nil {
		panic(errLoad)
	}

	db, err := database.StartDB(cfg.GetDB())
	if err != nil {
		return err
	}

	err = appMan.Init(ctx, cfg, db)
	if err != nil {
		return err
	}

	err = appMan.Run(ctx)
	if err != nil {
		return err
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	return nil
}

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{},
	}
	app.Version = "1.0"
	app.Name = "ApplicationName"
	app.Usage = "Usage"
	app.Description = "Event Processor"
	app.Copyright = "Copyright"
	app.EnableBashCompletion = true
	app.Action = runApplication
	app.Commands = []cli.Command{}
	sort.Sort(cli.FlagsByName(app.Flags))
	err := app.Run(os.Args)
	if err != nil {
		panic(err.Error())
	}

}
