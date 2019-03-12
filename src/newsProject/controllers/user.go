package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"newsProject/models"
	"newsProject/utils"
	"time"
)

// 注册控制器
type RegController struct {
	beego.Controller
}

// 登陆注册器
type LoginController struct {
	beego.Controller
}

/*
注册页面
uri: /Register
method: get
 */
func (this *RegController) ShowReg() {
	this.TplName = "register.html"
}

/*
处理注册
uri: /Register
method: post
 */
func (this *RegController) HandReg() {
	// 获取前端数据
	userName := this.GetString("userName")
	passWord := this.GetString("passWord")
	// 数据处理
	if userName == "" || passWord == "" {
		beego.Info("用户名或密码不能为空")
		this.TplName = "register.html"
		return
	}
	// 插入数据库
	o := orm.NewOrm()
	user := models.User{}
	// 1. 获取 orm 对象
	// 2. 获取插入对象
	// 3. 插入操作
	user.UserName = userName
	err := o.Read(&user,"UserName")
	if err != nil {
		beego.Info("用户不存在,可以注册",err)
	} else {
		beego.Info("用户存在,不可以注册",err)
		this.TplName = "register.html"
		return
	}
	// 加密密码
	enPassword := utils.EncryptStr(passWord)
	user.Password = enPassword
	_, err = o.Insert(&user)
	if err != nil {
		beego.Info("注册用户失败",err)
		this.TplName = "register.html"
		return
	}
	beego.Info("注册用户成功")
	this.Redirect("/",302)
	return
	// 4. 返回登陆
}

/*
登陆页面
uri: /
method: get
 */
func (this *LoginController) ShowLogin() {
	name := this.Ctx.GetCookie("userName")
	if name != ""{
		this.Data["userName"] = name
		this.Data["check"] = "checked"
	}
	this.TplName = "login.html"
}

/*
处理登陆 api 接口
uri: /
method: post
@params: userName string
		 password string
		 remember string
 */
func (this *LoginController) HandLogin() {
	beego.Info("start login.")
	userName := this.GetString("userName")
	passWord := this.GetString("passWord")
	remember := this.GetString("remember")
	var datas = make(map[string]interface{})
	datas["code"] = 0
	datas["msg"] = ""
	if userName == "" || passWord == "" {
		beego.Info(userName, "登陆失败: name or passwd empty")
		datas["code"] = 1
		datas["msg"] = "name or password error"
		datas["data"] = ""
		this.Data["json"] = datas
		this.ServeJSON()
		return
	}
	o := orm.NewOrm()
	user := models.User{}
	user.UserName = userName
	err := o.Read(&user, "UserName")
	if err != nil {
		beego.Info(userName, "登陆失败: name / passwd not match", err)
		datas["code"] = 1
		datas["msg"] = "name or password error"
		datas["data"] = ""
		this.Data["json"] = datas
		this.ServeJSON()
		return
	}
	if user.Password != utils.EncryptStr(passWord) {
		beego.Info(userName, "登陆失败: passwd / name not match", )
		datas["code"] = 1
		datas["msg"] = "name or password error"
		datas["data"] = ""
		this.Data["json"] = datas
		this.ServeJSON()
		return
	}
	if remember == "on" {
		beego.Info("开启记录用户名")
		this.Ctx.SetCookie("userName",userName,time.Second*3600)
	} else {
		this.Ctx.SetCookie("userName",userName,-1)
	}
	beego.Info("USERNAME: ", userName, "PASSWORD: ", passWord)
	beego.Info(userName, "登陆成功")
	this.SetSession("userName",userName)
	datas["data"] = map[string]interface{}{"access_token":this.GetSession("userName")}
	this.Data["json"] = datas
	this.ServeJSON()
	return
	//this.Redirect("/Article/ShowMenu",302)
}

/*
退出登陆
uri: /Logout
method: get
 */
func (this *LoginController) Logout() {
	// 1. 删除登陆状态
	// 2. 跳转到登陆页面
	userName := this.GetSession("userName")
	beego.Info(userName,"退出登陆.")
	this.DelSession("userName")
	this.Redirect("/",302)
}
