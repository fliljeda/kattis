package main

import(
    "fmt"
)

func main(){
    var a1, b1, a2, b2 int
    fmt.Scanf("%d%d%d%d", &a1, &b1, &a2, &b2)
    var gunnarsum, emmasum int
    gunnarsum = a1 + b1 + a2 + b2
    fmt.Scanf("%d%d%d%d", &a1, &b1, &a2, &b2)
    emmasum = a1 + b1 + a2 + b2
    if gunnarsum > emmasum {
        fmt.Println("Gunnar")
    }else if gunnarsum == emmasum{
        fmt.Println("Tie")
    }else{
        fmt.Println("Emma")
    }
}
