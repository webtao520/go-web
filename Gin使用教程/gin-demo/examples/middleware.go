// 自定义中间件

package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// HandlerFunc定义了gin中间件使用的处理程序作为返回值。
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 设置 example 变量
		c.Set("example", "12345")

		// 请求之前...
		// Next应该只在中间件内部使用。它执行调用处理程序内部链中的挂起处理程序。参见GitHub中的示例。
		c.Next()

		// 请求之后...
		// func Since(t Time) Duration
		// Since返回从t到现在经过的时间，等价于time.Now().Sub(t)。

		latency := time.Since(t)
		// 打印请求处理时间
		log.Print(latency)
		// 访问即将发送的响应状态码  返回当前请求的HTTP响应状态码。
		// 返回当前请求的HTTP响应状态码
		status := c.Writer.Status()
		log.Println(status)

	}
}

func main() {
	// New返回一个新的空白引擎实例，没有附加任何中间件
	r := gin.New()
	// 使用自定义的 Logger 中间件
	r.Use(Logger())
	// 定义路由
	r.GET("/test", func(c *gin.Context) {
		// func (c *Context) MustGet(key string) interface{}
		// MustGet返回给定键的值(如果存在的话)，否则会产生恐慌。
		example := c.MustGet("example").(string) //  interface{} 强制转换为 string
		log.Println(example)
	})
	// 监听 0.0.0.0:8080
	r.Run(":8080")
}
