package details

import (
	"log"
	"net"
	"os"
)

func GetHostName() (string, error) {
	return os.Hostname()
}

func GetIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAdder := conn.LocalAddr().(*net.UDPAddr)
	return localAdder.IP.String()
}
