package main

import(
    "fmt"
    "strconv"
)

func main(){
    var number int64
    fmt.Scanf("%d", &number)
    numberstring := strconv.FormatInt(number, 2)
    newstring := []byte(numberstring)
    length := len(newstring)
    for i := 0; i < length/2; i++ {
        newstring[i], newstring[length-i-1] = newstring[length-i-1], newstring[i]
    }
    numberstring = string(newstring)
    lol, _ := strconv.ParseInt(numberstring, 2, 64)
    fmt.Println(lol)
}
