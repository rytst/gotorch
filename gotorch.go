package gotorch

import (
//    "reflect"
    "math"
//    "errors"
)



/*
* For numerical type
*/
// type Number interface {
//     float64 | []float64
// }


/*
* Variable data structure
*/
type Variable[T any] struct {
    Data T
}


/*
* Forward method
*/
type I interface {
    Forward()
}


type Square struct {}

func (sqrt Square) Forward(X Variable[float64]) (Variable[float64], error) {
    var y Variable[float64]
    y.Data = math.Sqrt(X.Data)
    return y, nil
}

