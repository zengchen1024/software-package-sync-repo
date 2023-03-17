package synclock

import "github.com/opensourceways/software-package-sync-repo/syncrepo/domain"

type SyncInfo struct {
	LastCommit string
}

type RepoSyncLock interface {
	// TryLock try to lock the repo and return the sync info if locked
	TryLock(*domain.RepoInfo) (SyncInfo, error)

	// Unlock unlock the repo after saving the sync info
	Unlock(*domain.RepoInfo, SyncInfo) error
}
