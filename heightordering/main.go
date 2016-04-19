package main

import(
    "fmt"
)

func main(){
    var times int
    fmt.Scanf("%d", &times)
    for i := 0; i < times; i++ {
        sum := 0
        var printnr int
        fmt.Scanf("%d", &printnr)
        queue := make([]int,1)
        fmt.Scanf("%d", &queue[0])
        for j := 0; j < 19; j++ {
            var current int
            fmt.Scanf("%d", &current)
            inserted := false
            for k := range queue{
                if current < queue[k] {
                    sum += len(queue) - k
                    queue = append(queue, 0)
                    copy(queue[k+1:], queue[k:])
                    queue[k] = current
                    inserted = true;
                    break
                }
            }
            if(!inserted){
                queue = append(queue, current)
            }
        }
        fmt.Println(printnr, sum)
    }
}
