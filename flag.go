package main

import(
  "flag"
  "fmt"
)

func main(){
  flag.Parse()
  for i:=0;i<flag.NArg();i++{
    fmt.Println(flag.Arg(i))
  }
}
