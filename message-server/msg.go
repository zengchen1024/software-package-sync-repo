package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/google/go-github/v50/github"

	"github.com/opensourceways/software-package-sync-repo/syncrepo/app"
)

const (
	eventTypePush      = "push"
	msgHeaderUUID      = "X-GitHub-Delivery"
	msgHeaderUserAgent = "User-Agent"
	msgHeaderEventType = "X-GitHub-Event"
)

type msgToHandleCommitPushed struct {
	userAgent string
}

func (msg *msgToHandleCommitPushed) toCmd(payload []byte, header map[string]string) (
	cmd app.CmdToSyncRepo, err error,
) {
	eventType, err := msg.parseRequest(header)
	if err != nil {
		err = fmt.Errorf("invalid msg, err:%s", err.Error())

		return
	}

	if eventType != eventTypePush {
		err = errors.New("not pushed event")

		return
	}

	e := new(github.PushEvent)
	if err = json.Unmarshal(payload, e); err == nil {
		cmd, err = msg.genCmd(e)
	}

	return
}

func (msg *msgToHandleCommitPushed) genCmd(e *github.PushEvent) (cmd app.CmdToSyncRepo, err error) {
	repo := e.GetRepo()
	cmd.Owner = repo.GetOwner().GetLogin()
	cmd.Repo = repo.GetName()
	cmd.Endpoint = repo.GetURL()

	if v := strings.Split(e.GetRef(), "/"); len(v) == 0 {
		err = errors.New("can't parse branch")
	} else {
		cmd.Branch = v[len(v)-1]
	}

	return
}

func (msg *msgToHandleCommitPushed) parseRequest(header map[string]string) (
	eventType string, err error,
) {
	if header == nil {
		err = errors.New("no header")

		return
	}

	if header[msgHeaderUserAgent] != msg.userAgent {
		err = errors.New("unknown " + msgHeaderUserAgent)

		return
	}

	if eventType = header[msgHeaderEventType]; eventType == "" {
		err = errors.New("missing " + msgHeaderEventType)

		return
	}

	if header[msgHeaderUUID] == "" {
		err = errors.New("missing " + msgHeaderUUID)
	}

	return
}
