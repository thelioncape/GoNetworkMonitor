package configstorage

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// StorageConfigChooser takes the passed string and tries to parse it into a configuration option
func storageConfigChooser(text string) string {
	switch text {
	case "1":
		return "Flatfile"
	default:
		return "Undefined"
	}
}

func listStorageConfigs() {
	fmt.Println("1. Flatfile")
}

// SaveStorageConfig saves the chosen storage config to disk
func SaveStorageConfig(storage [2]string) {
	data := []byte(storage[0] + "\n" + storage[1])
	err := ioutil.WriteFile("storageconfig.cfg", data, 0733)
	if err != nil {
		log.Fatal(err)
	}
}

// ChooseStorage asks the user for the type of storage they would like to use and the location of the storage
// This should only be run on first launch or a configuration file loss
func ChooseStorage() [2]string {
	var storage [2]string
	storage[0] = chooseStorageType()
	storage[1] = chooseStorageLocation(storage[0])
	return storage
}

func chooseStorageType() string {
	fmt.Println("Please select a storage mechanism from the list below")
	listStorageConfigs()

	input := getUserChoice()
	storageType := storageConfigChooser(input)
	fmt.Println(storageType, "chosen")
	return storageType
}

func chooseStorageLocation(storagetype string) string {
	switch storagetype {

	case "Flatfile":
		return chooseFlatfileLocation()

	default:
		log.Fatal("Unknown storage type")
		return ""
	}
}

func chooseFlatfileLocation() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The file will be stored in the following location")
	fmt.Println(dir)
	return dir
}

func getUserChoice() string {
	fmt.Print("Enter your choice: ")
	var input string
	fmt.Scanln(&input)
	return input
}
