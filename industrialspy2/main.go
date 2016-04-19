package main

import(
    "fmt"
    "math"
    "strconv"
)
var numberofprimes int

func main(){
    var noi int
    fmt.Scanf("%d", &noi)
    for index := 0; index < noi; index++ {
        var numberstring string
        fmt.Scanf("%v", &numberstring)
        numberlength := len(numberstring)
        for length := 1; length <= numberlength; length++ {
            numberstring2 := make([]byte, numberlength, numberlength)
            copy(numberstring2, []byte(numberstring))
            allOrders(length, numberstring2, make([]byte, length, length), 0)
        }
        printprimes()
    }
}

func printprimes(){
    fmt.Println(numberofprimes)
    numberofprimes = 0
}

func allOrders(itg int, arsenal []byte, builder []byte, first int){ //iterations to go
    if itg == 0 {
        number,_ := strconv.Atoi(string(builder))
        if isPrime(number) {
            numberofprimes++
        }
    }else{
        checked := []byte{}
        for i := range arsenal {
            if !contains(checked, arsenal[i]) && !(first != 0 && arsenal[i] == '0'){
                checked = append(checked, arsenal[i])

                newarsenal := make([]byte, len(arsenal))
                copy(newarsenal, arsenal)

                newbuilder := make([]byte, len(builder))
                copy(newbuilder, builder)

                newbuilder[len(newbuilder)-itg] = arsenal[i]
                newarsenal = append(newarsenal[:i], newarsenal[i+1:]...)
                allOrders(itg-1, newarsenal, newbuilder, first +1)
            }
        }
    }
}

func isPrime(num int) bool {
    if(num < 2) {return false}
    top := int(math.Sqrt(float64(num)))
    if num == 2 {
        return true
    }
    if num%2==0 {
        return false
    }
    for i:=3; i<=top; i=i+2{
        if num%i == 0 {
            return false
        }
    }
    return true
}

func contains(list []byte, num byte) bool{
    for _,e := range list {
        if e == num {
            return true
        }
    }
    return false
}
