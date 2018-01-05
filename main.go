package main

import (
	"net"
	"log"
	"github.com/nickysemenza/gola/ola_proto"
	"github.com/golang/protobuf/proto"
	"github.com/nickysemenza/gola/ola_rpc"
	"fmt"
	"os"
	"strconv"
	"encoding/binary"
	"time"
	"github.com/pkg/errors"
)

const (
	PROTOCOLVERSION = 1
	VERSIONMASK     = 0xf0000000
	SIZEMASK        = 0x0fffffff
)
func main() {

	conn, err := net.Dial("tcp", "localhost:9010")
	defer conn.Close()
	if err != nil {
		log.Fatalln(err)
	}
	start := time.Now()


	test1 := new(ola_proto.OptionalUniverseRequest)
	//test1.Universe = proto.Int(33)
	sendMessage(conn, test1,"GetUniverseInfo")

	rsp := readData(conn)
	innards := new(ola_proto.UniverseInfoReply)

	if err := decipherMessage(rsp,innards); err != nil {
		log.Fatalln("error deciphering:", err)
	} else {
		log.Printf("yay: %v", innards)
	}




	elapsed := time.Since(start)
	log.Printf("test took %s", elapsed)
}

func decipherMessage(data []byte, whereTo proto.Message) error {
	rpcMessage := new(ola_rpc.RpcMessage)
	if err := proto.Unmarshal(data, rpcMessage); err != nil {
		log.Fatalln("Failed to parse ola_rpc.RpcMessage resp:", err)
		return err
	}

	switch *rpcMessage.Type {
	case ola_rpc.Type_RESPONSE:
		log.Printf("good to go")
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


//public PluginListReply getPlugins() {
//return (PluginListReply) callRpcMethod("GetPlugins", PluginListRequest.newBuilder().build());
//}
func getPlugins(conn net.Conn) *ola_proto.PluginListReply {

	x := new(ola_proto.PluginListRequest)
	callRpcMethod("GetPlugins",x)
	return nil
}
func callRpcMethod(s string, request *ola_proto.PluginListRequest) {

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
	log.Printf("header: %v", headerContent)


	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(headerContent))
	log.Println(bs)


	//conn, err := net.Dial("tcp", "localhost:9010")
	////conn, err := net.Dial("tcp", *dst)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	//	return
	//}

	conn.Write(bs)

	n, err := conn.Write(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return
	}
	fmt.Println("Sent " + strconv.Itoa(n) + " bytes")


}