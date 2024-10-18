package pkg

import (
	"cluscp/entities"
	"cluscp/pkg/utils"
	"fmt"
	"os"
)

// GenerateHostsFile generates a file with a list of hosts
// Receives no parameters get the data from the user stdin
// if has an error do a panic
func GenerateHostsFile() {
	// Define the type of hosts unknown
	hostOrIp := entities.HostsType(0)

	// Ask for the file name
	fileHost := utils.AskMinLength("Enter the file name to generate:", 1)

	// Ask if use ip or hostname
	respHostOrIp := utils.AskOptions("Do you want to use IP or Hostname? (I/H)", []string{"I", "H", "i", "h"})

	// Set the type of hosts
	hostOrIp.Set(respHostOrIp)

	// entities.HostsType(1) is the value of ip
	// entities.HostsType(2) is the value of hostname
	fmt.Println(hostOrIp)
	if hostOrIp == entities.HostsType(1) {
		// Ask for the ip
		ipsRange := utils.AskMinLength("Enter the Range of ips (example 10.0.0.1-200):", 7)
		ips, err := utils.GetIpList(ipsRange)
		if err != nil {
			panic(err)
		}
		createHostsFile(fileHost, ips)
	}
}

func createHostsFile(fileHost string, ips []string) {
	// Create the file
	file, err := os.Create(fileHost)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the ips in the file
	for _, ip := range ips {
		_, err = file.WriteString(ip + "\n")
		if err != nil {
			panic(err)
		}
	}

}
