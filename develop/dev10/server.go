package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func listenAndServe(conn net.Conn) {
	c := time.After(time.Minute) // для теста
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		select {
		case <-c:
			return
		default:
			fmt.Print("Receive message:", message)
			newMessage := strings.ToUpper(message)
			_, err := conn.Write([]byte(time.Now().String() + " " + newMessage + "\n"))
			if err != nil {
				log.Print(err)
			}
		}
	}
}

func main() {
	lis, _ := net.Listen("tcp", "localhost:8080")
	conn, _ := lis.Accept()

	listenAndServe(conn)

	err := lis.Close()
	if err != nil {
		log.Print(err)
	}
}
