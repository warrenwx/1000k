package main

import (
	"fmt"
	"net"
	"message"
	"github.com/golang/protobuf/proto"
	"util"
)

func main(){
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	util.CheckErrExit(err)
	defer conn.Close()
	// go processConnection(&conn)

	simpleMessageC := message.SimpleMessageC {
		A: 123,
		B: "this is B of simple message",
	}

	msgLength := uint64(simpleMessageC.XXX_Size())

	bytes, err := buildMessageHead(0, msgLength)
	util.CheckErrExit(err)
	fmt.Println("send binary bytes head:", bytes, " , len:", len(bytes))
	err = util.WriteConnBytes(&conn, bytes)
	util.CheckErrExit(err)

	bytes, err = proto.Marshal(&simpleMessageC)
	fmt.Println("send binary bytes msg:", bytes, " , len:", len(bytes))
	util.CheckErrExit(err)
	err = util.WriteConnBytes(&conn, bytes)
	util.CheckErrExit(err)

	fmt.Println("send message to ", conn.RemoteAddr(), " sucess.")
}

// 构建头消息
func buildMessageHead(msgType uint64, msgLength uint64) ([]byte, error) {
	msgHead := message.MessageHeader {
		MType: msgType,
		MLength: msgLength,
	}

	return proto.Marshal(&msgHead)
}

func processConnection(conn *net.Conn) {
	connection := (*conn)
	defer connection.Close()

	//recoveredMessage := message.WrappMessage{}
	//msgLength := unsafe.Sizeof(recoveredMessage)
	//buffer := make([]byte, msgLength)
	//fmt.Println("size:", recoveredMessage.XXX_Size)
	//readConnBytes(conn, buffer, unsafe.Sizeof(recoveredMessage))
	//fmt.Println("read message from server " + connection.RemoteAddr().String(), " with 10 bytes ", buffer[0:10])
}
