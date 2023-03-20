package main

import (
	"encoding/json"

	"github.com/opensourceways/software-package-sync-repo/syncrepo/app"
)

type server struct {
	service app.SyncService
}

func (s *server) handleCommitPushed(data []byte) error {
	msg := new(msgToHandleCommitPushed)

	if err := json.Unmarshal(data, msg); err != nil {
		return err
	}

	cmd, err := msg.toCmd()
	if err != nil {
		return err
	}

	return s.service.SyncRepo(&cmd)
}

type msgToHandleCommitPushed struct{}

func (msg *msgToHandleCommitPushed) toCmd() (cmd app.CmdToSyncRepo, err error) {
	return
}
