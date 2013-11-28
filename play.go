package main

import "fmt"

type A func(int, int)

func (f A) Serve(){
    fmt.Println("serve 2")
}

func serve(int, int){
    fmt.Println("serve 1")
}

func main(){
   a := A(serve)
   a(1,2)
   a.Serve()
}
