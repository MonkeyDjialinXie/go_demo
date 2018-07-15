package main
import "fmt"
import "math"

//几何体的基本接口
type geometry interface {
  area() float64
  perim() float64
}

type rect struct {
  width, height float64
}

type circle struct {
  radius float64
}

//在 Go 中实现一个接口，我们只需要实现接口中的所有方法
func (r rect) area() float64{
  return r.width * r.height
}

func (r rect) perim() float64{
  return 2*r.width + 2*r.height
}

func (c circle) area() float64{
  return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64{
  return 2 * math.Pi * c.radius
}

//一个一通用的 measure 变量为一个接口
func measure(g geometry){
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main(){
  //结构体类型 circle 和 rect 都实现了 geometry接口，所以可以使用它们的实例作为 measure 的参数。
  r := rect{width: 3, height: 4}
  c := circle{radius: 5}

  measure(r)
  measure(c)
}
