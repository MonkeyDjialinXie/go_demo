
//定时器 是当你想要在未来某一刻执行一次时使用的
//打点器 则是当你想要在固定的时间间隔重复执行准备的。
//这里是一个打点器的例子，它将定时的执行，直到我们将它停止。

package main
import "time"
import "fmt"
func main() {
  ticker := time.NewTicker(time.Millisecond * 500)
  go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
  }()

  
  time.Sleep(time.Millisecond * 1600)
      ticker.Stop()
      fmt.Println("Ticker stopped")
}
