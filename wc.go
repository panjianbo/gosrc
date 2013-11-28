package main
import (
"os"
"bufio"
"fmt"
"io"
)

func main(){
  if len(os.Args) <= 1{
     fmt.Println("please input filename.")
     os.Exit(0)
  }
  f, err := os.Open(os.Args[1])
  if err != nil {
     fmt.Printf("%v\n", err)
     os.Exit(1)
  }
  defer f.Close()

  br := bufio.NewReader(f)
  counter := 0
  total := 0
  for {
    line, prefix, err := br.ReadLine()
    if err == io.EOF{
      break
    }else if err == nil{
      counter ++
      if prefix{
         fmt.Println("big line")
      }
      total = total + len(line)
    }else{
      fmt.Printf("%v\n", err)
      break
    }
  }
  fmt.Println(os.Args[1], counter, total)
}
