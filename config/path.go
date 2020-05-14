package config

import (
	"github.com/shibukawa/configdir"
)

func getPath() (string, bool) {
	configDir := configdir.New("flexpool", "flexpool-cli")
	path := configDir.QueryFolderContainsFile("settings.json")
	if path == nil {
		return "", false
	}
	return path.Path, true
}
