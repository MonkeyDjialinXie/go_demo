package main
import "fmt"
import "net/url"
import "strings"
func main() {

//它包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段。
  s := "postgres://user:pass@host.com:5432/path?k=v#f"

//解析这个 URL 并确保解析没有出错。
  u, err := url.Parse(s)
      if err != nil {
          panic(err)
  }

//直接访问 scheme。
  fmt.Println(u.Scheme)

//User 包含了所有的认证信息，这里调用 Username和 Password 来获取独立值。
  fmt.Println(u.User)
  fmt.Println(u.User.Username())
  p, _ := u.User.Password()
  fmt.Println(p)

//Host 同时包括主机名和端口信息，如过端口存在的话，使用 strings.Split() 从 Host 中手动提取端口。
    fmt.Println(u.Host)
    h := strings.Split(u.Host, ":")
    fmt.Println(h[0])
    fmt.Println(h[1])

//路径和查询片段信息
    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])

}
