// 定义路由日志格式

/**
如果你想要记录指定格式（如 JSON、键值）的信息，可以通过 gin.DebugPrintRouteFunc 来定义这个格式，
在下面这个例子中，我们将通过标准日志包记录所有路由信息，你也可以根据需要自定义日志格式：
*/
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 默认路由输出格式  DebugPrintRouteFunc调试日志的输出格式。
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.POST("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "foo")
	})

	r.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	// Listen and Server in http://0.0.0.0:8080
	r.Run()
}
