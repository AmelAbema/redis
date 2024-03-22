package main

import (
	"context"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
	"redis/cmd/app"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, os.Interrupt)
		<-ch
		log.Println("gracefully shutdown sequence stated")

		cancel()

		time.Sleep(3 * time.Second)
		log.Println("gracefully shutdown sequence ended")
		os.Exit(1)
	}()

	appc := &cli.App{
		Name: "redis",
		Commands: []*cli.Command{
			app.New().CliStart(),
			app.New().CliShow(),
		},
	}

	if err := appc.RunContext(ctx, os.Args); err != nil {
		log.Fatalln("start app fail: ", err)
	}
}
