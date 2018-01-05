package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/nickysemenza/gola/ola_proto"
	"log"
	"net"
)

type OlaClient struct {
	Conn    net.Conn
	Address string
}

//Create a new OlaClient connecting to at the provided address
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

//Closes the connection
func (o *OlaClient) Close() {
	o.Conn.Close()
}

//Calls the GetPlugins RPC function
func (o *OlaClient) GetPlugins() (resp *ola_proto.PluginListReply, err error) {

	req := new(ola_proto.PluginListRequest)
	resp = new(ola_proto.PluginListReply)

	err = callRpcMethod(o.Conn, "GetPlugins", req, resp)
	return
}

//Calls the GetUniverseInfo RPC function
func (o *OlaClient) GetUniverseList() (resp *ola_proto.UniverseInfoReply, err error) {

	req := new(ola_proto.OptionalUniverseRequest)
	resp = new(ola_proto.UniverseInfoReply)

	err = callRpcMethod(o.Conn, "GetUniverseInfo", req, resp)
	return
}

//Calls the GetUniverseInfo RPC function, with the universe parameter
func (o *OlaClient) GetUniverseInfo(universe int) (resp *ola_proto.UniverseInfoReply, err error) {

	req := new(ola_proto.OptionalUniverseRequest)
	resp = new(ola_proto.UniverseInfoReply)

	req.Universe = proto.Int(universe)
	err = callRpcMethod(o.Conn, "GetUniverseInfo", req, resp)
	return
}

//Calls the GetDmx RPC function, with the universe parameter
func (o *OlaClient) GetDmx(universe int) (resp *ola_proto.DmxData, err error) {

	req := new(ola_proto.UniverseRequest)
	resp = new(ola_proto.DmxData)

	req.Universe = proto.Int(universe)
	err = callRpcMethod(o.Conn, "GetDmx", req, resp)
	return
}

//Calls the SendDmx RPC function, with the universe and data parameters
func (o *OlaClient) SendDmx(universe int, values []byte) (status bool, err error) {

	req := new(ola_proto.DmxData)
	resp := new(ola_proto.Ack)

	req.Universe = proto.Int(universe)
	req.Data = values
	err = callRpcMethod(o.Conn, "UpdateDmxData", req, resp)

	return true, err
}
