package util

import (
	"fmt"
	"os"
	"io/ioutil"
	"net"
)


func CheckErrExit(err error) {
	if err != nil {
		fmt.Println("Fatal Error:", err.Error())
		os.Exit(-1)
	}
}

func WriteFile(data []byte) error {
	logFile := "serial"
	err := ioutil.WriteFile(logFile, data, 0644)
	return err
}

func ReadFile() ([]byte, error){
	logFile := "serial"
	data, err := ioutil.ReadFile(logFile)

	return data, err
}

// 从conn中读取长度为 readLength 的字节到 buffer 中
func ReadConnBytes(conn *net.Conn, buffer []byte, readLength int) error {
	connection := (*conn)

	curPos :=  0
	for {
		count, err := connection.Read(buffer[curPos:])
		if err != nil {
			return err
		}

		curPos += count
		if curPos >= readLength {
			return nil
		}
	}

}

// 向 conn 中写入 buffer 中长度为 contentLength 的字节
func WriteConnBytes(conn *net.Conn, buffer []byte) error {
	connection := *conn

	_, err := connection.Write(buffer)
	
	return err
}