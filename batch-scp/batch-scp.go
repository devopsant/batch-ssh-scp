package main

// batch-scp.go

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"strings"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)


func sftpconnect(user, password, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
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

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}

	return sftpClient, nil
}

func scpCopy(localFilePath string, remoteDir string, user string, userPwd string, ip string) error {
	fmt.Println("uploading.....：", ip)
	var (
		sftpClient *sftp.Client
		err        error
	)
	sftpClient, err = sftpconnect(user, userPwd, ip, 22)
	if err != nil {
		fmt.Println("upload error:", err)
		return err
	}
	defer sftpClient.Close()
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		fmt.Println("transfer error:", err)
		return err
	}
	defer srcFile.Close()

	var remoteFileName = path.Base(localFilePath)
	dstFile, err := sftpClient.Create(path.Join(remoteDir, remoteFileName))
	if err != nil {
		fmt.Println("transfer error:", err)
		return err
	}
	defer dstFile.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf[0:n])
	}
	fmt.Println("upload done：", ip)
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

	fmt.Println("---bscp start run---")

	fmt.Println("please input properties file name,default is host.properties,don't need properties postfix:")
	configFileName := getInput()
	configFileName = strings.TrimSpace(configFileName)
	if configFileName == ""{
		configFileName = "host"
	}
	hostMaps := initConfig("./" + configFileName + ".properties")

	for {
		fmt.Println("input local file pull path:")
		localFilePath := getInput()
		localFilePath = strings.ReplaceAll(localFilePath, "\\", "/")
		fmt.Println(localFilePath)

		fmt.Println("input remote server dir full path:")
		remoteDir := getInput()
		fmt.Println(remoteDir)

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
			remoteDir = strings.ReplaceAll(remoteDir, " ", "")
			localFilePath = strings.ReplaceAll(localFilePath, " ", "")
			err := scpCopy(localFilePath, remoteDir, user, userpwd, k)
			if err != nil {
				fmt.Println("error in execute. ip:" + k + " error:" + err.Error())
				break
			}
		}
	}
	time.Sleep(100 * time.Hour)
}
