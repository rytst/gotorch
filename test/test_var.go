package main

import (
    "fmt"
    . "github.com/rytst/gotorch"
)


func main() {
    var X = Variable[float64]{256}
    var sqrt = Square{}
    var y, err = sqrt.Forward(X)

    fmt.Println(y)
    fmt.Println(err)
}
