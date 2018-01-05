package main

import (
	"log"
	"net"
	"time"
)

func main() {
	start := time.Now()
	conn, err := net.Dial("tcp", "localhost:9010")
	defer conn.Close()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("GetPlugins: %v", GetPlugins(conn))
	log.Printf("GetUniverseList: %v", GetUniverseList(conn))
	log.Printf("GetUniverseInfo: 1: %v", GetUniverseInfo(conn, 1))
	log.Printf("GetUniverseInfo: 3: %v", GetUniverseInfo(conn, 3))
	log.Printf("GetDmx: 1: %v", GetDmx(conn, 1).Data)
	log.Printf("GetDmx: 3: %v", GetDmx(conn, 3).Data)

	q := []byte{2, 3, 5, 7, 11, 13, 255, 0}
	log.Printf("SendDmx: 3: %v", SendDmx(conn, 3, q))

	elapsed := time.Since(start)
	log.Printf("test took %s", elapsed)
}
