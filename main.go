package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func buildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "nats-address",
			Value:       "127.0.0.1:4222",
			Destination: &NatsAddress,
		},
		&cli.StringFlag{
			Name:        "nats-topic",
			Value:       "centrifugo.client.>",
			Destination: &NatsTopic,
		},
	}
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	app := cli.NewApp()

	app.Flags = buildFlags()
	app.Action = action

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
