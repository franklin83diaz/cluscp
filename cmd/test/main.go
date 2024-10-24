package main

import (
	"cluscp/pkg"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func main() {
	listHost := []string{"127.0.0.1"}
	var sessions []*ssh.Session
	var err error
	sessions, err = pkg.ConnectSSH(listHost, "password")
	if err != nil {
		panic(err)
	}
	session := sessions[0]

	//check if ncopy is installed
	output, err := session.CombinedOutput("ncopy -h")
	if err != nil {
		fmt.Println("ncopy is not installed")
	}

	println(string(output))

	session.Close()

}
