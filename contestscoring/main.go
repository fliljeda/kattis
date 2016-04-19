package main

import(
    "fmt"
)

func main(){
    var time int
    var problem, solved string
    fmt.Scanf("%d%v%v", &time, &problem, &solved)
    problems := make(map[string]int)

    problemssolved := 0
    sum := 0
    for time != -1 {
        _, exists := problems[problem]
        if !exists {
            if solved == "right" {
                sum += time
                problemssolved++
            }else{
                problems[problem] = 20
            }
        }else{
            if solved == "right" {
                sum += problems[problem] + time
                problemssolved++
            }else{
                problems[problem] += 20
            }
        }

        fmt.Scanf("%d%v%v", &time, &problem, &solved)
    }
    fmt.Println(problemssolved, sum)
}
