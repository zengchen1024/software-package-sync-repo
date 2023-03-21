package main

import (
	"context"

	"github.com/opensourceways/software-package-sync-repo/message-server/kafka"
	"github.com/opensourceways/software-package-sync-repo/syncrepo/app"
)

type server struct {
	service   app.SyncService
	userAgent string
}

func (s *server) run(cfg *subscription, ctx context.Context) error {
	err := kafka.Subscriber().Subscribe(
		cfg.Group,
		map[string]kafka.Handler{
			cfg.Topic: s.handleCommitPushed,
		},
	)
	if err != nil {
		return err
	}

	<-ctx.Done()

	return nil
}

func (s *server) handleCommitPushed(data []byte, header map[string]string) error {
	msg := msgToHandleCommitPushed{s.userAgent}

	cmd, err := msg.toCmd(data, header)
	if err != nil {
		return err
	}

	return s.service.SyncRepo(&cmd)
}
