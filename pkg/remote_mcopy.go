package pkg

import (
	"cluscp/pkg/utils"
	"log"
	"net"
	"os/user"

	"golang.org/x/crypto/ssh"
)

var password = "admin"
var username = "admin"

func ConnectSSH(listHost []string, authMod string) (sessions []*ssh.Session, err error) {

	if username == "" {
		currentUser, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		username = currentUser.Username
	}

	//get key
	key, _, err := utils.GenerateKey()
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	var config *ssh.ClientConfig
	if authMod == "password" {
		config = &ssh.ClientConfig{
			User: username,
			Auth: []ssh.AuthMethod{
				ssh.Password(password),
			},
			HostKeyCallback: func(hostname string, remote net.Addr, k ssh.PublicKey) error {
				return nil
			},
			Config: ssh.Config{},
		}
	} else {
		config = &ssh.ClientConfig{
			User: username,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
			Config: ssh.Config{},
		}
	}

	// Dial the remote server
	client, err := ssh.Dial("tcp", listHost[0]+":22", config)
	if err != nil {
		return
	}

	sessions = make([]*ssh.Session, 1)

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}

	sessions[0] = session

	return sessions, nil

}
