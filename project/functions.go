package main

import "fmt"

//函数单返回值
func plus(a int,b int) int {
  return a + b
}

func main(){
  c := plus(1,2)
  fmt.Println(c)
}
