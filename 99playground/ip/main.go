package main

/*
URL: https://github.com/mccoyst/myip/blob/master/myip.go
URL: http://changsijay.com/2013/07/28/golang-get-ip-address/
*/

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}

	iface, _ := net.InterfaceByName("本地连接")
	if iface != nil {
		addrs, _ := iface.Addrs()
		if len(addrs) == 1 {
			fmt.Printf("%v\n", addrs[0])
		}
		if len(addrs) == 2 {
			fmt.Printf("%v\n", addrs[1])
		}
	}

	list, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for i, iface := range list {
		fmt.Printf("%d name=%s %v\n", i, iface.Name, iface)
		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}

		for j, addr := range addrs {
			fmt.Printf(" %d %v\n", j, addr)
		}
	}
}
