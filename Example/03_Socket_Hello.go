
package main

import (
	"fmt"
	"net"
)

func doServerStuff(conn net.Conn) {
	fmt.Println("Accept a NEW Connect from client ...")
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		fmt.Printf("\n Received data: %v", string(buf[:len]))
	}
}

func main() {
	fmt.Println("Starting the Server at localhost:50000 ...")

	//Create listener
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error Listening", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		go doServerStuff(conn)
	}

}
