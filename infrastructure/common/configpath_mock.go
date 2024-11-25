package common

import (
	"path/filepath"
	"testing"
)

type configPathMockService struct {
	localPath string
}

func NewConfigPathMockService(t *testing.T) (ConfigPathService, error) {
	t.Helper()
	return &configPathMockService{t.TempDir()}, nil
}

func (lp *configPathMockService) GetJoinedPath(filename string) string {
	return filepath.Join(lp.GetLocalPath(), filename)
}

func (lp *configPathMockService) GetLocalPath() string {
	return lp.localPath
}
