package main

import (
    "fmt"
    "github.com/rytst/gotorch"
)


func main() {
    var X = gotorch.Variable[[]int]{[]int{1,2,3}}
    fmt.Println(X.Data)
}
