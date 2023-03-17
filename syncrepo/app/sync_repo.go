package app

import (
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/software-package-sync-repo/syncrepo/domain"
	"github.com/opensourceways/software-package-sync-repo/syncrepo/domain/platform"
	"github.com/opensourceways/software-package-sync-repo/syncrepo/domain/synclock"
	"github.com/opensourceways/software-package-sync-repo/syncrepo/domain/syncrepo"
	"github.com/opensourceways/software-package-sync-repo/utils"
)

type CmdToSyncRepo struct {
	Owner string
	syncrepo.OriginRepo
}

func (cmd *CmdToSyncRepo) repoInfo() domain.RepoInfo {
	return domain.RepoInfo{
		Owner:  cmd.Owner,
		Repo:   cmd.Repo,
		Branch: cmd.Branch,
	}
}

type SyncService interface {
	SyncRepo(*CmdToSyncRepo) error
}

func NewSyncService(
	p platform.Platform,
	l synclock.RepoSyncLock,
	s syncrepo.SyncRepo,
) *syncService {
	return &syncService{
		lock:     l,
		platform: p,
		syncrepo: s,
	}
}

type syncService struct {
	lock     synclock.RepoSyncLock
	platform platform.Platform
	syncrepo syncrepo.SyncRepo
}

func (s *syncService) SyncRepo(cmd *CmdToSyncRepo) error {
	repo := cmd.repoInfo()

	info, err := s.lock.TryLock(&repo)
	if err != nil {
		return err
	}

	defer func() {
		s.unlock(&repo, &info)
	}()

	v, err := s.platform.GetLastCommit(&repo)
	if err == nil && info.LastCommit == v {
		return nil
	}

	v, err = s.syncrepo.Sync(&cmd.OriginRepo)
	if err == nil {
		info.LastCommit = v
	}

	return err
}

func (s *syncService) unlock(repoInfo *domain.RepoInfo, syncInfo *synclock.SyncInfo) {
	err := utils.Retry(func() error {
		return s.lock.Unlock(repoInfo, *syncInfo)
	})

	if err == nil {
		return
	}

	logrus.Errorf(
		"unlock repo(%s) failed, dead lock happened",
		repoInfo.String(),
	)
}
