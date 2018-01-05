package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/nickysemenza/gola/ola_proto"
	"log"
	"net"
)

//OlaClient holds connection info
type OlaClient struct {
	Conn    net.Conn
	Address string
}

//New creates a new OlaClient connecting to at the provided address
func New(address string) *OlaClient {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	return &OlaClient{
		Address: address,
		Conn:    conn,
	}
}

//Close closes the connection
func (o *OlaClient) Close() {
	o.Conn.Close()
}

//GetPlugins calls the GetPlugins RPC function
func (o *OlaClient) GetPlugins() (resp *ola_proto.PluginListReply, err error) {

	req := new(ola_proto.PluginListRequest)
	resp = new(ola_proto.PluginListReply)

	err = callRPCMethod(o.Conn, "GetPlugins", req, resp)
	return
}

//GetUniverseList calls the GetUniverseInfo RPC function
func (o *OlaClient) GetUniverseList() (resp *ola_proto.UniverseInfoReply, err error) {

	req := new(ola_proto.OptionalUniverseRequest)
	resp = new(ola_proto.UniverseInfoReply)

	err = callRPCMethod(o.Conn, "GetUniverseInfo", req, resp)
	return
}

//GetUniverseInfo calls the GetUniverseInfo RPC function, with the universe parameter
func (o *OlaClient) GetUniverseInfo(universe int) (resp *ola_proto.UniverseInfoReply, err error) {

	req := new(ola_proto.OptionalUniverseRequest)
	resp = new(ola_proto.UniverseInfoReply)

	req.Universe = proto.Int(universe)
	err = callRPCMethod(o.Conn, "GetUniverseInfo", req, resp)
	return
}

//GetDmx calls the GetDmx RPC function, with the universe parameter
func (o *OlaClient) GetDmx(universe int) (resp *ola_proto.DmxData, err error) {

	req := new(ola_proto.UniverseRequest)
	resp = new(ola_proto.DmxData)

	req.Universe = proto.Int(universe)
	err = callRPCMethod(o.Conn, "GetDmx", req, resp)
	return
}

//SendDmx calls the SendDmx RPC function, with the universe and data parameters
func (o *OlaClient) SendDmx(universe int, values []byte) (status bool, err error) {

	req := new(ola_proto.DmxData)
	resp := new(ola_proto.Ack)

	req.Universe = proto.Int(universe)
	req.Data = values
	err = callRPCMethod(o.Conn, "UpdateDmxData", req, resp)

	return true, err
}
