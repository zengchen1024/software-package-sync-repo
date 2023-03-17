package platformimpl

import "github.com/opensourceways/software-package-sync-repo/syncrepo/domain"

func NewGithub() github {
	return github{}
}

type github struct{}

func (impl github) GetLastCommit(*domain.RepoInfo) (string, error) {
	return "", nil
}
