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

var host = flag.String("host", "", "A host host name of the Go engine")

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

	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			log.Fatal(err)
			wg.Done()
		}
	}()

	go func() {
		_, err :=
		go io.Copy(conn, os.Stdin)
		if err != nil {
			log.Fatal(err)
			wg.Done()
		}
	}()

	wg.Wait()
}
