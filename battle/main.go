package main

import(
    "fmt"
)

func main(){
    var number, count int
    fmt.Scanf("%d", &number)
    for number != 0{
        count++
        fmt.Println("SET", count)
        var names = make([]string, number)
        var name string
        for i := 0; i<number/2; i++ {
            fmt.Scanf("%v", &name)
            names[i] = name
            fmt.Scanf("%v", &name)
            names[len(names)-i-1] = name
        }
        if len(names)%2 == 1{
            fmt.Scanf("%v", &name)
            names[len(names)/2] = name
        }
        for j:=0;j<number;j++{
            fmt.Println(names[j])
        }
        fmt.Scanf("%d", &number)
    }
}
