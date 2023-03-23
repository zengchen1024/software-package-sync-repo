package synclockimpl

import (
	"github.com/google/uuid"

	"github.com/opensourceways/software-package-sync-repo/syncrepo/domain"
)

const (
	busy = "busy"
	free = "free"
)

type RepoInfoDO struct {
	Id         uuid.UUID `gorm:"column:uuid;type:uuid"`
	Owner      string    `gorm:"column:owner"`
	Repo       string    `gorm:"column:repo"`
	Branch     string    `gorm:"column:branch"`
	Status     string    `gorm:"column:status"`
	LastCommit string    `gorm:"column:last_commit"`
}

func toRepoDO(i *domain.RepoInfo, do *RepoInfoDO) {
	*do = RepoInfoDO{
		Id:     uuid.New(),
		Owner:  i.Owner,
		Repo:   i.Repo,
		Branch: i.Branch,
		Status: busy,
	}
}

func (r *RepoInfoDO) isBusy() bool {
	return r.Status == busy
}
