package main

import (
	"net"
	"github.com/nickysemenza/gola/ola_proto"
	"github.com/golang/protobuf/proto"
)

func GetPlugins(conn net.Conn) *ola_proto.PluginListReply {

	req := new(ola_proto.PluginListRequest)
	resp := new(ola_proto.PluginListReply)

	callRpcMethod(conn, "GetPlugins", req, resp)
	return resp
}
func GetUniverseList(conn net.Conn) *ola_proto.UniverseInfoReply {

	req := new(ola_proto.OptionalUniverseRequest)
	resp := new(ola_proto.UniverseInfoReply)

	callRpcMethod(conn, "GetUniverseInfo", req, resp)
	return resp
}
func GetUniverseInfo(conn net.Conn, universe int) *ola_proto.UniverseInfoReply {

	req := new(ola_proto.OptionalUniverseRequest)
	resp := new(ola_proto.UniverseInfoReply)

	req.Universe = proto.Int(universe)
	callRpcMethod(conn, "GetUniverseInfo", req, resp)

	return resp
}

func GetDmx(conn net.Conn, universe int) *ola_proto.DmxData {

	req := new(ola_proto.UniverseRequest)
	resp := new(ola_proto.DmxData)

	req.Universe = proto.Int(universe)
	callRpcMethod(conn, "GetDmx", req, resp)

	return resp
}

func SendDmx(conn net.Conn, universe int, values []byte) bool {

	req := new(ola_proto.DmxData)
	resp := new(ola_proto.Ack)

	req.Universe = proto.Int(universe)
	req.Data = values
	callRpcMethod(conn, "UpdateDmxData", req, resp)

	return true
}