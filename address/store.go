package address

import (
	"encoding/json"

	"github.com/shibukawa/configdir"
)

// GetAddresses gets the config addresses
func GetAddresses() []string {
	configDirs := configdir.New("flexpool", "flexpool-cli")
	folder := configDirs.QueryFolderContainsFile("settings.json")
	var addresses []string
	if folder != nil {
		data, _ := folder.ReadFile("settings.json")
		json.Unmarshal(data, &addresses)
	}
	return addresses
}

func setAddresses(addresses []string) error {
	configDirs := configdir.New("flexpool", "flexpool-cli")
	folders := configDirs.QueryFolders(configdir.Global)
	data, _ := json.Marshal(addresses)
	return folders[0].WriteFile("settings.json", data)
}

func containsAddress(addressChecksummed string) bool {
	for _, addr := range GetAddresses() {
		if addr == addressChecksummed {
			return true
		}
	}

	return false
}
