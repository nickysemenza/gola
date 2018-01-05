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

	if x, err := GetPlugins(conn); err != nil {
		log.Printf("GetPlugins: %v", err)
	}  else {
		log.Printf("GetPlugins: %v", x)
	}

	if x, err := GetUniverseList(conn); err != nil {
		log.Printf("GetUniverseList: %v", err)
	}  else {
		log.Printf("GetUniverseList: %v", x)
	}

	if x, err := GetUniverseInfo(conn, 1); err != nil {
		log.Printf("GetUniverseInfo: 1: %v", err)
	}  else {
		log.Printf("GetUniverseInfo: 1: %v", x)
	}

	if x, err := GetUniverseInfo(conn, 3); err != nil {
		log.Printf("GetUniverseInfo: 3: %v", err)
	}  else {
		log.Printf("GetUniverseInfo: 3: %v", x)
	}

	if x, err := GetDmx(conn, 1); err != nil {
		log.Printf("GetDmx: 1: %v", err)
	}  else {
		log.Printf("GetDmx: 1: %v", x.Data)
	}

	if x, err := GetDmx(conn, 3); err != nil {
		log.Printf("GetDmx: 3: %v", err)
	}  else {
		log.Printf("GetDmx: 3: %v", x.Data)
	}

	if x, err := SendDmx(conn, 1, []byte{2, 3, 5, 7, 11, 13, 255, 0}); err != nil {
		log.Printf("SendDmx: 1: %v", err)
	}  else {
		log.Printf("SendDmx: 1: %v", x)
	}

	if x, err := SendDmx(conn, 3, []byte{2, 3, 5, 7, 11, 13, 255, 0}); err != nil {
		log.Printf("SendDmx: 3: %v", err)
	}  else {
		log.Printf("SendDmx: 3: %v", x)
	}

	elapsed := time.Since(start)
	log.Printf("test took %s", elapsed)
}
