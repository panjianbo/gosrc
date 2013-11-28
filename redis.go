package main

import "menteslibres.net/gosexy/redis"
import "fmt"

func main(){
    var client *redis.Client

    client = redis.New()
    client.Connect("localhost", 6379)

    s,_ := client.Ping()
    fmt.Println(s)
    client.Incr("hello")
    s, _ = client.Get("hello")
    fmt.Println(s)
    client.Quit()
}
