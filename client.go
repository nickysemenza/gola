package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/nickysemenza/gola/ola_proto"
	"net"
	"log"
)

type OlaClient struct {
	Conn net.Conn
	Address string
}

func New(address string) *OlaClient {
	conn, err := net.Dial("tcp", "localhost:9010")
	if err != nil {
		log.Fatal(err)
	}
	return &OlaClient{
		Address: address,
		Conn: conn,
	}
}
func (o *OlaClient) Close() {
	o.Conn.Close()
}
func (o *OlaClient) GetPlugins() (resp *ola_proto.PluginListReply, err error) {

	req := new(ola_proto.PluginListRequest)
	resp = new(ola_proto.PluginListReply)

	err = callRpcMethod(o.Conn, "GetPlugins", req, resp)
	return
}
func (o *OlaClient) GetUniverseList() (resp *ola_proto.UniverseInfoReply, err error) {

	req := new(ola_proto.OptionalUniverseRequest)
	resp = new(ola_proto.UniverseInfoReply)

	err = callRpcMethod(o.Conn, "GetUniverseInfo", req, resp)
	return
}
func (o *OlaClient) GetUniverseInfo( universe int) (resp *ola_proto.UniverseInfoReply, err error) {

	req := new(ola_proto.OptionalUniverseRequest)
	resp = new(ola_proto.UniverseInfoReply)

	req.Universe = proto.Int(universe)
	err = callRpcMethod(o.Conn, "GetUniverseInfo", req, resp)
	return
}

func (o *OlaClient) GetDmx( universe int) (resp *ola_proto.DmxData, err error) {

	req := new(ola_proto.UniverseRequest)
	resp = new(ola_proto.DmxData)

	req.Universe = proto.Int(universe)
	err = callRpcMethod(o.Conn, "GetDmx", req, resp)
	return
}

func (o *OlaClient) SendDmx(universe int, values []byte) (status bool, err error) {

	req := new(ola_proto.DmxData)
	resp := new(ola_proto.Ack)

	req.Universe = proto.Int(universe)
	req.Data = values
	err = callRpcMethod(o.Conn, "UpdateDmxData", req, resp)

	return true, err
}
