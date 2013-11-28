package main

import (
  "fmt"
  "github.com/PuerkitoBio/purell"
)

func ExampleNormalizeURLString() {
  if normalized, err := purell.NormalizeURLString("hTTp://someWEBsite.com:80/Amazing%41%3f/url/",
    purell.FlagsAllGreedy); err != nil {
    panic(err)
  } else {
    fmt.Print(normalized)
  }
  //Output: http://somewebsite.com:80/Amazing%3F/url/
}


func main(){
   ExampleNormalizeURLString() 
}
