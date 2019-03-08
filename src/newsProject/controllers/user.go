package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"newsProject/models"
	"newsProject/utils"
)

// 注册控制器
type RegController struct {
	beego.Controller
}

// 登陆注册器
type LoginController struct {
	beego.Controller
}

// 注册页面
func (this *RegController) ShowReg() {
	this.TplName = "register.html"
}

// 注册post
/*

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

func (this *LoginController) ShowLogin() {
	this.TplName = "login.html"
}
func (this *LoginController) HandLogin() {
	beego.Info("start login.")
	userName := this.GetString("userName")
	passWord := this.GetString("passWord")
	if userName == "" || passWord == "" {
		beego.Info(userName, "登陆失败: name or passwd empty")
		this.TplName = "login.html"
		return
	}
	o := orm.NewOrm()
	user := models.User{}
	user.UserName = userName
	err := o.Read(&user, "UserName")
	if err != nil {
		beego.Info(userName, "登陆失败: name / passwd not match", err)
		this.TplName = "login.html"
		return
	}
	if user.Password != utils.EncryptStr(passWord) {
		beego.Info(userName, "登陆失败: passwd / name not match", )
		this.TplName = "login.html"
		return
	}
	beego.Info("USERNAME: ", userName, "PASSWORD: ", passWord)
	fmt.Println(user.UserName, user.Password)
	beego.Info(userName, "登陆成功")
	this.Redirect("/ShowMenu",302)
}
