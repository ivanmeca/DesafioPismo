package main

import (
	"context"
	"github.com/ivanmeca/DesafioPismo/v2/config"
	"github.com/ivanmeca/DesafioPismo/v2/internal/app"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"sort"
)

func runApplication(cli *cli.Context) error {

	appMan := app.NewApp()
	cfg, errLoad := config.Load()
	if errLoad != nil {
		panic(errLoad)
	}

	err := appMan.Init(cfg)
	if err != nil {
		return err
	}

	return appMan.Run(gracefullyShutdown())
}

func gracefullyShutdown() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		cancel()
	}()
	return ctx
}

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{},
	}
	app.Version = "1.0"
	app.Name = "Pismo challenge app"
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
