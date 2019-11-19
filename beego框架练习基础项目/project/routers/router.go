package routers

import (
	"project/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	//beego.Router("/post1", &controllers.MainController{})
	//beego.Router("/register", &controllers.MainController{})
    //注意：如果实现自定义的get请求方法，请求将不会访问默认方法
    beego.Router("/login", &controllers.MainController{},"get:Login;post:Loginpost")
    beego.Router("/index", &controllers.MainController{},"get:ShowIndexGet")
    beego.Router("/addArticle", &controllers.MainController{},"get:AddArticleGet;post:AddArticlePost")
    beego.Router("/moreInfo", &controllers.MainController{},"get:MoreInfoGet")
    beego.Router("/update", &controllers.MainController{},"get:UpdateGet;post:UpdatePost")
    beego.Router("/del", &controllers.MainController{},"get:DelGet")
}
