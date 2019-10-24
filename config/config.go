package config

import (
	"os"

	configstorage "github.com/thelioncape/GoNetworkMonitor/config/storage"
)

// GetStorageType returns the storage type and the storage location that has been selected or detected
func GetStorageType() (typeandlocation [2]string, err error) {
	var lines [2]string
	if exists("storageconfig.cfg") {
		lines = configstorage.LoadStorageConfig()
	} else {
		lines = configstorage.ChooseStorage()
	}

	return lines, err
}

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
