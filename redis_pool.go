package main

import "menteslibres.net/gosexy/redis"
import "fmt"

const MAX_POOL_SIZE = 10
var RedisClientPool chan *redis.Client

func GetRedisClient() *redis.Client {

    if RedisClientPool == nil {
        RedisClientPool = make(chan *redis.Client, MAX_POOL_SIZE)
    }
    if len(RedisClientPool) == 0{
        go func(){
           for i := 0; i < MAX_POOL_SIZE; i++ {  
                client := redis.New()  
                err := client.Connect("localhost", 6379)  
                if err != nil {  
                    panic(err)  
                }     
                PutRedisClient(client)  
            }     
 
        }()
    }
    return <-RedisClientPool
}

func PutRedisClient(client *redis.Client){
    if RedisClientPool == nil {
        RedisClientPool = make(chan *redis.Client, MAX_POOL_SIZE)
    }
    if len(RedisClientPool) >= MAX_POOL_SIZE {  
        client.Quit()  
        return  
    }  
    RedisClientPool <- client
}

func ClearRedisClient(){
    if RedisClientPool != nil {
        count := len(RedisClientPool)
        for i:=0;i<count;i++{
           client := <-RedisClientPool
           client.Quit()
        }
    }
}

func main(){
   finish := make(chan int, 1000)
   client2 := GetRedisClient()
   client2.Set("test", 0)
   PutRedisClient(client2)
   for i:=0;i<1000;i++{
      go func(i int, c chan int){
         client := GetRedisClient()
         s1, _ := client.Ping()
         client.Incr("test")
         s2, _ := client.Get("test")
         fmt.Println(i,s1, s2)
         PutRedisClient(client)
         c<-1
      }(i, finish)
   }
   for i:=0;i<1000;i++{
      <-finish
   }
   fmt.Println("redis pool len:", len(RedisClientPool))
   ClearRedisClient()
}
