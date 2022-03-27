package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

//func to handle error
func loggFatal(err error) {
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
}

//Client server
func main() {
	con, err := net.Dial("tcp", ":8888")
	loggFatal(err)
	defer con.Close()

	clientReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name? ")
	msg, _ := clientReader.ReadString('\n')
	fmt.Println("welcome to the room", msg)
	fmt.Println("-----------------------------------------")
	//User = msg

	go Reader(con)
	Writer(con, msg)

}

func Reader(con net.Conn) string {
	toRead := ""
	if con != nil {
		toRead += "network reading client message"

		for {
			data := make([]byte, 4390)
			read, err := con.Read(data)
			if err == io.EOF {
				fmt.Println(err)
				con.Close()
				os.Exit(0)
			}
			fmt.Println(string(data[:read]))
			fmt.Println("-----------------------------------------")
		}
	} else {
		toRead += "network not reading message"
	}
	return toRead
}

func Writer(con net.Conn, user string) string {
	toWrite := ""
	if con != nil || user != "" {
		toWrite += "network writing message"

		for {
			clientReader := bufio.NewReader(os.Stdin)
			msg, err := clientReader.ReadString('\n')
			msg = strings.TrimSpace(msg)
			user = strings.Trim(user, "\n")
			if strings.TrimSpace(msg) == "STOP" {
				fmt.Println(user + " exiting...")
				con.Write([]byte(user + " exiting..."))
				return ""
			}
			if _, err = con.Write([]byte(user + " --> " + msg + "\n")); err != nil {
				log.Printf("failed to send the client request: %v\n", err)
			}
		}
	} else {
		toWrite += "network not writing message"
	}
	return toWrite
}
