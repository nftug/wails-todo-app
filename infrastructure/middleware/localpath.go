package middleware

import (
	"path/filepath"

	"github.com/kirsle/configdir"
	"github.com/samber/do"
)

type LocalPathService interface {
	GetJoinedPath(filename string) string
	GetLocalPath() string
}

type localPathService struct {
	localPath string
}

func NewLocalPathService(i *do.Injector) (LocalPathService, error) {
	configPath := configdir.LocalConfig("wails-todo")
	if err := configdir.MakePath(configPath); err != nil {
		return nil, err
	}
	return &localPathService{localPath: configPath}, nil
}

func (lp *localPathService) GetJoinedPath(filename string) string {
	return filepath.Join(lp.localPath, filename)
}

func (lp *localPathService) GetLocalPath() string {
	return lp.localPath
}
