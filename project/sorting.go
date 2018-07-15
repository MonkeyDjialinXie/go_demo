package main
import "fmt"
import "sort"

//sort 包实现了内置和用户自定义数据类型的排序功能
func main() {

  //注意排序是原地更新的，所以他会改变给定的序列并且不返回一个新值。
  strs := []string{"c","a","b"}
  sort.Strings(strs)
  fmt.Println("Strings:" ,strs)

  ints := []int{7,5,9}
  sort.Ints(ints)
  fmt.Println("Ints:", ints)

//可以使用 sort 来检查一个序列是不是已经是排好序的。
  s := sort.IntsAreSorted(ints)
  fmt.Println("Sorted:",s)


}
