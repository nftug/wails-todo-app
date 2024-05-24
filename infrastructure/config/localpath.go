package config

import (
	"log"
	"path/filepath"

	"github.com/kirsle/configdir"
)

type LocalPathService struct {
	localPath string
}

func NewLocalPathService() *LocalPathService {
	configPath := configdir.LocalConfig("wails-todo")
	if err := configdir.MakePath(configPath); err != nil {
		log.Fatal("cannot make local config path")
	}
	return &LocalPathService{localPath: configPath}
}

func (c *LocalPathService) GetJoinedPath(filename string) string {
	return filepath.Join(c.localPath, filename)
}

func (c *LocalPathService) GetLocalPath() string {
	return c.localPath
}
