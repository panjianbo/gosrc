package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func ServiceConn(connSocket net.Conn) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd in ServiceConn", r)
		}
	}()

	//业务逻辑处理
	//处理读数据
	go func(conn net.Conn) {
		for {
			data := make([]byte, 1024)
			count, err := conn.Read(data)
			if err != nil {
                                fmt.Println("Read error")
				return
			}
			fmt.Println("read data:", count)
			count, err = conn.Write(data[:count])
			if err != nil {
                                fmt.Println("Write error")
				return
			}
			fmt.Println("write data:", count)
		}
	}(connSocket)
}

func ServerListenSocket(listenSocket net.Listener) (err error) {
	for {
		connSocket, err := listenSocket.Accept()
		if err != nil {
                        fmt.Println("Accept Error: %v", err)
			return err
		}
		defer connSocket.Close()

		//对connSocket进行管理
		go ServiceConn(connSocket)
	}
}

func serveTcp(port int) (err error) {
	listenSocket, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("Listen error %v", err)
		return
	}
	defer listenSocket.Close()

	go ServerListenSocket(listenSocket)
	return err
}

func main() {
	err:=serveTcp(8080)
        fmt.Println("get err: %v", err)
	time.Sleep(time.Hour)
}

