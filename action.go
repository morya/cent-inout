package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/centrifugal/protocol"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func action(c *cli.Context) error {
	opt := []nats.Option{
		nats.ReconnectBufSize(-1),
		nats.MaxReconnects(-1),
	}
	nc, err := nats.Connect(NatsAddress, opt...)
	if err != nil {
		return err
	}

	_, err = nc.Subscribe(NatsTopic, handler)
	if err != nil {
		return err
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)

	<-sig
	nc.Close()

	return err
}

func handler(msg *nats.Msg) {
	p := &protocol.Push{}
	if err := p.UnmarshalVT(msg.Data); err != nil {
		logrus.WithError(err).Debugf("WrapHandle UnmarshalVT msg err: %v, msg data: %v", err, msg.Data)
		return
	}

	// if p.Join == nil && p.Leave == nil {
	// 	// centrifuge v5
	// 	return
	// }

	if p.Join != nil {
		logrus.Infof("got join %s", p.Join.String())
	} else if p.Leave != nil {
		logrus.Infof("got leave %s", p.Leave.String())
	}
}
