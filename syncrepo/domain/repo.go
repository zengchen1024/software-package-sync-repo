package domain

import "fmt"

type RepoInfo struct {
	Owner  string
	Repo   string
	Branch string
}

func (r *RepoInfo) String() string {
	return fmt.Sprintf("%s/%s/%s", r.Owner, r.Repo, r.Branch)
}
