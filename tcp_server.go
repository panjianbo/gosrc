package main

import (
"fmt"
"net"
"os"
)

func main(){
        ln, err := net.Listen("tcp", ":8080")
        if err != nil{
            fmt.Printf("Fatal error:%s", err.Error())
            return
        }
        defer ln.Close()
        for{
                conn, err := ln.Accept()
                if err != nil{
                        fmt.Printf("Fatal error:%s", err.Error())
                        continue
                }
                go func (conn net.Conn) {
                    for{
                        data := make([]byte, 1024)
                        count, err := conn.Read(data)
                        if err != nil{
                                fmt.Printf("Fatal error:%s", err.Error())
                                os.Exit(1)
                        }
                        fmt.Println("read data:", count)
                        count, err = conn.Write(data[:count])
                        if err != nil{
                                fmt.Printf("Fatal error:%s", err.Error())
                                os.Exit(1)
                        }
                        fmt.Println("write data:", count)
                    }
                }(conn)
                defer conn.Close()
        }
}
