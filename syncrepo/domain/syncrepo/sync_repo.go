package syncrepo

type OriginRepo struct {
	CloneURL string
	Repo     string
	Branch   string
}

type SyncRepo interface {
	Sync(*OriginRepo) (string, error)
}
