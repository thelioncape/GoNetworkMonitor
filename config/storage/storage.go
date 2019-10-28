package configstorage

import (
	"bufio"
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
		return "SQLite"
	default:
		return "Undefined"
	}
}

func listStorageConfigs() {
	fmt.Println("1. SQLite")
}

// SaveStorageConfig saves the chosen storage config to disk
func SaveStorageConfig(storage [2]string) {
	data := []byte(storage[0] + "\n" + storage[1] + "\n")
	err := ioutil.WriteFile("storageconfig.cfg", data, 0700)
	if err != nil {
		log.Fatal(err)
	}
}

// LoadStorageConfig reads the config file in the same directory and returns the storage type and location
// without the trailing backslash
func LoadStorageConfig() [2]string {
	var lines [2]string
	f, err := os.Open("storageconfig.cfg")
	if err != nil {
		log.Fatal(err, "\nThere was an error when reading the configuration file (storageconfig.cfg)")
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	lines[0], err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err, "\nThere was an error when reading the configuration file (storageconfig.cfg)")
	}
	lines[1], err = reader.ReadString('\n')
	check(err)

	return lines

}

// ChooseStorage asks the user for the type of storage they would like to use and the location of the storage
// This should only be run on first launch or a configuration file loss
func ChooseStorage() [2]string {
	var storage [2]string
	storage[0] = chooseStorageType()
	storage[1] = chooseStorageLocation(storage[0])
	SaveStorageConfig(storage)
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

	case "SQLite":
		return chooseSQLiteLocation()

	default:
		log.Fatal("Unknown storage type")
	}

	return ""
}

func chooseSQLiteLocation() string {
	fmt.Print("Please enter a directory to store the SQLite database: ")
	var input string
	fmt.Scanln(&input)

	path, err := filepath.Abs(input)

	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll(path, 600)

	if err != nil {
		log.Fatal(err, "\nIssue accessing folder - check permissions")
	}

	fmt.Println("The database will be stored in the following location")
	fmt.Println(path)
	return path
}

func getUserChoice() string {
	fmt.Print("Enter your choice: ")
	var input string
	fmt.Scanln(&input)
	return input
}

func check(err error) {
	if err != nil {
		log.Fatal(err, "There was an error reading the config file (storageconfig.cfg)")
	}
}
