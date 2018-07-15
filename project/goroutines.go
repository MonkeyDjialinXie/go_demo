package main
import "fmt"

func f(from string) {
    for i := 0; i < 100; i++ {
        fmt.Println(from, ":", i)
    }
}

func main(){
  //使用一般的方式调并同时运行。
  f("direct")

//在一个 Go 协程中调用这个函数。这个新的 Go 协程将会并行的执行这个函数调用。
  go f("goroutine")

//匿名函数启动一个 Go 协程
  go func(msg string) {
    for i := 0; i < 100; i++ {
        fmt.Println(msg,":",i)
    }
  }("going")

    var input string
    //Scanln 代码需要我们在程序退出前按下任意键结束。
    fmt.Scanln(&input)
    fmt.Println("done")

//  将首先看到阻塞式调用的输出，然后是两个 Go 协程的交替输出。这种交替的情况表示 Go 运行时是以异步的方式运行协程的
}
