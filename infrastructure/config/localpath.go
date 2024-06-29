package config

import (
	"path/filepath"

	"github.com/kirsle/configdir"
	"github.com/samber/do"
)

type LocalPathService struct {
	localPath string
}

func NewLocalPathService(i *do.Injector) (*LocalPathService, error) {
	configPath := configdir.LocalConfig("wails-todo")
	if err := configdir.MakePath(configPath); err != nil {
		return nil, err
	}
	return &LocalPathService{localPath: configPath}, nil
}

func (c *LocalPathService) GetJoinedPath(filename string) string {
	return filepath.Join(c.localPath, filename)
}

func (c *LocalPathService) GetLocalPath() string {
	return c.localPath
}
