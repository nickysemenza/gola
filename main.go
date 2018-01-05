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
)

const (
	PROTOCOL_VERSION = 1
	VERSION_MASK = 0xf0000000
	SIZE_MASK = 0x0fffffff
)
func main() {

	conn, err := net.Dial("tcp", "localhost:9010")
	defer conn.Close()
	if err != nil {
		log.Fatalln(err)
	}
	start := time.Now()


	//test1 := new(ola_proto.Ack)
	test1 := new(ola_proto.OptionalUniverseRequest)
	//test1.Universe = proto.Int32(2)


	dataToSend, err := proto.Marshal(test1)
	if err != nil {
		log.Fatalln(err)
	}
	//log.Printf("%v",dataToSend)
	//log.Printf("%v",test1)
	log.Println("-----------")

	var test ola_rpc.Type
	test = ola_rpc.Type_REQUEST

	rpc_message := new(ola_rpc.RpcMessage)
	rpc_message.Type = &test
	//rpc_message.Id = proto.Uint32(1)
	rpc_message.Name = proto.String("GetUniverseInfo")
	rpc_message.Buffer = dataToSend


	dataToSend2, err := proto.Marshal(rpc_message)
	if err != nil {
		log.Print("hi1")
		log.Fatalln(err)
	}
	log.Printf("%v", rpc_message)
	log.Printf("%v",dataToSend2)

	sendDataToDest(conn, dataToSend2)


	rpc_message2 := new(ola_rpc.RpcMessage)
	rsp := readData(conn)
	if err := proto.Unmarshal(rsp, rpc_message2); err != nil {
		log.Fatalln("Failed to parse resp:", err)
	}
	innerBuffer :=rpc_message2.GetBuffer()
	log.Printf("%v", rpc_message2)

	innards := new(ola_proto.UniverseInfoReply)
	if err := proto.Unmarshal(innerBuffer, innards); err != nil {
		log.Fatalln("Failed to parse resp:", err)
	}
	log.Printf("%v", innards.Universe[0].Name)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)



	//var reply ola_proto.UniverseInfoReply

}

func readData(conn net.Conn) []byte {


	header := make([]byte, 4)
	conn.Read(header)

	headerValue := int(binary.LittleEndian.Uint32(header))
	size := headerValue  & SIZE_MASK
	log.Printf("expecing %d bytes",size)


	data := make([]byte, size)
	n, err:= conn.Read(data)
	if err != nil {
		log.Print("hi2")
		log.Fatalln(err)
	}
	log.Printf("received %d bytes",n)
	return data
	//log.Print(data)
}

//func sendDataToDest(data []byte, dst *string) {
func sendDataToDest(conn net.Conn, data []byte) {

	headerContent := (PROTOCOL_VERSION << 28) & VERSION_MASK
	headerContent |= len(data) & SIZE_MASK
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