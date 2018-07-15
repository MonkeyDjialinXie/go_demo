package main
import "fmt"

//(int, int) 在这个函数中标志着这个函数返回 2 个 int。
func value() (int,int){
  return 3,7
}

func main(){
  a,b := value()
  fmt.Println(a);
  fmt.Println(b);

//仅仅想返回值的一部分的话，你可以使用空白定义符 _。
  _, c := value();
  fmt.Println(c);
}
