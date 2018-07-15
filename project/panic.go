package main
import "os"

//


//panic 意味着有些出乎意料的错误发生。通常我们用它来表示程序正常运行中不应该出现的，或者我们没有处理好的错误。
func main() {
  panic("a problem")

  _, err := os.Create("/tmp/file")
  if err != nil{
    panic(err)
  }

}


/**panic: a problem

goroutine 1 [running]:
main.main()
	/Users/jialinxie/Documents/整理/go/project/panic.go:4 +0x39
exit status 2*/

//运行程序将会引起 panic，输出一个错误消息和 Go 运行时栈信息，并且返回一个非零的状态码。
