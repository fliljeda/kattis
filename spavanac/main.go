package main

import(
    "fmt"
)

func main(){
    var hours, minutes int
    fmt.Scanf("%d%d", &hours, &minutes)
    minutes = minutes - 45
    if minutes < 0 {
        minutes = 60 + minutes
        hours--
        if hours < 0 {
            hours = 23
        }
    }
    fmt.Printf("%d %d\n", hours, minutes)
}
