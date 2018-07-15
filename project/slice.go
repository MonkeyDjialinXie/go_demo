package main

import "fmt"

func main(){

  //不像数组，slice 的类型仅由它所包含的元素决定（不像数组中还需要元素的个数)
  //要创建一个长度非零的空slice，需要使用内建的方法 make。这里我们创建了一个长度为3的 string 类型 slice（初始化为零值）
  s := make([]string, 3)

  fmt.Println("emp",s)

  s[0] = "a"
  s[1] = "b"
  s[2] = "c"

  fmt.Println("set:", s)
  fmt.Println("get:", s[2])

  fmt.Println("len:", len(s))

//内建的 append，它返回一个包含了一个或者多个新值的 slice
  s = append(s, "d")
  s = append(s, "e", "f")
  fmt.Println("apd:", s)

  c := make([]string, len(s))
  copy(c, s)
  fmt.Println("cpy:", c)

  // slice[low:high] 语法进行“切片”操作
  //= 是赋值， := 是声明变量并赋值。
  l := s[2:5]
  fmt.Println("sl1:", l)

  l = s[:5]
  fmt.Println("sl2:", l)

  l = s[2:]
  fmt.Println("sl3:", l)

  t := []string{"g", "h", "i"}
  fmt.Println("dcl:", t)

//组成多维数据结构。内部的 slice 长度可以不同，这和多位数组不同
  twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
  fmt.Println("2d: ", twoD)


}
