package synclockimpl

import (
	"github.com/opensourceways/software-package-sync-repo/syncrepo/domain"
	"github.com/opensourceways/software-package-sync-repo/syncrepo/domain/synclock"
)

func NewRepoSyncLock() syncLock {
	return syncLock{}
}

type syncLock struct {
}

func (impl syncLock) TryLock(*domain.RepoInfo) (r synclock.SyncInfo, err error) {
	return
}

func (impl syncLock) Unlock(*domain.RepoInfo, synclock.SyncInfo) error {
	return nil
}
