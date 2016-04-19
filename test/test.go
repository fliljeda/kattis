package main
import (
    "fmt"
    //"strconv"
)

func main(){
    lets := make([]byte, 5,5)
    lets[2] = 123
    gel := false
    for i := range lets {
        if lets[i] == 123 {

            gel = true
        }
    }
    fmt.Println(gel)
}
