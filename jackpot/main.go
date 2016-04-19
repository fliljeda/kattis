package main
import(
    "fmt"
)

func main(){
    var noslots int
    fmt.Scanf("%d", &noslots)
    for i := 0; i < noslots; i++ {
        var noi int
        fmt.Scanf("%d", &noi)
        previous := 1
        overbil := false
        for j := 0; j < noi; j++ {
            var per int
            fmt.Scanf("%d", &per)
            previous = per * previous / commonDivider(previous, per)
        }
        if previous > 1000000000 {
            overbil = true
        }
        if overbil {
            fmt.Println("More than a billion.")
        }else{
            fmt.Println(previous)
        }
    }
}

func commonDivider(a,b int) int{
    if b == 0 {
        return a
    }else{
        return commonDivider(b, a%b)
    }
}
