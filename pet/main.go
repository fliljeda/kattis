package main

import(
    "fmt"
)

func main(){
    var leadersum, leadernr int
    for i := 1; i <= 5; i++ {
        var sum, current int
        fmt.Scanf("%d", &current)
        sum += current
        fmt.Scanf("%d", &current)
        sum += current
        fmt.Scanf("%d", &current)
        sum += current
        fmt.Scanf("%d", &current)
        sum += current

        if(sum > leadersum){
            leadersum = sum
            leadernr = i
        }

    }
    fmt.Println(leadernr, leadersum)
}
