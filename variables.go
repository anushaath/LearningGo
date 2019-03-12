package main

import (
    "fmt"
    "math/cmplx"
)
   
var (
    kin bool = false
    iot complex128 = cmplx.Sqrt(-5 +56i)
    x   uint64 = 1<<64 -1
)
    
func main(){

  fmt.Println("Type : %T, Value = %v\n", kin, kin)
  fmt.Println("Type : %T, Value = %v\n", iot, iot)
  fmt.Println("Type : %T, Value = %v\n", x, x)
  
}
