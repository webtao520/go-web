// 通过自定义结构体绑定表单请求数据

package main
    
import "github.com/gin-gonic/gin"
    
// 最基本的结构体
type StructA struct {
    FieldA string `form:"field_a"`
}
    
// 嵌套结构体的结构体
type StructB struct {
    NestedStruct StructA
    FieldB string `form:"field_b"`
}
    
// 嵌套结构体指针的结构体
type StructC struct {
    NestedStructPointer *StructA
    FieldC string `form:"field_c"`
}
    
// 嵌套匿名结构体的结构体
type StructD struct {
    NestedAnonyStruct struct {
        FieldX string `form:"field_x"`
    }
    FieldD string `form:"field_d"`
}
    
// 返回 StructB
func GetDataB(c *gin.Context) {
    var b StructB
    // 读取请求数据并写入结构体b
    c.Bind(&b)
    // 返回 JSON 格式响应
    c.JSON(200, gin.H{
        "a": b.NestedStruct,
        "b": b.FieldB,
    })
}
    
// 返回 StructC
func GetDataC(c *gin.Context) {
    var b StructC
    c.Bind(&b)
    c.JSON(200, gin.H{
        "a": b.NestedStructPointer,
        "c": b.FieldC,
    })
}
    
// 返回 StructD
func GetDataD(c *gin.Context) {
    var b StructD
    c.Bind(&b)
    c.JSON(200, gin.H{
        "x": b.NestedAnonyStruct,
        "d": b.FieldD,
    })
}
    
func main()  {
    r := gin.Default()
    r.GET("/getb", GetDataB)
    r.GET("/getc", GetDataC)
    r.GET("/getd", GetDataD)
    
    r.Run()
}