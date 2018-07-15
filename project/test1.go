package main

import (
  "fmt"
)

func main(){
  s:= []int{1,2,3}
  ss:= s[1:]  //s[1:]除了1之后后面输出
  fmt.Println(ss)
  ss = append(ss,4)  //对ss拓展一个4
  fmt.Println(ss)

  for _,v:= range ss{
    v += 10
  }

  for i:= range ss{
    ss[i] += 10
  }

  fmt.Println(s)

}

输出：
[2 3]
[2 3 4]
[1 2 3]
