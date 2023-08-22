package details

import (
	"log"
	"net"
	"os"
)

func GetHostname() (hostname string, err error) {
	hostname, err = os.Hostname()
	return
}

func GetLocalIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP, err
}
