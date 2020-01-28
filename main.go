package main

import "net"

func main() {
	// port, err := net.LookupPort("tcp", "ssh")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(port)

	// addrs, _ := net.LookupHost("www.facebook.com")
	// fmt.Println(addrs)

	// ints, _ := net.Interfaces()
	// fmt.Println(ints)

	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:8080")
	tcpConn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer tcpConn.Close()
	buffer := make([]byte, 1024)
	buffer = []byte("/ GET HTTP1.1 \r\n\r\n")
	tcpConn.Write(buffer)
}
