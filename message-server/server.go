package main

import "github.com/opensourceways/software-package-sync-repo/syncrepo/app"

type server struct {
	service   app.SyncService
	userAgent string
}

func (s *server) handleCommitPushed(data []byte, header map[string]string) error {
	msg := msgToHandleCommitPushed{s.userAgent}

	cmd, err := msg.toCmd(data, header)
	if err != nil {
		return err
	}

	return s.service.SyncRepo(&cmd)
}
