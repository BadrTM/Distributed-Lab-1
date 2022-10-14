package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func readServerReply(conn net.Conn) {
	reader := bufio.NewReader(conn)
	msg, _ := reader.ReadString('\n')
	fmt.Println(msg)
}

func main() {
	// Getting command line arguments
	//msgP := flag.String("msg", "Default Msg: Msg from client", "The msg you want to sent")
	//flag.Parse()
	stdIn := bufio.NewReader(os.Stdin)
	conn, _ := net.Dial("tcp", ":8030")

	for {
		fmt.Print("Send Message -> ")
		msg, _ := stdIn.ReadString('\n')
		fmt.Fprintln(conn, msg)
		readServerReply(conn)
	}
}
