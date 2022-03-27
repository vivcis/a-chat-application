package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

var (
	openConnection = make(map[net.Conn]bool)
	deadConnection = make(chan net.Conn)
)

//var User string

func main() {
	// Start the server and listen for incoming connections.
	log.Printf("Server started on port :8888")

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	// Close the listener when the application closes.
	defer listener.Close()
	log.Println("waiting for client...")

	// run loop forever, until exit.
	for {
		con, err := listener.Accept()
		if err != nil {
			panic(err)

		}

		if strings.TrimSpace(con.RemoteAddr().String()) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}
		fmt.Println("Client " + con.RemoteAddr().String() + " accepted.")
		openConnection[con] = true
		go BroadcastMessage(con)
	}
}

func BroadcastMessage(conn net.Conn) string {
	toBroadcast := ""
	if conn != nil {
		toBroadcast += "message broad-casted to other clients"

		data := make([]byte, 4390)
		for {
			read, err := conn.Read(data)
			if err != nil {
				//log.Print(err)
				return ""
			}

			for item := range openConnection {
				if item != conn {
					daytime := time.Now().Format(time.RFC822)
					item.Write(append([]byte(daytime+": "), data[:read]...))

				}
			}
			if strings.Contains(string(data), "exiting...") {
				fmt.Println(string(data))
				return ""
			}
		}

	} else {
		toBroadcast += "message not broad-casted to other clients"
	}
	deadConnection <- conn
	return toBroadcast
}
