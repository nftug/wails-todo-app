package common

import (
	"path/filepath"

	"github.com/kirsle/configdir"
	"github.com/samber/do"
)

type ConfigPathService interface {
	GetJoinedPath(filename string) string
	GetLocalPath() string
}

type configPathService struct {
	localPath string
}

func NewConfigPathService(i *do.Injector) (ConfigPathService, error) {
	configPath := configdir.LocalConfig("wails-todo")
	if err := configdir.MakePath(configPath); err != nil {
		return nil, err
	}
	return &configPathService{localPath: configPath}, nil
}

func (lp *configPathService) GetJoinedPath(filename string) string {
	return filepath.Join(lp.GetLocalPath(), filename)
}

func (lp *configPathService) GetLocalPath() string {
	return lp.localPath
}
