package pkg

import (
	"cluscp/entities"
	"cluscp/pkg/utils"
	"os"
)

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
	if hostOrIp == entities.HostsType(1) {
		// Ask for the ip
		ips := utils.AskMinLength("Enter the Range of ips (example 10.0.0.1-200):", 7)
		createHostsFile(fileHost, ips)
	}
}

func createHostsFile(fileHost string, ips string) {
	// Create the file
	file, err := os.Create(fileHost)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// ips = strings.TrimSpace(ips)
	// ipsSplit := strings.Split(ips, "-")

	// ipStart := strings.Split(ipsSplit[0], ".")
	// ipEnd := strings.Split(ipsSplit[1], ".")

	// Write the ips in the file
	_, err = file.WriteString(ips)
	if err != nil {
		panic(err)
	}

}
