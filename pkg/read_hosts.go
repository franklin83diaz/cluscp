package pkg

import (
	"fmt"
	"os"
)

// Read Hosts from file
// Receives a string with the file path and name of the file
// Returns a slice of strings with the hosts
func ReadHosts(fileHost string) ([]string, error) {

	//Read the file
	file, err := os.Open(fileHost)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileInfo.Size()

	if fileSize == 0 {
		return nil, fmt.Errorf("file %s is empty", fileHost)
	}

	if fileSize > 5120000000 /* 5MB */ {
		return nil, fmt.Errorf("file %s is too large", fileHost)
	}

	//Create Buffer
	buffer := make([]byte, fileSize)

	//Read the file
	file.Read(buffer)

	//!TODO: Implement the logic to read the hosts from the file
	return []string{}, nil

}
