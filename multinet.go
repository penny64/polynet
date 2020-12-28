package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"io"
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
	MAGIC := "YOU HAVE CONNECTED\n"
	var polymap map[float64]float64
	polymap = make(map[float64]float64)
	_, err := c.Write([]byte(MAGIC))
	if err != nil {
		fmt.Println("failed to write:", err)
	}
	fmt.Println("Connection from:", c.RemoteAddr())
	for ;;{
		poly, err := polynominalRead(c)
		if err == io.EOF {
			fmt.Println("Connection closed")
			break
		} else if err != nil {
			fmt.Println("failed to parse input:", err)
			continue
			}
		go networkReply(c, processPoly(poly, polymap))
			if err != nil {
			fmt.Println("failed to send reply:", err)
		}

	}
	err = c.Close()
	if err != nil {
                fmt.Println("failed to close connection:", err)
        }

}


func networkReply(c net.Conn, poly float64) (err error){
	message := strconv.FormatFloat(poly, 'f', -1, 64)
	_, err = c.Write([]byte(message))
	if err != nil {
		return
        }
	return
}

func polynominalRead(c net.Conn) (poly float64, err error) {
        buffer := make([]byte, 1024)
        DATA, err := c.Read(buffer)
	if err != nil {
		return
        }

	poly, err = strconv.ParseFloat(string(buffer[:DATA-1]), 64)
	return
}

//The server should calculate the value of Y where Y = -1/2*X^2 + X + 8 using recursion.

func processPoly(poly float64, polymap map[float64]float64) (polyresult float64) {
	polyresult, ok := polymap[poly]
	if !ok {
		polyresult = recursivePolynominal(poly)
		polymap[poly] = polyresult
	}
	return
}
func recursivePolynominal(X float64) (Y float64) {
	Y = (-.5*(X*X))+8
	return
}
