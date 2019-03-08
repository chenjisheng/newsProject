package routers

import (
	"newsProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    // web icon router 网页icon
    beego.Router("/favicon.ico",&controllers.MainController{})
    // register router 注册
    beego.Router("/Register",&controllers.RegController{},"get:ShowReg;post:HandReg")  // 用户注册
    // login router 登陆
	beego.Router("/",&controllers.LoginController{},"get:ShowLogin;post:HandLogin")
    // show articles router 登陆后显示
    beego.Router("/ShowMenu",&controllers.ArticleController{},"get:ShowMenu")
	// 文章列表
    beego.Router("/ShowArticle",&controllers.ArticleController{},"get:ShowArticleList")

    // upload image router 添加照片
    beego.Router("/UploadImg",&controllers.ArticleController{},"post:HandUploadImg")
    //  add article router 添加文章
    beego.Router("/AddArticle",&controllers.ArticleController{},"get:ShowAddArticle;post:HandAddArticle")
    // article details router 文章详情
    beego.Router("/ArticleDetail/:id",&controllers.ArticleController{},"get:ShowArticleDetail")
    // delete article router 删除文章
    beego.Router("/ArticleDelete/:id",&controllers.ArticleController{},"get:HandDelete")
    // update article router 更新文章
    beego.Router("/ArticleUpdate",&controllers.ArticleController{},"get:ShowArticledetailUpdate;post:HandUpdate")
    // article type show
    beego.Router("/ArticleTypePage",&controllers.ArticleController{},"get:ArticleTypePage")
    // show  article type 返回所有文章类型
    beego.Router("/ArticleTypeAll",&controllers.ArticleController{},"get:ArticleTypeAll")
    // add article type
    beego.Router("/AddArticleType",&controllers.ArticleController{},"post:AddArticleType")
    // delete article type
    beego.Router("/DeleteArticleType/:id",&controllers.ArticleController{},"post:DeleteArticleType")
}
