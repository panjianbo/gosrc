package main

import (
	"fmt"
	"net"
)

func tcp_test(ch chan int) {
	address := "p.api.pc120.com:80"
	//address := "10.20.216.116:80"
        defer func() {
            if err := recover(); err != nil {
                fmt.Println(err)
            }
        }()

	con, err := net.Dial("tcp", address)
	defer con.Close()
        if err != nil{
           fmt.Println("Connect error")
           ch <- 1
           return
        }

	post_data := "POST /spp/ HTTP/1.0\r\nHost: p.api.pc120.com\r\nAccept: */*\r\nContent-Type: multipart/form-data; boundary=-VisualSeawind-\r\nAccept: */*\r\nAccept-Encoding: gzip, deflate\r\nUser-Agent: Microsoft-ATL-Native/8.00\r\nConnection: keep-alive\r\nAccept-Language: zh-CN\r\nContent-Length: 224\r\n\r\n"

	fmt.Fprintf(con, post_data)

	buf := make([]byte, 512)
	length, err := con.Read(buf)
	if err != nil {
                fmt.Println("Read Error")
		ch <- 1
		return
	}
	fmt.Println(string(buf[length-14 : length]))
        ch <- 1
}

func main() {
    ch := make(chan int)

    loops := 1
   
    for j:=0; j<loops; j++ {

	counter := 1

	for i := 0; i < counter; i++ {
		go tcp_test(ch)
	}

	for i := 0; i < counter; i++ {
		<-ch
	}
    }
    fmt.Println("all done")
}
