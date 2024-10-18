package pkg

import (
	"bufio"
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

	scanner := bufio.NewScanner(file)
	hosts := []string{}

	for scanner.Scan() {
		hosts = append(hosts, scanner.Text())
	}

	//!TODO: Implement the logic to read the hosts from the file
	return hosts, nil

}
