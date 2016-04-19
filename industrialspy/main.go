package main

import(
    "fmt"
    "strconv"
    "math"
)
var numberofprimes int
var numberschecked []int

func main(){
    var noi int //number of inputs
    fmt.Scanf("%d", &noi)
    for i:=0;i<noi;i++ {
        var numberstring string
        fmt.Scanf("%v", &numberstring)
        allNumbers(numberstring)
        fmt.Println(numberofprimes)
        numberofprimes = 0
    }
}

//tries all orders of numbers with the given string
func allNumbers(numberstring string){
        numberslice := []byte(numberstring)
        length := len(numberslice)
        for i := 1; i <= length; i++ { //size of number
            checked := make([]byte, length, length) //Which first digit has alreaddy been cehcked
            numberschecked = make([]int, 0, 0) //Which numbers has been checked
            for y := 0; y < length; y++{ //start of number
                check := false
                for j := range checked {
                    if checked[j] == numberslice[y] {
                        check = true
                    }
                }
                if check != true && numberslice[y] != 48{ //If char has been checked or isnt zero
                  checked[y] = numberslice[y]
                  builder := make([]byte, i, i)
                  exhaustOrders(y, builder, numberslice, 0)
                }
            }
        }
}

//Exhausts all orders with the fixed amount of characters
func exhaustOrders(currentchar int, builder []byte, prevarsenal []byte, index int){
    if index == len(builder) {
        //fmt.Println(string(builder))
        s := string(builder)
        n, _ := strconv.Atoi(s)
        if notInList(n, numberschecked) && isPrime(n) { //if n hasnt been checked and is a prime
            //fmt.Println(string(builder))
            numberschecked = append(numberschecked, n)
            numberofprimes++
        }
    }else{
        arsenal := make([]byte, len(prevarsenal))
        copy(arsenal, prevarsenal)
        builder[index] = arsenal[currentchar] //
        arsenal = append(arsenal[:currentchar], arsenal[currentchar+1:]...) //delete currentchar
        for i := 0; i < len(arsenal); i++ { //next number in line
            exhaustOrders(i, builder, arsenal, index + 1)
        }
        if len(arsenal) == 0 {
            exhaustOrders(0, builder, arsenal, index + 1)
        }
    }
}

func isPrime(num int) bool {
    if(num < 2) {return false}
    top := int(math.Sqrt(float64(num)))
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

func notInList(num int, numberschecked []int) bool {
    for i := range numberschecked {
        if num == numberschecked[i] {
            return false
        }
    }
    return true
}
