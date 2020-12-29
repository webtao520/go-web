// 绑定 HTML 复选框 checkbox
package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

type myForm struct {
  Colors []string `form:"colors[]"`
}

func main()  {
  r := gin.Default()
  // LoadHTMLGlob加载由glob模式标识的HTML文件，并将结果与​​HTML渲染器关联。
  r.LoadHTMLGlob("templates/**/*")
  r.GET("/colors", func(c *gin.Context) {
	  // type H map[string]interface{}
    c.HTML(http.StatusOK, "checkbox/color.tmpl", gin.H{
      "title": "Select Colors",
    })
  })
  r.POST("/colors", func(c *gin.Context) {
    var fakeForm myForm
    // ShouldBind 和 Bind 类似，不过会在出错时退出而不是返回400状态码
    c.ShouldBind(&fakeForm)
    c.JSON(200, gin.H{"color": fakeForm.Colors})
  })
  r.Run()
}