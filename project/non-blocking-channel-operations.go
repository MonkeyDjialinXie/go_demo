package main
import "fmt"

//使用带一个 default 子句的 select 来实现非阻塞 的发送、接收，甚至是非阻塞的多路 select
func main(){
  messages := make(chan string)
  signals := make(chan bool)


//这里是一个非阻塞接收的例子。如果在 messages 中存在，然后 select 将这个值带入 <-messages case中。如果不是，就直接到 default 分支中。
  select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
  }

//非阻塞发送
  msg := "hi"
  select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
  }

//在 messages和 signals 上同时使用非阻塞的接受操作。
  select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}
