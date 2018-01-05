package main

import (
	"net"
	"log"
	"github.com/golang/protobuf/proto"
	"github.com/nickysemenza/gola/ola_rpc"
	"fmt"
	"os"
	"encoding/binary"
	"github.com/pkg/errors"
)

const (
	PROTOCOLVERSION = 1
	VERSIONMASK     = 0xf0000000
	SIZEMASK        = 0x0fffffff
)

func decipherMessage(data []byte, whereTo proto.Message) error {
	rpcMessage := new(ola_rpc.RpcMessage)
	if err := proto.Unmarshal(data, rpcMessage); err != nil {
		log.Fatalln("Failed to parse ola_rpc.RpcMessage resp:", err)
		return err
	}

	switch *rpcMessage.Type {
	case ola_rpc.Type_RESPONSE_FAILED:
		return errors.New("response failed")
	}

	innerBuffer := rpcMessage.GetBuffer()
	if err := proto.Unmarshal(innerBuffer, whereTo); err != nil {
		log.Fatalln("Failed to parse inner resp:", err)
		return err
	}
	return nil
}

func sendMessage(conn net.Conn, pb proto.Message, rpcFunction string) {

	dataToSend, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("couldn't marshal inner pb", err)
	}

	var t ola_rpc.Type
	t = ola_rpc.Type_REQUEST

	rpcMessage := new(ola_rpc.RpcMessage)
	rpcMessage.Type = &t
	rpcMessage.Id = proto.Uint32(1)
	rpcMessage.Name = proto.String(rpcFunction)
	rpcMessage.Buffer = dataToSend


	encodedRpcMessage, err := proto.Marshal(rpcMessage)
	if err != nil {
		log.Fatalln("couldn't marshal outer pb", err)
	}
	//log.Printf("%v", rpcMessage)
	//log.Printf("%v", encodedRpcMessage)

	sendDataToDest(conn, encodedRpcMessage)
}


func callRpcMethod(conn net.Conn, rpcFunction string, pb proto.Message, pb2 proto.Message) error {
	sendMessage(conn, pb, rpcFunction)

	rsp := readData(conn)

	return decipherMessage(rsp,pb2)
}

func readData(conn net.Conn) []byte {


	header := make([]byte, 4)
	conn.Read(header)

	headerValue := int(binary.LittleEndian.Uint32(header))
	size := headerValue  & SIZEMASK
	//log.Printf("expecing %d bytes",size)


	data := make([]byte, size)
	_, err:= conn.Read(data)
	if err != nil {
		log.Fatalln(err)
	}
	//log.Printf("received %d bytes",n)
	return data
}

//func sendDataToDest(data []byte, dst *string) {
func sendDataToDest(conn net.Conn, data []byte) {

	headerContent := (PROTOCOLVERSION << 28) & VERSIONMASK
	headerContent |= len(data) & SIZEMASK
	//log.Printf("header: %v", headerContent)


	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(headerContent))


	conn.Write(bs)

	_	, err := conn.Write(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return
	}
	//fmt.Println("Sent " + strconv.Itoa(n) + " bytes")


}