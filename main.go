package main

import "net"

import "fmt"

import "log"

func main() {
	port, err := net.LookupPort("tcp", "ssh")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(port)

	addrs, _ := net.LookupHost("www.facebook.com")
	fmt.Println(addrs)

	ints, _ := net.Interfaces()
	fmt.Println(ints)
}
