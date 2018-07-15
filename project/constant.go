package main

import "fmt"
import "math"

const s string = "consttant"

func main(){
  fmt.Println(s)

  const n = 5000000

  const d = 3e20 / n

  fmt.Println(int64(d))

  fmt.Println(math.Sin(n))
}
