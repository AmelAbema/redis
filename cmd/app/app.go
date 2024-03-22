package app

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
	"log"
	"redis/pkg/storage"
)

type Redis struct{}

func New() *Redis { return &Redis{} }

func (s *Redis) CliStart() *cli.Command {
	return &cli.Command{
		Name:        "start",
		Usage:       "start service",
		Action:      s.Action,
		Subcommands: []*cli.Command{},
	}
}
func (s *Redis) CliShow() *cli.Command {
	return &cli.Command{
		Name:        "show",
		Usage:       "show how service works",
		Action:      s.ShowAction,
		Subcommands: []*cli.Command{},
	}
}

func errHandle(str string, err error) {
	if err == nil {
		return
	}
	log.Fatalf("Msg: %v, Err: %v", str, err)
}

func (*Redis) Action(cliC *cli.Context) error {
	ctx := cliC.Context
	//load config
	appConfig := storage.App{}
	errHandle("config load error", env.Parse(&appConfig))

	storageModule, errStor := storage.NewStore(ctx, &appConfig)
	errHandle("storage load error", errStor)

	//launch app
	wgroup, _ := errgroup.WithContext(ctx)
	wgroup.Go(func() error {
		return storageModule.Do()
	})
	return errors.WithStack(wgroup.Wait())
}

func (*Redis) ShowAction(cliC *cli.Context) error {
	ctx := cliC.Context

	storageModule, errStor := storage.NewStore(ctx, &storage.App{})
	errHandle("storage load error", errStor)
	//launch app
	wgroup, _ := errgroup.WithContext(ctx)
	wgroup.Go(func() error {
		fmt.Println()
		return storageModule.Show()
	})
	return errors.WithStack(wgroup.Wait())
}
