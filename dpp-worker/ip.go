package main

import (
	"errors"
	"log"
	"net"
	"os"
)

// hostIP looks up the IP address of the local machine.
func hostIP() (string, error) {
	name, err := os.Hostname()
	if err != nil {
		log.Printf("hostIP: %v\n", err)
		return "", err
	}

	addrs, err := net.LookupHost(name)
	if err != nil {
		log.Printf("hostIP: %v\n", err)
		return "", err
	}

	for _, a := range addrs {
		return a, nil
	}
	return "", errors.New("Unable to resolve address")
}
