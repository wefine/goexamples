package main

//implements the SFTP Service. Needs to be installed and managed using NSSM.
import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type configuration struct {
	Source      string
	Host        string
	Port        int
	Username    string
	Password    string
	Destination string
}

func main() {

	conf := configuration{
		Username:    "root",
		Password:    "",
		Port:        22,
		Destination: "/home/dist",
	}

	if len(os.Args) < 3 {
		os.Exit(1)
	}

	conf.Host = os.Args[1]
	conf.Source = os.Args[2]

	if len(os.Args) == 4 {
		conf.Destination = os.Args[3]
	}

	err := uploadFiles(conf)
	if err != nil {
		log.Println(err)
	}
}

func uploadFiles(conf configuration) error {
	//init sftp client
	var authMethods []ssh.AuthMethod

	keyboardInteractiveChallenge := func(
		user,
		instruction string,
		questions []string,
		echos []bool,
	) (answers []string, err error) {
		if len(questions) == 0 {
			return []string{}, nil
		}
		return []string{conf.Password}, nil
	}
	//add KeyboardInteractive and Password auth methods
	authMethods = append(authMethods, ssh.KeyboardInteractive(keyboardInteractiveChallenge))
	authMethods = append(authMethods, ssh.Password(conf.Password))

	config := &ssh.ClientConfig{
		User:            conf.Username,
		Auth:            authMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	//get source folder
	files, err := ioutil.ReadDir(conf.Source)
	if err != nil {
		log.Fatal(err)
	}

	if len(files) > 0 {
		//open sftp connection
		client, err := ssh.Dial("tcp", conf.Host+":"+strconv.Itoa(conf.Port), config)
		if err != nil {
			log.Fatal(err)
		}

		sftp, err := sftp.NewClient(client)
		if err != nil {
			log.Fatal(err)
		}
		defer sftp.Close()

		for _, f := range files {
			filename := f.Name()
			if strings.HasSuffix(filename, ".zip") {
				sourcePath := conf.Source + "\\" + filename
				fmt.Println(sourcePath)
				b, err := ioutil.ReadFile(sourcePath)
				if err != nil {
					log.Fatal(err)
				}

				//write file to server
				destPath := conf.Destination + "/" + filename
				f, err := sftp.Create(destPath)
				if err != nil {
					log.Fatal(err)
				}

				defer f.Close()
				f.Write(b)
			}
		}
	}
	return nil
}
