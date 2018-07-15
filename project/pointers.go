package main
import "fmt"

func main(){
  i := 1
  fmt.Println("initial:",i);

  //传入值类型
  zeroval(i)
  fmt.Println("zeroval:", i)

//会报错cannot use &i (type *int) as type int in argument to zeroval
// zeroval(&i)
// fmt.Println("zeroval:", i)

  //传入指针类型
  //通过 &i 语法来取得 i 的内存地址
  zeroptr(&i)
  fmt.Println("zeroptr:", i)

  fmt.Println("pointer:", &i)
}


func zeroval(ival int){
   ival = 0
}

// *int 参数，意味着它用了一个 int指针
func zeroptr(iptr *int) {
  //*iptr 接着解引用 这个指针，从它内存地址得到这个地址对应的当前值。对一个解引用的指针赋值将会改变这个指针引用的真实地址的值
    *iptr = 0
}

//zeroval 在 main 函数中不能改变 i 的值，但是zeroptr 可以，因为它有一个这个变量的内存地址的引用
