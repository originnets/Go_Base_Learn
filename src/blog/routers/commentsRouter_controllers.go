package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["blog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["blog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "About",
			Router: `/about`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["blog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Comment",
			Router: `/comment`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["blog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Details",
			Router: `/details`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["blog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Error",
			Router: `/error`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["blog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Message",
			Router: `/message`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
