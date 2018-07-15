package main

import "fmt"

  func main(){

  nums := []int{2, 3, 4}

  sum := 0

  //_, 表示要的是值而不是下标
  for _, num := range nums{
    sum += num
    fmt.Println("sum",sum)
  }

  sum1 := 0
  for  num := range nums{
    sum1 += num
    fmt.Println("sum1",sum1)
  }


  for i, num:= range nums{
    if num == 3 {
      fmt.Println("index",i)
    }
  }

  kvs := map[string]string{"a": "apple", "b": "banana"}

  for k, v:= range kvs {
     fmt.Printf("%s -> %s\n", k, v)
  }

  for i, c := range "go" {
        fmt.Println(i, c)
  }
}
