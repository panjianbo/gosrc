package main
import "fmt"
func main(){
    a := make([][]byte, 0)
    b := make([]byte, 0)
    a = append(a, b)
    fmt.Println(a)
}
