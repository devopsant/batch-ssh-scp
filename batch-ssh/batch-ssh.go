package main

// batch-scp.go

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

func sshconnect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

func singleSsh(cmd string, user string, userPwd string, ip string) error {
	fmt.Println("connecting.....：", ip)
	sshSession, err := sshconnect(user, userPwd, ip, 22)
	if err != nil {
		fmt.Println("ssh error:", err)
		return err
	}
	defer sshSession.Close()

	cmdResult, err := sshSession.Output(cmd)
	if err != nil {
		fmt.Println("ssh execute error:", err)
		return err
	}
	fmt.Println("ssh finish:" + ip + "\n", string(cmdResult))
	return nil
}

func initConfig(path string) map[string]string {
	config := make(map[string]string)

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config
}

func getInput() string {
	in := bufio.NewReader(os.Stdin)
	str, _, err := in.ReadLine()
	if err != nil {
		return err.Error()
	}
	return string(str)
}

func main() {

	fmt.Println("---bssh start run---")

	fmt.Println("please input properties file name,default is host.properties,don't need properties postfix:")
	configFileName := getInput()
	configFileName = strings.TrimSpace(configFileName)
	if configFileName == ""{
		configFileName = "host"
	}
	hostMaps := initConfig("./" + configFileName + ".properties")

	for {
		fmt.Println("input shell command to execute,don't include enter:")
		cmd := getInput()
		fmt.Println(cmd)
		for k, v := range hostMaps {
			if strings.Contains(k, "#") {
				continue
			}
			if v == "" {
				fmt.Println("value is null：", k)
				continue
			}
			user := strings.Split(v, "//")[0]
			userpwd := strings.Split(v, "//")[1]
			err := singleSsh(cmd, user, userpwd, k)
			if err != nil {
				fmt.Println("error in execute. ip:" + k + " error:" + err.Error())
				break
			}
		}
		fmt.Println("execute finish")
	}

	time.Sleep(100 * time.Hour)

}
