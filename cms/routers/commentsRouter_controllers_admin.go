package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)


func init (){

    beego.GlobalControllerRouter["cms/controllers/admin:CronController"] = append(beego.GlobalControllerRouter["cms/controllers/admin:CronController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "cron",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})
}