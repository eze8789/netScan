package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var ipAddress string
	var ports string
	var wrkrs int

	flag.StringVar(&ipAddress, "ip", "127.0.0.1", "Specify IP Address to scan.")
	flag.StringVar(&ports, "ports", "1-1024", "Specify port range to perform scan. Allowed formats: 80 or 1-1024")
	flag.IntVar(&wrkrs, "w", 8, "Amount of concurrent process scanning ports")
	flag.Parse()

	scanPorts, err := parsePorts(ports)
	if err != nil {
		log.Fatal(err)
	}

	pChann := make(chan int, wrkrs)
	openPorts := make(chan int)
	defer close(pChann)
	defer close(openPorts)

	for w := 0; w < cap(pChann); w++ {
		go scan(ipAddress, pChann, openPorts)
	}

	go populateChann(scanPorts, pChann)

	for i := 0; i < len(scanPorts); i++ {
		if p := <-openPorts; p != 0 {
			fmt.Printf("port is open: %d\n", p)
		}
	}
}
