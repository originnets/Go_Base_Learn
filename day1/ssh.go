package main

import (
	"golang.org/x/crypto/ssh"
	"time"
	"fmt"
	"log"
	"net"
	"os"
)

func connect(user, password, host string, port int)(*ssh.Session, error) {
	var (
		auth	[]ssh.AuthMethod
		addr	string
		clientConfig	*ssh.ClientConfig
		cilent	*ssh.Client
		session	*ssh.Session
		err	error
 	)
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: auth,
		Timeout: 30* time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	addr = fmt.Sprintf("%s:%d", host, port)

	if  cilent, err = ssh.Dial("tcp", addr, clientConfig); err !=nil {
		return nil, err
	}

	if session, err = cilent.NewSession();err != nil {
		return nil, err
	}
	return session, nil
}




func main() {

	session, err := connect("root", "123456", "192.168.1.218", 22)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	session.Run("ls /; pwd")
	session.Stdin = os.Stdin
}

