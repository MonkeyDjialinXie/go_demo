package main
import "fmt"

func main(){
  //创建map使用内建的 make:make(map[key-type]val-type).
  m := make(map[string]int)

//make[key] = val 语法来设置键值对
  m["k1"] = 7
  m["k2"] = 13
  fmt.Println("map:",m)

//name[key] 来获取一个键的值
  v1 := m["k1"]
  fmt.Println("v1: ", v1)

// 调用内建的 len 时，返回的是键值对数目
  fmt.Println("len:", len(m))

//delete 可以从一个 map 中移除键值对
  delete(m, "k2")
  fmt.Println("map:", m)

//当从一个 map 中取值时，可选的第二返回值指示这个键是在这个 map 中。这可以用来消除键不存在和键有零值，像 0 或者 "" 而产生的歧义
  _, prs1 := m["k2"]
  fmt.Println("prs1:", prs1)

  prs2 := m["k2"]
  fmt.Println("prs2:", prs2)

//map[k:v k:v]的格式输出的
  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("map:", n)


}
