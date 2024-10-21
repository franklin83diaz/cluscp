package main

import (
	"cluscp/pkg"
	"fmt"
)

func main() {
	pkg.GenerateHostsFile()
	hosts, err := pkg.ReadHosts("node-hosts")
	if err != nil {
		panic(err)
	}

	for _, host := range hosts {
		fmt.Println(host)
	}

}
