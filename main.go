package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	var ipAddress string
	var ports string

	flag.StringVar(&ipAddress, "ip", "127.0.0.1", "Specify IP Address to scan.")
	flag.StringVar(&ports, "ports", "1-1024", "Specify port range to perform scan. Allowed formats: 80 or 1-1024")
	flag.Parse()

	scanPorts, err := parsePorts(ports)
	if err != nil {
		log.Fatal(err)
	}

	for i := range scanPorts {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ipAddress, scanPorts[i]))
		if err != nil {
			fmt.Printf("port is closed: %d\n", scanPorts[i])
			continue
		}
		conn.Close()
		fmt.Printf("port is open: %d\n", scanPorts[i])
	}
}

// parsePorts return a slice of integers and an error after parse an string with a port range (i.e: single port: 80 multiple ports: 1-1024)
func parsePorts(ports string) ([]int, error) {
	portRanges := strings.Split(ports, "-")
	if len(portRanges) > 2 {
		return nil, errors.New("unable to parse ports")
	}

	if len(portRanges) == 1 {
		p, err := strconv.Atoi(portRanges[0])
		if err != nil {
			return nil, err
		}
		return []int{p}, nil
	}

	fp, err := strconv.Atoi(portRanges[0])
	if err != nil {
		return nil, err
	}
	tp, err := strconv.Atoi(portRanges[1])
	if err != nil {
		return nil, err
	}

	if fp <= 0 || tp <= 0 {
		return nil, errors.New("ports needs to be greater than 0")
	}

	var values []int

	for p := fp; p <= tp; p++ {
		values = append(values, p)
	}
	return values, nil
}
