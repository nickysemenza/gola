/*
Package gola is a Open Lighting client for golang.

It communicates over sockets/RPC/protobuf to an ola server.
*/
package gola

import (
	"net"

	"github.com/golang/protobuf/proto"
	"github.com/nickysemenza/gola/ola_proto"
)

//Client holds connection info
type Client struct {
	Conn    net.Conn
	Address string
}

//New creates a new OlaClient connecting to at the provided address
func New(address string) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &Client{
		Address: address,
		Conn:    conn,
	}, nil
}

//Close closes the connection
func (c *Client) Close() {
	c.Conn.Close()
}

//GetPlugins calls the GetPlugins RPC function
func (c *Client) GetPlugins() (resp *ola_proto.PluginListReply, err error) {

	req := new(ola_proto.PluginListRequest)
	resp = new(ola_proto.PluginListReply)

	err = c.callRPCMethod("GetPlugins", req, resp)
	return
}

//GetUniverseList calls the GetUniverseInfo RPC function
func (c *Client) GetUniverseList() (resp *ola_proto.UniverseInfoReply, err error) {

	req := new(ola_proto.OptionalUniverseRequest)
	resp = new(ola_proto.UniverseInfoReply)

	err = c.callRPCMethod("GetUniverseInfo", req, resp)
	return
}

//GetUniverseInfo calls the GetUniverseInfo RPC function, with the universe parameter
func (c *Client) GetUniverseInfo(universe int) (resp *ola_proto.UniverseInfoReply, err error) {

	req := new(ola_proto.OptionalUniverseRequest)
	resp = new(ola_proto.UniverseInfoReply)

	req.Universe = proto.Int(universe)
	err = c.callRPCMethod("GetUniverseInfo", req, resp)
	return
}

//GetDmx calls the GetDmx RPC function, with the universe parameter
func (c *Client) GetDmx(universe int) (resp *ola_proto.DmxData, err error) {

	req := new(ola_proto.UniverseRequest)
	resp = new(ola_proto.DmxData)

	req.Universe = proto.Int(universe)
	err = c.callRPCMethod("GetDmx", req, resp)
	return
}

//SendDmx calls the SendDmx RPC function, with the universe and data parameters
func (c *Client) SendDmx(universe int, values []byte) (status bool, err error) {

	req := new(ola_proto.DmxData)
	resp := new(ola_proto.Ack)

	req.Universe = proto.Int(universe)
	req.Data = values
	err = c.callRPCMethod("UpdateDmxData", req, resp)

	return true, err
}
