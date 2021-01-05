// 绑定 URL 路由参数

package main

import "github.com/gin-gonic/gin"

type Student struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var student Student
		// 将路由参数绑定到结构体中
		if err := c.ShouldBindUri(&student); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": student.Name, "uuid": student.ID})
	})
	route.Run(":8088")
}
