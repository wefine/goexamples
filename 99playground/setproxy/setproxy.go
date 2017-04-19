package main

import (
	"os/exec"
	"fmt"
	"os"
	"log"
	"strings"
	"bufio"
	"regexp"
)

func main() {
	var proxy string
	if len(os.Args) != 2 {
		ip := execCmd("w | grep ${USER} | awk '{print $3}' | uniq | tail -n 1")
		ip = strings.TrimSpace(ip)
		port := "8888"

		proxy = "http://" + ip + ":" + port
	} else {
		proxy = os.Args[1]
	}

	if len(proxy) < 10 {
		log.Println("Wrong proxy = ", proxy)
		os.Exit(1)
	}

	log.Println("Using proxy: " + proxy)

	proxyGit(proxy)
	proxyYarn(proxy)
	proxyNpm(proxy)
	proxyBower(proxy)
	proxyBashrc(proxy)
	proxySystemdDocker(proxy)

	log.Println("Done!")
}

func proxyGit(proxy string) {
	if isExisted("git") {
		log.Println("Proxying git...")

		// git config --global http.proxy http://username:password@host:port
		// git config --global https.proxy http://username:password@host:port
		execCmd("git config --global http.proxy " + proxy)
		execCmd("git config --global https.proxy " + proxy)
		execCmd("git config --global http.sslverify false")
		execCmd("git config --global url.'https://'.insteadOf git:// ")
	}
}

func proxyYarn(proxy string) {
	if isExisted("yarn") {
		log.Println("Proxying yarn...")

		// yarn config set proxy http://username:password@host:port
		// yarn config set https-proxy http://username:password@host:port
		execCmd("yarn config set proxy " + proxy)
		execCmd("yarn config set https-proxy " + proxy)
	}
}

func proxyNpm(proxy string) {
	if isExisted("npm") {
		log.Println("Proxying npm...")

		// npm config set proxy http://username:password@host:port
		// npm config set https-proxy http://username:password@host:port
		execCmd("npm config set proxy " + proxy)
		execCmd("npm config set https-proxy " + proxy)
	}
}

func proxyBower(proxy string) {
	if isExisted("bower") {
		log.Println("Proxying bower...")

		bowerrcPath := os.Getenv("HOME") + "/.bowerrc"
		if _, err := os.Stat(bowerrcPath); os.IsNotExist(err) {
			os.Create(bowerrcPath)
		}

		count := 0
		lines, _ := readLines(bowerrcPath)
		for i, line := range lines {
			if strings.Index(line, "proxy") != -1 {
				lines[i] = "  \"proxy\": \"" + proxy + "\","
				if strings.Index(line, "https") != -1 {
					lines[i] = "  \"https-proxy\": \"" + proxy + "\","
				}

				count++
			}

			if count == 2 {
				break
			}
		}

		if len(lines) == 0 {
			lines = append(lines, "{")
			lines = append(lines, "  \"proxy\": \""+proxy+"\",")
			lines = append(lines, "  \"https-proxy\": \""+proxy+"\",")
			lines = append(lines, "}")
		}

		writeLines(lines, bowerrcPath)
	}
}

func proxyBashrc(proxy string) {
	log.Println("Proxying bashrc...")

	bashrcPath := os.Getenv("HOME") + "/.bashrc"
	if _, err := os.Stat(bashrcPath); os.IsNotExist(err) {
		os.Create(bashrcPath)
	}

	lines, _ := readLines(bashrcPath)

	lines = append(lines, "export http_proxy=\""+proxy+"\"")
	lines = append(lines, "export https_proxy=\"${http_proxy}\"")

	writeLines(lines, bashrcPath)
}

func proxySystemdDocker(proxy string) {
	log.Println("Proxying SystemdDocker...")

	dockerConfPath := "/etc/systemd/system/docker.service.d/docker.conf"
	file, err := os.OpenFile(dockerConfPath, os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Permission denied for writing to ", dockerConfPath)
		}
		return
	}
	defer file.Close()

	lines, _ := readLines(dockerConfPath)
	for i, line := range lines {
		if strings.Index(line, "Environment") != -1 {
			eqIndex := strings.Index(line, "=") + 1

			props := line[eqIndex:]
			props = regexp.MustCompile(`\s+`).ReplaceAllString(props, " ")

			a := strings.Split(props, " ")
			for i, prop := range a {
				if strings.Contains(strings.ToUpper(prop), "HTTP_PROXY") {
					a[i] = "\"HTTP_PROXY=" + proxy + "\" "
				}
				if strings.Contains(strings.ToUpper(prop), "HTTPS_PROXY") {
					a[i] = "\"HTTPS_PROXY=" + proxy + "\" "
				}
			}

			lines[i] = line[:eqIndex] + strings.Join(a, "")
		}
	}

	writeLines(lines, dockerConfPath)
	execCmd("systemctl daemon-reload")
}

func isExisted(cmd string) bool {
	_, err := exec.LookPath(cmd)
	if err != nil {
		return false
	}
	return true
}

func execCmd(cmd string) string {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return fmt.Sprintf("Failed to execute command: %s", cmd)
	}

	return string(out)
}

func readLines(path string) (lines [] string, err error) {
	var (
		file *os.File
	)

	if file, err = os.Open(path); err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return
}

func writeLines(lines [] string, path string) (err error) {
	var file *os.File

	if file, err = os.Create(path); err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, line := range lines {
		writer.WriteString(line + "\n")
	}

	return nil
}
