package synclockimpl

import (
	"errors"

	"github.com/opensourceways/software-package-sync-repo/syncrepo/domain"
	"github.com/opensourceways/software-package-sync-repo/syncrepo/domain/synclock"
)

func NewRepoSyncLock(cli dbClient) syncLock {
	return syncLock{cli: cli}
}

type syncLock struct {
	cli dbClient
}

func (impl syncLock) TryLock(i *domain.RepoInfo) (r synclock.SyncInfo, err error) {
	var do RepoInfoDO
	toRepoDO(i, &do)

	filter := RepoInfoDO{
		Owner:  i.Owner,
		Repo:   i.Repo,
		Branch: i.Branch,
	}

	err = impl.cli.FirstOrCreate(&filter, &do)
	if err == nil || !impl.cli.IsRowExists(err) {
		return
	}

	if do.isBusy() {
		err = errors.New("record busy")

		return
	}

	r.LastCommit = do.LastCommit

	filter.Status = free
	filter.LastCommit = do.LastCommit
	err = impl.cli.UpdateRecord(&filter, &RepoInfoDO{Status: busy})

	return
}

func (impl syncLock) Unlock(i *domain.RepoInfo, s synclock.SyncInfo) error {
	return impl.cli.UpdateRecord(
		&RepoInfoDO{Owner: i.Owner, Repo: i.Repo, Branch: i.Branch, Status: busy},
		&RepoInfoDO{Status: free, LastCommit: s.LastCommit},
	)
}
