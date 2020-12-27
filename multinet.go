package main

import (
	"fmt"
	"net"
	"os"
)


func main() {
	ln, err := net.Listen("tcp", ":4535")
		if err != nil {
			fmt.Println("failed to bind port:", err)
			os.Exit(1)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("failed to handlet connection:", err)
			os.Exit(1)
		}
		go connectionHandler(conn)
	}

}

func connectionHandler(c net.Conn) {
	text := "YOU HAVE CONNECTED\n"
	_, err := c.Write([]byte(text))
	if err != nil {
		fmt.Println("failed to write:", err)
	}
	fmt.Println("Connection from %s:", c.String)
	err = c.Close()
	if err != nil {
                fmt.Println("failed to close connection:", err)
        }

}

