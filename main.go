package main

import "syscall"

import "os"

import "fmt"

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

	/* TCP CLIENT
	fmt.Println(os.Getenv("PROJECT_NAME"))
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:8080")
	tcpConn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer tcpConn.Close()
	buffer := make([]byte, 1024)
	buffer = []byte("HEAD / HTTP/1.1\r\n\r\n")
	tcpConn.Write(buffer)
	*/

	/* IP SERVER
	ipAddr, _ := net.ResolveIPAddr("ip4", "localhost")
	ipConn, err := net.ListenIP("ip4:icmp", ipAddr)
	if err != nil {
		log.Fatalln(err)
	}
	//defer ipConn.Close()
	buffer := make([]byte, 1024)
	for {
		n, _, err := ipConn.ReadFrom(buffer)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(buffer[:n])
	}
	*/

	fd, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	file := os.NewFile(uintptr(fd), "")

	buffer := make([]byte, 1024*4)
	for {
		n, _ := file.Read(buffer)
		fmt.Println(buffer[:n])
	}
}
