package syncrepoimpl

import (
	"fmt"
	"strings"
)

const repoEnpointPrefix = "https://"

type Config struct {
	WorkDir       string     `json:"work_dir"        required:"true"`
	TargetRepo    targetRepo `json:"target_repo"     required:"true"`
	SyncRepoShell string     `json:"sync_repo_shell" required:"true"`
}

func (cfg *Config) Validate() error {
	return cfg.TargetRepo.validate()
}

// targetRepo
type targetRepo struct {
	Enpoint    string     `json:"endpoint"    required:"true"`
	Credential credential `json:"credential"  required:"true"`
}

func (t *targetRepo) validate() error {
	if !strings.HasSuffix(t.Enpoint, repoEnpointPrefix) {
		return fmt.Errorf("unsupported protocol")
	}

	return nil
}

func (t *targetRepo) remoteURL() string {
	e := strings.TrimSuffix(t.Enpoint, "/")

	return fmt.Sprintf(
		"%s%s:%sxi@%s/",
		repoEnpointPrefix,
		t.Credential.UserName,
		t.Credential.Token,
		strings.TrimPrefix(e, repoEnpointPrefix),
	)
}

// credential
type credential struct {
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}
