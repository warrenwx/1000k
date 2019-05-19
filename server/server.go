package main

import (
	"net"
	"util"
	"fmt"
	"github.com/golang/protobuf/proto"
	"message"
)



func main(){
	netListen, err := net.Listen("tcp", "0.0.0.0:9090")
	util.CheckErrExit(err)
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		go processConnection(&conn)
	}
}

const HEADER_LENGTH = 2

func processConnection(conn *net.Conn) {
	//defer connection.Close()

	bufferHeader := make([]byte, HEADER_LENGTH)
	messageHeader := message.MessageHeader{}

	readHeader(conn, bufferHeader, HEADER_LENGTH)
	fmt.Println("recieve binary bytes head:", bufferHeader, " , len:", len(bufferHeader))
	proto.Unmarshal(bufferHeader, &messageHeader)
	switch messageHeader.MType {
	case 0:
		fmt.Println("MType 0")
	case 1:
		fmt.Println("MType 1")
	default:
		fmt.Println("MType none")
	}

	messageLength := messageHeader.MLength
	fmt.Println("MLength :", messageHeader.MLength)

	buffer := make([]byte, messageLength)
	readConnBytes(conn, buffer, messageLength)

	simpleMessageC := message.SimpleMessageC{}
	fmt.Println("recieve binary bytes msg:", buffer, " , len:", len(buffer))
	proto.Unmarshal(buffer, &simpleMessageC)
	fmt.Printf("read message from client %s %v\n", (*conn).RemoteAddr().String(), simpleMessageC)
}

func readHeader(conn *net.Conn, buffer []byte, readLength uint64) error {
	connection := (*conn)

	curPos :=  0
	for {
		count, err := connection.Read(buffer[curPos:])
		if err != nil {
			return err
		}

		curPos += count
		if uint64(curPos) >= readLength {
			return nil
		}
	}

}

// 从conn中读取长度为 readLength 的字节到 buffer 中
func readConnBytes(conn *net.Conn, buffer []byte, readLength uint64) error {
	connection := (*conn)

	curPos :=  0
	for {
		count, err := connection.Read(buffer[curPos:])
		if err != nil {
			return err
		}

		curPos += count
		if uint64(curPos) >= readLength {
			return nil
		}
	}

}