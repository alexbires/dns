package main

import (
	"flag"
	"fmt"

	"github.com/abires/dns/header"
	"github.com/abires/dns/network"
)

var (
	destination string
)

func main() {
	server := flag.String("ip", "8.8.8.8", "which server we are connecting to")
	flag.Parse()
	destination = *server + ":53"
	header := &header.DNSHeader{}
	header.SetID()
	header.SetQuery(true)
	header.SetNumberofQuestions(1)
	header.SetNumberofNameServers(0)
	header.SetNumberofAdditional(0)
	//fmt.Println(header)
	headerBuffer := header.ToByteBuffer()
	byteArray := headerBuffer.Bytes()
	//stream := []byte{145, 150}
	fmt.Println(byteArray)
	network.ConnectToDNSServer(destination, byteArray)
}
