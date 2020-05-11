package network

import (
	"fmt"
	"net"
)

func ConnectToDNSServer(destination string, data []byte) {
	conn, err := net.Dial("udp", destination)

	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	conn.Write(data)
}
