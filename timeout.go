package main

import(
"fmt"
"time"
)

func timeout_test(){
    done:=make(chan bool, 1)
    timeout:=time.After(time.Second*5)
    go func(){
      fmt.Println("begin_sleep..")
      time.Sleep(time.Second*10)
      fmt.Println("end_sleep..")
      done <- true
    }()

    select{
    case <- done:
      fmt.Println("finished")
      return
    case <- timeout:
      fmt.Println("timeout")
      return
    }
}

func main(){
   fmt.Println("go")
   go timeout_test()
   time.Sleep(time.Second*20)
   fmt.Println("go end")
}
