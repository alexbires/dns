package main

import (
	"flag"
	"fmt"
	"github.com/abires/dns/header"
)

var (
	destination string
)

func main() {
	server := flag.String("ip", "8.8.8.8", "which server we are connecting to")
	flag.Parse()
	tmp := *server
	fmt.Println(tmp)
	header := &dnsHeader{}
	header.setID("duckduckgo.com")
	fmt.Println(header)
}
