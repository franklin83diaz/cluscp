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
	// if the user choose ip
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
	// if the user choose hostname
	if hostOrIp == entities.HostsType(2) {
		// Ask for the hostname
		hostname := utils.AskMinLength("Enter the hostname:", 1)
		start := utils.AskNumber("Counter start:", 0, 4294967295)
		end := utils.AskNumber("Counter end:", 0, 4294967295)
		domain := utils.AskMinLength("Enter the domain:", 1)

		fqdnList, err := utils.GetFqdnList(hostname, start, end, domain)
		if err != nil {
			panic(err)
		}

		createHostsFile(fileHost, fqdnList)
	}
}

func createHostsFile(fileHost string, hosts []string) {
	// Create the file
	file, err := os.Create(fileHost)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the ips in the file
	for _, host := range hosts {
		_, err = file.WriteString(host + "\n")
		if err != nil {
			panic(err)
		}
	}

}
