package main

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// TODO: add UDP support
func scan(ipAddress, output string, pChann <-chan int, openPorts chan<- int) {
	for v := range pChann {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ipAddress, v))
		if err != nil {
			writeOutput(output, ipAddress, "closed", v)
			openPorts <- 0
			continue
		}
		conn.Close()
		openPorts <- v
	}
}

func populateChann(n []int, c chan int) {
	for _, v := range n {
		c <- v
	}
}

// parsePorts return a slice of integers and an error after parse an string with a port range (i.e: single port: 80 multiple ports: 1-1024)
func parsePorts(ports string) ([]int, error) {
	portRanges := strings.Split(ports, "-")
	//nolint:gomnd
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
