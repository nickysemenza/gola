package main

import (
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/nickysemenza/gola/ola_rpc"
	"github.com/pkg/errors"
	"log"
	"net"
	"os"
)

const (
	protocolVersion = 1
	versionMask     = 0xf0000000
	sizeMask        = 0x0fffffff
)

//Deciphers an incoming message by unwrapping the outer ola_rpc.RpcMessage first
func decipherMessage(data []byte, whereTo proto.Message) error {
	rpcMessage := new(ola_rpc.RpcMessage)
	if err := proto.Unmarshal(data, rpcMessage); err != nil {
		log.Fatalln("Failed to parse ola_rpc.RpcMessage resp:", err)
		return err
	}

	switch *rpcMessage.Type {
	case ola_rpc.Type_RESPONSE_FAILED:
		//Buffer now probably contain an error msg
		str := fmt.Sprintf("%s", rpcMessage.Buffer)
		return errors.New("RESPONSE_FAILED: " + str)
	}

	innerBuffer := rpcMessage.GetBuffer()
	if err := proto.Unmarshal(innerBuffer, whereTo); err != nil {
		log.Fatalln("Failed to parse inner resp:", err)
		return err
	}
	return nil
}

//Sends a message to an RPC function
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

	encodedRPCMessage, err := proto.Marshal(rpcMessage)
	if err != nil {
		log.Fatalln("couldn't marshal outer pb", err)
	}
	//log.Printf("%v", rpcMessage)
	//log.Printf("%v", encodedRPCMessage)

	sendDataToDest(conn, encodedRPCMessage)
}

//callRPCMethod calls an RPC message, unpacking the response into resp proto.Message
func callRPCMethod(conn net.Conn, rpcFunction string, pb proto.Message, responseMessage proto.Message) error {
	sendMessage(conn, pb, rpcFunction)
	rsp := readData(conn)

	return decipherMessage(rsp, responseMessage)
}

//Reads the 4 bytes of header and then body, returns the body
func readData(conn net.Conn) []byte {

	header := make([]byte, 4)
	conn.Read(header)

	headerValue := int(binary.LittleEndian.Uint32(header))
	size := headerValue & sizeMask
	//log.Printf("expecing %d bytes",size)

	data := make([]byte, size)
	_, err := conn.Read(data)
	if err != nil {
		log.Fatalln(err)
	}
	//log.Printf("received %d bytes",n)
	return data
}

//Sends data over the connection, pre-pending it with a 4 byte header
func sendDataToDest(conn net.Conn, data []byte) {

	headerContent := (protocolVersion << 28) & versionMask
	headerContent |= len(data) & sizeMask
	//log.Printf("header: %v", headerContent)

	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(headerContent))

	conn.Write(bs)

	_, err := conn.Write(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return
	}
	//fmt.Println("Sent " + strconv.Itoa(n) + " bytes")

}
