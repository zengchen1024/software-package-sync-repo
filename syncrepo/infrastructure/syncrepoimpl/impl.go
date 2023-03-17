package syncrepoimpl

import (
	"errors"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/opensourceways/software-package-sync-repo/syncrepo/domain/syncrepo"
	"github.com/opensourceways/software-package-sync-repo/utils"
)

const lastCommitTag = "===last commit==="

func NewSyncRepo(cfg *Config) *syncRepo {
	return &syncRepo{
		shell:      cfg.SyncRepoShell,
		workDir:    cfg.WorkDir,
		targetRepo: cfg.TargetRepo.remoteURL(),
	}
}

type syncRepo struct {
	shell      string
	workDir    string
	targetRepo string
}

func (impl *syncRepo) SyncRepo(origin *syncrepo.OriginRepo) (string, error) {
	params := []string{
		impl.shell,
		impl.workDir, lastCommitTag,
		origin.Repo, origin.Branch, origin.Endpoint,
		impl.targetRepo + origin.Repo,
	}

	v, err, _ := utils.RunCmd(params...)
	if err != nil {
		logrus.Errorf(
			"run sync shell, err=%s, params=%v",
			err.Error(), params[:len(params)-1],
		)

		return "", err
	}

	if r := strings.Split(string(v), lastCommitTag); len(r) == 2 {
		return r[1], nil
	}

	return "", errors.New("can't parse last commit")
}
