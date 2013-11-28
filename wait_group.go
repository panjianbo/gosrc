package main

import (
  "fmt"
  "sync"
  "time"
)

type Service struct{
    ch chan bool
    waitGroup *sync.WaitGroup
}

func NewService *Service{
    s:= &Service{
    ch:make(chan bool),
    waitGroup:&sync.WaitGroup{},
    }
    s.waitGroup.Add(1)
    return s
}

func (s *Service) ServeListen(listener, *net.TcpListener){
}

func main(){
    var wg sync.WaitGroup
    for i:=0;i<100;i++{
       wg.Add(1)
       go func(i int){
          defer wg.Done()
          fmt.Println(i)
          time.Sleep(time.Second*2)
          fmt.Println(i+1000)
       }(i)
    }
    wg.Wait()
    fmt.Println("all done")
}
