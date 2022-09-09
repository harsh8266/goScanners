package main

import (
	"fmt"
	"net"
)

func main() {

	for i := 0; i <= 1023; i++ {
		address := fmt.Sprintf("192.168.1.1:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port is closed or filtered.
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
