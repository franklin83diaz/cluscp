package main

import (
	"cluscp/pkg"
	"cluscp/pkg/utils"
	"flag"
	"fmt"
)

func main() {

	fileName := flag.String("out", "nodes.hosts", "Output file name for hosts file")
	ips := flag.String("ips", "", "Range of IPs to generate example: 10.0.0.1-200")
	host := flag.String("hostname", "", "name base for hosts example: node")
	domainName := flag.String("domain", "", "Domain name for hosts file example: adaptivecomputing.com")
	start := flag.Int("start", 1, "Start number for hosts")
	end := flag.Int("end", 1, "End number for hosts")

	flag.Parse()

	if flag.NFlag() < 1 {
		pkg.GenerateHostsFile()
		return
	}

	if *ips != "" && (*host != "" || *domainName != "") {
		fmt.Println("Please provide either -ips or -hosts and -domain, not both")
		return
	}

	if *ips != "" {
		ips, err := utils.GetIpList(*ips)
		if err != nil {
			panic(err)
		}
		pkg.CreateHostsFile(*fileName, ips)
		return
	}

	if *host != "" && *domainName != "" {
		fqdnList, err := utils.GetFqdnList(*host, *start, *end, *domainName)
		if err != nil {
			panic(err)
		}
		pkg.CreateHostsFile(*fileName, fqdnList)
		return
	}

	//help
	flag.PrintDefaults()
}
