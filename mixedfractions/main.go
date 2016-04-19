package main

import(
    "fmt"
    "io"
)

func main(){
    for true {
        var upper, lower int
        _, err := fmt.Scanf("%d", &upper)
        if err == io.EOF{
            break
        }
        _, err = fmt.Scanf("%d", &lower)
        if err == io.EOF {
            break;
        }
        if lower == 0 && upper == 0 {
            continue
        }
        first := upper/lower
        var second int
        second = upper%lower
        fmt.Println(first, second, "/", lower)
    }

}
