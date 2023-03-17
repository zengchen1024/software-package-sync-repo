package platform

import "github.com/opensourceways/software-package-sync-repo/syncrepo/domain"

type Platform interface {
	GetLastCommit(*domain.RepoInfo) (string, error)
}
