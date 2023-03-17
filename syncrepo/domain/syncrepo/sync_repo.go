package syncrepo

type OriginRepo struct {
	Endpoint string
	Repo     string
	Branch   string
}

type SyncRepo interface {
	Sync(*OriginRepo) (string, error)
}
