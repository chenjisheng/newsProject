package routers

import (
    "github.com/astaxie/beego/context"
    "newsProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    // 路由过滤器,判断用户是否登陆
    beego.InsertFilter("/Article/*",beego.BeforeRouter,FilterFunc)
    // 网页icon, 网站小图片
    beego.Router("/favicon.ico",&controllers.MainController{})
    // 注册用户
    beego.Router("/Register",&controllers.RegController{},"get:ShowReg;post:HandReg")
    // 登陆 api 接口
	beego.Router("/",&controllers.LoginController{},"get:ShowLogin;post:HandLogin")
    // 退出登陆
    beego.Router("/Article/Logout",&controllers.LoginController{},"get:Logout")
    // 主页面
    beego.Router("/Article/ShowMenu",&controllers.ArticleController{},"get:ShowMenu")
	// 文章列表 api 接口
    beego.Router("/Article/ShowArticle",&controllers.ArticleController{},"get:ShowArticleList")
    // 上传照片 api 接口
    beego.Router("/Article/UploadImg",&controllers.ArticleController{},"post:HandUploadImg")
    // 添加文章
    beego.Router("/Article/AddArticle",&controllers.ArticleController{},"get:ShowAddArticle;post:HandAddArticle")
    // 文章详情
    beego.Router("/Article/ArticleDetail/:id",&controllers.ArticleController{},"get:ShowArticleDetail")
    // 删除文章 api 接口
    beego.Router("/Article/ArticleDelete/:id",&controllers.ArticleController{},"get:HandDelete")
    // 编辑文章页面
    beego.Router("/Article/ArticleUpdate",&controllers.ArticleController{},"get:ShowArticledetailUpdate;post:HandUpdate")
    // 添加文章类型
    beego.Router("/Article/AddArticleType",&controllers.ArticleController{},"get:ArticleTypePage;post:AddArticleType")
    // 所有文章类型 api 接口
    beego.Router("/Article/ArticleTypeAll",&controllers.ArticleController{},"get:ArticleTypeAll")
   // 删除文章类型 api 接口
    beego.Router("/Article/DeleteArticleType/:id",&controllers.ArticleController{},"post:DeleteArticleType")
}

var FilterFunc = func(ctx *context.Context) {
    userName := ctx.Input.Session("userName")
    if userName == nil {
        ctx.Redirect(302,"/")
    }
}