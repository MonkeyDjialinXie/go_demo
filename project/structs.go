package main
import "fmt"

//person 结构体包含了 name 和 age 两个字段
type person struct{
  name string
  age int
}

func main(){
  fmt.Println(person{"Bobo",20})

  fmt.Println(person{name: "Alice", age: 30})

  fmt.Println(person{name: "Fred"})

//& 前缀生成一个结构体指针
  fmt.Println(&person{name:"Mike",age:20})

//使用点来访问结构体字段。
  s := person{name:"Seam",age:50}

  sp := &s
  fmt.Println(sp.age)

//对结构体指针使用. - 指针会被自动解引用
  sp.age = 51
  fmt.Println(sp.age)
//已经改变了原s的值，结构体是可变的
  fmt.Println(s)

  sb := s
  fmt.Println(sb.age)
  sb.age = 52
//不会改变原s的值
  fmt.Println(sb.age)
  fmt.Println(s)




}
