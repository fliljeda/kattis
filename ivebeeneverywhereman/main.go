package main

import (
    "fmt"
)

func main(){
    var trips int
    fmt.Scanf("%d", &trips)
    for i := 0; i < trips;i++ {
        var flights int
        sum := 0
        fmt.Scanf("%d", &flights)
        var cities []string
        for y := 0; y < flights; y++{
            var city string
            fmt.Scanf("%v", &city)
            if(!contains(cities,city)){
                cities = append(cities, city)
                sum++
            }
        }
        fmt.Println(sum)
    }
}

func contains(cities []string, city string) bool {
    for i := range cities{
        if(cities[i] == city){
            return true
        }
    }
    return false
}
