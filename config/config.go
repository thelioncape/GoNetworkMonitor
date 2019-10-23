package config

import (
	"bufio"
	"log"
	"os"

	configstorage "github.com/thelioncape/GoNetworkMonitor/config/storage"
)

// GetStorageType returns the storage type and the storage location that has been selected or detected
func GetStorageType() (typeandlocation [2]string, err error) {
	var lines [2]string
	if exists("storageconfig.cfg") {
		f, err := os.Open("storageconfig.cfg")
		if err != nil {

		}
		defer f.Close()

		reader := bufio.NewReader(f)
		lines[0], err = reader.ReadString('\n')
		check(err)
		lines[1], err = reader.ReadString('\n')
		check(err)
	} else {
		storageType := configstorage.ChooseStorage()

	}

	return lines, err
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
