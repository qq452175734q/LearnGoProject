package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"path"
	"project/models"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	var FilterUser = func(ctx *context.Context) {
		_, ok := ctx.Input.Session("uid").(int)
		if !ok {
			ctx.Redirect(302, "/login")
		}
	}
	beego.InsertFilter("/user/:id([0-9]+)",beego.BeforeRouter,FilterUser)
/*新增
	//orm对象
	o := orm.NewOrm()
	//插入数据的结构体对象
	user := models.User{}
	//对结构体赋值
	user.Name ="zhangsan"
	user.Pwd = "123"
	//插入
	_,err:=o.Insert(&user)
	if err!=nil {
		beego.Info("插入失败",err)
		return
	}
*/
	/*查询
	//1有orm对象
	o:=orm.NewOrm()
	//2查询的对象
	user := models.User{}
	//3指定查询的字段值
	//user.Id = 1 根据ID查询
	//err := o.Read(&user)
	user.Name = "zhangsan"
	err := o.Read(&user,"Name")
	//4查询
	if err!=nil {
		beego.Info("查询失败",err)
		return
	}
	beego.Info("查询成功",user)
	*/

	/*更新
	o := orm.NewOrm()
	user := models.User{}
	user.Id = 1
	err := o.Read(&user)
	if err == nil {
		user.Pwd = "123456"
		_,err := o.Update(&user,"Pwd")
		if err != nil {
			beego.Info("更新失败",err)
		}
	}*/

	/*删除
	o := orm.NewOrm()
	user := models.User{}
	user.Id = 1
	_,err := o.Delete(&user)
	if err!=nil {
		beego.Info("删除失败",err)
		return
	}
	beego.Info("删除成功",&user)
	*/

	//c.Data["Website"] = "beego.me"
	//c.Data["Data"] = "beego请求"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.TplName = "register.html"
}

func (c *MainController) Post() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Data"] = "Post请求"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	//c.TplName = "postIndex.html"

	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	if userName == "" || pwd == "" {
		beego.Info("数据不能为空")
		c.Redirect("/register",302)
		return
	}

	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	user.Pwd = pwd
	_,err := o.Insert(&user)
	if err != nil {
		beego.Info("注册失败",err)
		c.Redirect("/register",302)
	}
	beego.Info("注册成功",&user)
	c.Data["userName"] = userName
	c.TplName = "login.html"
	//c.Redirect("/login",1) //速度快
}

func (c *MainController)  Login() {
	c.TplName = "login.html"
}

func (c *MainController)  Loginpost() {
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	if userName == "" || pwd == "" {
		beego.Info("数据不能为空")
		c.Redirect("/login",302)
		return
	}
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	user.Pwd = pwd
	err := o.Read(&user,"Name","Pwd")
	if err != nil {
		beego.Info("账号或者密码错误",err)
		c.Redirect("/login",303)
		return
	}
	c.Redirect("/index",302)
}

//显示页
func (c *MainController) ShowIndexGet()  {
	o := orm.NewOrm()
	var artic []models.Article
	_,err := o.QueryTable("Article").All(&artic)
	if err!=nil {
		beego.Info("查询文章出现错误")
		return
	}
	beego.Info(artic)
	c.Data["articles"] = artic
	c.TplName = "index.html"
}

//显示详情页
func (c *MainController) MoreInfoGet() {

	id,err := c.GetInt("id")
	if err!=nil {
		beego.Info("获取ID错误")
	}
	o := orm.NewOrm()
	art:= models.Article{Id:id}
	err = o.Read(&art)
	if err!=nil {
		beego.Info("查询错误")
	}
	c.Data["art"] = art
	c.TplName = "content.html"
}





func (c *MainController) AddArticleGet()  {
	c.TplName = "add.html"
}
//添加
func (c *MainController) AddArticlePost()  {
	articleName := c.GetString("articleName")
	selectType := c.GetString("select")
	content := c.GetString("content")
	if articleName == "" ||  selectType == "" || content == ""{
		beego.Info("数据不能为空")
		return
	}

	f,h,err := c.GetFile("uploadname")
	defer f.Close()

	//规定格式
	ext := path.Ext(h.Filename) //获取文件后缀
	if ext != ".jpg" && ext != ".png" {
		beego.Info("上传文件格式错误")
	}
	//限制图片大小
	if h.Size > 5000000 {
		beego.Info("上传文件过大")
	}
	//对文件重命名，防止文件名重复
	filename := time.Now().Format("20060102150405")+ ext

	if err != nil {
		beego.Info("文件上传失败",err)
	}else {
		c.SaveToFile("uploadname","./static/img/"+filename)
	}

	article := models.Article{}
	article.Title = articleName
	article.Type = selectType
	article.Acontent = content
	article.Aimg = "/static/img/"+filename
	o := orm.NewOrm()
	o.Insert(&article)

	c.Redirect("index.html",302)
}

//更新
func (c *MainController) UpdateGet()  {

	id,err := c.GetInt("id")
	if err!=nil {
		beego.Info("获取ID错误")
	}
	o := orm.NewOrm()
	art:= models.Article{Id:id}
	err = o.Read(&art)
	if err!=nil {
		beego.Info("查询错误")
	}
	c.Data["art"] = art
	c.TplName = "update.html"
}

//更新
func (c *MainController) UpdatePost()  {
	id,err := c.GetInt("id")
	if err != nil {
		beego.Info("未获取到Id")
		return
	}
	title := c.GetString("articleName")
	context := c.GetString("content")

	if title == "" || context == "" {
		beego.Info("缺少标题或内容")
		return
	}

	f,h,err := c.GetFile("uploadname")
	var filename string
	if err != nil {
		beego.Info("未获取到文件")
		return
	}else {
		defer f.Close()
		//规定格式
		ext := path.Ext(h.Filename) //获取文件后缀
		if ext != ".jpg" && ext != ".png" {
			beego.Info("上传文件格式错误")
			return
		}
		//限制图片大小
		if h.Size > 5000000 {
			beego.Info("上传文件过大")
			return
		}
		//对文件重命名，防止文件名重复
		if err != nil {
			beego.Info("文件上传失败",err)
			return
		}else {
			filename = time.Now().Format("20060102150405")+ ext
			beego.Info("开始保存文件啦~","./static/img/"+filename)
			c.SaveToFile("uploadname","./static/img/"+filename)
		}
	}
	o := orm.NewOrm()
	art := models.Article{Id:id}
	err = o.Read(&art)
	if err != nil {
		beego.Info("查询失败",err)
		return
	}
	art.Title = title
	art.Acontent = context
	art.Aimg = "/static/img/" + filename
	_,err = o.Update(&art,"Title","Acontent","Aimg")
	if err != nil {
		beego.Info("更新失败",err)
		return
	}
	c.Redirect("/index",302)
}

//删除
func (c *MainController) DelGet()  {
	id,err := c.GetInt("id")
	if err!=nil {
		beego.Info("id错误",err)
		return
	}
	o := orm.NewOrm()
	art := models.Article{Id:id}
	err = o.Read(&art)
	if err!=nil {
		beego.Info("查询失败",err)
		return
	}
	_,err =o.Delete(&art)
	if err!=nil {
		beego.Info("删除失败",err)
		return
	}
	c.Redirect("/index",302)
}