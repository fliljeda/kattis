package main

import(
    "fmt"
)

func main(){
    input := 0
    fmt.Scanf("%d", &input)
    for input != -1 {
        i, previous,sum := 0, 0, 0
        for i < input {
            var mph, time int
            fmt.Scanf("%d%d", &mph, &time)
            sum += mph*(time-previous)
            previous = time
            i++
        }
        fmt.Println(sum, "miles")
        fmt.Scanf("%d", &input)
    }
}
