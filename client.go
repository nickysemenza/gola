package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/nickysemenza/gola/ola_proto"
	"net"
)

func GetPlugins(conn net.Conn) (resp *ola_proto.PluginListReply, err error) {

	req := new(ola_proto.PluginListRequest)
	resp = new(ola_proto.PluginListReply)

	err = callRpcMethod(conn, "GetPlugins", req, resp)
	return
}
func GetUniverseList(conn net.Conn) (resp *ola_proto.UniverseInfoReply, err error) {

	req := new(ola_proto.OptionalUniverseRequest)
	resp = new(ola_proto.UniverseInfoReply)

	err = callRpcMethod(conn, "GetUniverseInfo", req, resp)
	return
}
func GetUniverseInfo(conn net.Conn, universe int) (resp *ola_proto.UniverseInfoReply, err error) {

	req := new(ola_proto.OptionalUniverseRequest)
	resp = new(ola_proto.UniverseInfoReply)

	req.Universe = proto.Int(universe)
	err = callRpcMethod(conn, "GetUniverseInfo", req, resp)
	return
}

func GetDmx(conn net.Conn, universe int) (resp *ola_proto.DmxData, err error) {

	req := new(ola_proto.UniverseRequest)
	resp = new(ola_proto.DmxData)

	req.Universe = proto.Int(universe)
	err = callRpcMethod(conn, "GetDmx", req, resp)
	return
}

func SendDmx(conn net.Conn, universe int, values []byte) (status bool, err error) {

	req := new(ola_proto.DmxData)
	resp := new(ola_proto.Ack)

	req.Universe = proto.Int(universe)
	req.Data = values
	err = callRpcMethod(conn, "UpdateDmxData", req, resp)

	return true, err
}
