package admin 

import (
	_"fmt"
)

type MainController struct {
	BaseController
}

func (c *MainController) Index(){
	c.TplName = "admin/index.html"
}