package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.1:443")
	if err == nil {

		fmt.Println("Connection open")
		conn.Close()
	}

}
