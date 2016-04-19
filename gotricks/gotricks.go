package main

import(
    "fmt"
)

func main(){
    //declaring array
    array := [3]int{1,2,3}

    //slice
    slice := []int{1,2,3}
    slice2 := make([]string, 5, 5) //([]type, length, cap)
    slice2[0] = "lol"
    copy(slice2[1:], []string{"hej","jag","heter","fredrik"})

    //insert (maybe lose last element)
    i := 1
    copy(array[i+1:], array[i:]) //if capacity is reached last element will drop out
    array[i] = 5

    //insert with append
    last := slice[len(slice)-1]
    copy(slice[i+1:], slice[i:])
    slice = append(slice, last)
    slice[i] = 4

    //length of slice
    length := len(slice)

    //Print
    fmt.Println(array)
    fmt.Print(slice, "\n") //comma makes a space between
    fmt.Printf("%v %v\n", slice2[0], slice2[3])
    fmt.Print("Length of slice: ", length, "\n")

    //IO verbs
    fmt.Println("\nIO-VERBS")
    test := 15
    fmt.Printf("%b\n", test) // base 2  (binary)
    fmt.Printf("%d\n", test) // base 10 (decimal)
    fmt.Printf("%x\n", test) // base 16 (hexa)
    test2 := "test"
    fmt.Printf("%v\n", test2) // value
    fmt.Printf("%v\n", test)

    fmt.Printf("%s\n", test2) // String
    fmt.Printf("%t\n", true) // bool

    //String handling
    fmt.Println("\nSTRING-HANDLING")
    teststring := "test this string"
    fmt.Println(teststring[0])
    fmt.Printf("%c\n", teststring[0])

}
