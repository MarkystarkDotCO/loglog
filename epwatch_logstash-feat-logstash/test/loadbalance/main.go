package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	var text string = "a"
	for i := 0; i < 1250; i++ {
		text = text + "a"
	}

	servAddr := "0.0.0.0:6514"

	conn, err := net.Dial("udp", servAddr)
	defer conn.Close()

	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
	for {
		fmt.Fprintf(conn, text)
		// p := make([]byte, 2048)
		// _, err = bufio.NewReader(conn).Read(p)
		// if err == nil {
		// 	fmt.Printf("%s\n", p)
		// } else {
		// 	fmt.Printf("Some error %v\n", err)
		// }
	}

	// conn.Close()
}
