# gola
[WIP] Open Lighting Project client for golang using RPC over sockets

[![Go Report Card](https://goreportcard.com/badge/github.com/nickysemenza/gola)](https://goreportcard.com/report/github.com/nickysemenza/gola)
[![GoDoc](https://godoc.org/github.com/nickysemenza/gola?status.svg)](https://godoc.org/github.com/nickysemenza/gola)

Example:
```
client := New("localhost:9010")
defer client.Close()
# get DMX on universe 1

if x, err := client.GetDmx(1); err != nil {
	log.Printf("GetDmx: 1: %v", err)
}  else {
	log.Printf("GetDmx: 1: %v", x.Data)
}
```
