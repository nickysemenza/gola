package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()
	client := New("localhost:9010")
	defer client.Close()

	if x, err := client.GetPlugins(); err != nil {
		log.Printf("GetPlugins: %v", err)
	} else {
		log.Printf("GetPlugins: %v", x)
	}

	if x, err := client.GetUniverseList(); err != nil {
		log.Printf("GetUniverseList: %v", err)
	} else {
		log.Printf("GetUniverseList: %v", x)
	}

	if x, err := client.GetUniverseInfo(1); err != nil {
		log.Printf("GetUniverseInfo: 1: %v", err)
	} else {
		log.Printf("GetUniverseInfo: 1: %v", x)
	}

	if x, err := client.GetUniverseInfo(3); err != nil {
		log.Printf("GetUniverseInfo: 3: %v", err)
	} else {
		log.Printf("GetUniverseInfo: 3: %v", x)
	}

	if x, err := client.GetDmx(1); err != nil {
		log.Printf("GetDmx: 1: %v", err)
	} else {
		log.Printf("GetDmx: 1: %v", x.Data)
	}

	if x, err := client.GetDmx(3); err != nil {
		log.Printf("GetDmx: 3: %v", err)
	} else {
		log.Printf("GetDmx: 3: %v", x.Data)
	}

	if x, err := client.SendDmx(1, []byte{2, 3, 5, 7, 11, 13, 255, 0}); err != nil {
		log.Printf("SendDmx: 1: %v", err)
	} else {
		log.Printf("SendDmx: 1: %v", x)
	}

	if x, err := client.SendDmx(3, []byte{2, 3, 5, 7, 11, 13, 255, 0}); err != nil {
		log.Printf("SendDmx: 3: %v", err)
	} else {
		log.Printf("SendDmx: 3: %v", x)
	}

	elapsed := time.Since(start)
	log.Printf("test took %s", elapsed)
}
