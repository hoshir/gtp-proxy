package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

var host = flag.String("host", "", "A host")

func main() {
	flag.Parse()

	time.Sleep(5 * time.Second)

	dest := fmt.Sprintf("%s:6000", *host)
	conn, err := net.Dial("tcp", dest)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	// Receive loop
	go io.Copy(os.Stdout, conn)

	// Send loop
	go io.Copy(conn, os.Stdin)

	wg.Wait()
}
