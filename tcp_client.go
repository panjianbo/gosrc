package main

import (
  "fmt"
  "net"
  "os"
  "strconv"
  "time"
)

func main(){
        servAddr := "localhost:8080"
        tcpAddr, err:=net.ResolveTCPAddr("tcp", servAddr)
        if err != nil{
                fmt.Println("ResolveTCPAddr failed:", err.Error())
                os.Exit(1)
        }

        conn, err := net.DialTCP("tcp", nil, tcpAddr)
        if err != nil{
                fmt.Println("Dail failed:", err.Error())
                os.Exit(1)
        }
        defer conn.Close()

        go func (conn net.Conn){
                for{
                        data := make([]byte, 1024)
                        fmt.Println("begin read")
                        count, err := conn.Read(data)
                        if err != nil{
                                fmt.Println("Read from server failed:", err.Error())
                                os.Exit(1)
                        }
                        fmt.Println("get data count", count)
                }
        }(conn)

        strData := "test"
        for i:=0;i<1000;i++{
                count, err := conn.Write([]byte(strData+strconv.Itoa(i)))
                if err != nil{
                    fmt.Println("write to server failed:", err.Error())
                    os.Exit(1)
                }
                time.Sleep(time.Second)
                fmt.Println("send...", i, count)
        }
        time.Sleep(time.Second)
}
