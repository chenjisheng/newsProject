package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"newsProject/models"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type ArticleController struct {
	beego.Controller
}

/*
添加文章类型页面
uri: /Article/AddArticleType
method: get
 */
func (this *ArticleController) ArticleTypePage() {
	this.Data["userName"] = this.Ctx.GetCookie("userName")
	this.Layout = "base.html"
	this.TplName = "addType.html"
}

/*
处理添加文章类型
uri: /Article/AddArticleType
method: post
@params: typeName string
 */
func (this *ArticleController) AddArticleType() {
	typeName := this.GetString("articleType")
	if typeName == "" {
		beego.Info("添加类型错误")
		this.Redirect("/Article/ArticleTypePage", 302)
		return
	}
	o := orm.NewOrm()
	var articlType models.ArticleType
	articlType.TypeName = typeName
	res := o.Read(&articlType, "TypeName")
	if res == nil {
		beego.Info("类型存在,不需要插入", res)
		this.Redirect("/Article/ArticleTypePage", 302)
		return
	}
	_, err := o.Insert(&articlType)
	if err != nil {
		beego.Info("添加文章类型失败")
	}
	this.Redirect("/Article/AddArticleType", 302)
}

/*
所有文章类型 api 接口
uri: /Article/ArticleTypeAll
method: get
return: json
 */
func (this *ArticleController) ArticleTypeAll() {
	o := orm.NewOrm()
	var datas = map[string]interface{}{}
	var counts int64
	var articleTypes []models.ArticleType
	_, err := o.QueryTable("ArticleType").All(&articleTypes)
	if err != nil {
		beego.Info("查询数据错误")
		datas["count"] = counts
		datas["code"] = 1
		datas["msg"] = "failed"
		datas["data"] = articleTypes
	} else {
		datas["code"] = 0
		datas["msg"] = "success"
		counts,_ = o.QueryTable("ArticleType").Count()
		datas["count"] = counts
		datas["data"] = articleTypes
	}
	this.Data["json"] = datas
	this.ServeJSON()
	return
}

/*
删除文章类型 api 接口
uri: /Article/DeleteArticleType/:id
method: post
return: json
 */
func (this *ArticleController) DeleteArticleType() {
	var datas = make(map[string]interface{})
	datas["code"] = 1
	datas["msg"] = "failed"
	id, err := this.GetInt(":id")
	if err != nil {
		datas["data"] = ""
		this.Data["json"] = datas
		this.ServeJSON()
		return
	}
	o := orm.NewOrm()
	articleType := models.ArticleType{Id: id}
	_, err = o.Delete(&articleType)
	if err != nil {
		datas["data"] = ""
		this.Data["json"] = datas
		this.ServeJSON()
		return
	}
	beego.Info("删除类型成功")
	datas["code"] = 0
	datas["msg"] = "successed"
	this.Data["json"] = datas
	this.ServeJSON()
	return

}

/*
编辑文章页面
uri: /Article/ArticleUpdate?id=
method: get
@params: id int
*/
func (this *ArticleController) ShowArticledetailUpdate() {
	o := orm.NewOrm()
	articleType := []models.ArticleType{}
	o.QueryTable("ArticleType").All(&articleType)
	id := this.GetString("id")
	newId, _ := strconv.Atoi(id)
	_, article := selectArticleData(newId)
	this.Data["articleType"] = articleType
	this.Data["article"] = article
	this.Data["userName"] = this.Ctx.GetCookie("userName")
	this.Layout = "base.html"
	this.TplName = "updateArticle.html"
	return
}

/*
处理文章更新
uri: /Article/ArticleUpdate?id=
method: post
@params: id int
 */
func (this *ArticleController) HandUpdate() {
	path_ := this.Ctx.Request.URL.Path
	query := this.Ctx.Request.URL.RawQuery
	url := path_ + "?" + query
	beego.Info("当前请求的URL: ", url)
	id := this.GetString("id")
	newId, _ := strconv.Atoi(id)
	o, article := selectArticleData(newId)
	articleName := this.GetString("articleName")
	articleContent := this.GetString("articleContent")
	articleImg := this.GetString("articleImg")
	TypeName := this.GetString("articleType")
	beego.Info("更新后的类型为: ", TypeName)
	// 类型判断
	if TypeName == "" {
		beego.Info("下拉框数据错误")
		this.Redirect("/Article/AddArticle", 302)
		return
	}
	var articleType models.ArticleType
	articleType.TypeName = TypeName
	err := o.Read(&articleType, "TypeName")
	if err != nil {
		beego.Info("获取类型错误", err)
		this.Redirect("/Article/AddArticle", 302)
		return
	}
	article.ArticleType = &articleType

	// 更新新值
	article.Title = articleName
	article.Content = articleContent
	article.Img = articleImg
	_, err = o.Update(&article)
	if err != nil {
		beego.Info("更新失败", err)
	}
	this.Redirect("/Article/ShowMenu", 302)
}

/*
删除文章 api 接口
uri: /Article/ArticleDelete/:id
method: get
 */
func (this *ArticleController) HandDelete() {
	id := this.GetString(":id")
	newId, _ := strconv.Atoi(id)
	o, article := selectArticleData(newId)
	imgName := "static/img/" + filepath.Base(article.Img)
	_, err := os.Stat(imgName)
	if err != nil {
		beego.Info("文件不存在", err)
	} else {
		beego.Info("文件存在,开始删除文件")
		os.Remove(imgName)
	}
	var datas = map[string]interface{}{}
	datas["code"] = 0
	datas["msg"] = ""
	datas["data"] = article
	this.Data["json"] = datas
	o.Delete(&article, "Id")
	this.ServeJSON()
	return
}

/*
显示文章详情
uri: /Article/ArticleDetail/:id
method: get
 */
func (this *ArticleController) ShowArticleDetail() {
	id := this.GetString(":id")
	newId, _ := strconv.Atoi(id)
	o, article := selectArticleData(newId)
	//article := models.Article{Id:newId}
	article.Count += 1
	// 根据文章表的文章类型Id 查询文章类型
	var articleType models.ArticleType
	articleType.Id = article.ArticleType.Id
	o.Read(&articleType, "Id")
	article.ArticleType.TypeName = articleType.TypeName
	beego.Info(article.ArticleType.TypeName)

	o.Update(&article, "Count")
	this.Data["article"] = article
	this.Data["userName"] = this.Ctx.GetCookie("userName")
	this.Layout = "base.html"
	this.TplName = "articleDetail.html"
}

/*
主页面
uri: /Article/ShowMenu
method: get
 */
func (this *ArticleController) ShowMenu() {
	o := orm.NewOrm()
	articleType := []models.ArticleType{}
	o.QueryTable("ArticleType").All(&articleType)
	this.Data["userName"] = this.Ctx.GetCookie("userName")
	this.Data["articleType"] = articleType
	this.Layout = "base.html"
	this.TplName = "articleList.html"
}

/*
文章列表
uri: /Article/ShowArticle
method: get
@params: page int
		 limit int
		 articletype string
 */
func (this *ArticleController) ShowArticleList() {
	// 当前页面数
	page, err := this.GetInt("page")
	// 当前页面显示数量
	limit, err := this.GetInt("limit")
	// 查询的类型
	typeName := this.GetString("articleType")
	o := orm.NewOrm()
	var articles []models.Article // 文章表
	qs := o.QueryTable("Article")
	var counts int64
	if typeName == "" {
		beego.Info("查询的文章类型为: 所有")
		qs.Limit(limit, limit*(page-1)).RelatedSel("ArticleType").All(&articles) // 1. pagesize 2. start 数据库限制查询;
		counts, err = qs.RelatedSel("ArticleType").Count()
		if err != nil {
			beego.Info("查询总数错误.")
		}
	} else {
		beego.Info("查询的文章类型为: ", typeName)
		qs.Limit(limit, limit*(page-1)).RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).All(&articles) // 1. pagesize 2. start 数据库限制查询;
		counts, err = qs.RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).Count()
		if err != nil {
			beego.Info("查询总数错误.")
		}
	}
	beego.Info("PAGE: ", page, "LIMIT: ", limit)
	// 查询数据
	// 将数据传递给视图
	//_, err = qs.All(&articles) // select * from article;
	var datas = map[string]interface{}{}
	datas["code"] = 0
	datas["msg"] = ""
	datas["count"] = counts
	datas["data"] = articles
	this.Data["json"] = datas
	this.ServeJSON()
	return
}

/*
显示添加文章
uri: /Article/AddArticle
method: get
 */
func (this *ArticleController) ShowAddArticle() {
	o := orm.NewOrm()
	var articleType []models.ArticleType
	_, err := o.QueryTable("ArticleType").All(&articleType)
	if err != nil {
		beego.Info("查询文章类型错误", err)
	}
	beego.Info(articleType)
	this.Data["userName"] = this.Ctx.GetCookie("userName")
	this.Data["articleType"] = articleType
	this.Layout = "base.html"
	this.TplName = "addArticle.html"
}

/*
处理添加文章
uri: /Article/AddArticle
method: post
@params: articelName string
		 articelContent string
		 articleImg string
		 articleType string
  */
func (this *ArticleController) HandAddArticle() {
	// 获取数据
	// 查询数据
	// 插入数据
	o := orm.NewOrm()
	article := models.Article{}
	articleName := this.GetString("articleName")
	articleContent := this.GetString("articleContent")
	articleImg := this.GetString("articleImg")
	article.Title = articleName
	article.Content = strings.Replace(articleContent, "\n", "", 10)
	article.Time = time.Now().Add(time.Second * 28800)
	article.Img = articleImg
	// 获取到下拉框的数据
	TypeName := this.GetString("articleType")
	beego.Info("插入的类型为: ", TypeName)
	// 类型判断
	if TypeName == "" {
		beego.Info("下拉框数据错误")
		this.Redirect("/Article/AddArticle", 302)
		return
	}
	var articleType models.ArticleType
	articleType.TypeName = TypeName
	err := o.Read(&articleType, "TypeName")
	if err != nil {
		beego.Info("获取类型错误", err)
		this.Redirect("/Article/AddArticle", 302)
		return
	}
	article.ArticleType = &articleType
	_, err = o.Insert(&article)
	if err != nil {
		beego.Info("插入数据失败")
		this.Redirect("/Article/AddArticle", 302)
		return
	}
	this.Redirect("/Article/ShowMenu", 302)
}

/*
上传图片 api 接口
uri: /Article/UploadImg
method: post
@params: file []byte
 */
func (this *ArticleController) HandUploadImg() {
	file, header, err := this.Ctx.Request.FormFile("file")
	// 检查是否上传文件
	code := 1
	errmsg := "failed"
	infomsg := "successed"
	var datas = make(map[string]interface{})
	if err != nil {
		beego.Info("未上传文件")
	} else {
		defer file.Close()
		beego.Info("上传的文件名称为: ", header.Filename)
		// 判断文件格式
		// 判断文件大小
		// 存的时候不能重名
		ext := path.Ext(header.Filename)
		if ext == ".jpg" || ext == ".png" || ext == ".jpeg" || ext == ".pdf" {
			beego.Info("上传文件格式正确")
		} else {
			beego.Info("上传文件格式不正确")
			datas["code"] = code
			datas["msg"] = errmsg + "file error"
			this.Data["json"] = datas
			this.ServeJSON()
			return
		}
		// 设置阀值为 5 M 1024 *1024 * 5 = 5M
		if header.Size >= 5000000 {
			currentSize := header.Size / 1024 / 1024
			beego.Info("文件过大,不允许上传,当前文件大小为: ", currentSize, "M")
			datas["code"] = code
			datas["msg"] = errmsg + "large file"
			this.Data["json"] = datas
			this.ServeJSON()
			return
		}
		fileName := time.Now().Format("2006-01-02-15-04-05")
		err = this.SaveToFile("file", "static/img/"+fileName+ext)
		if err != nil {
			beego.Info("Save file", fileName+ext, "Falied.", err)
			datas["code"] = code
			datas["msg"] = errmsg + "save file"
			this.Data["json"] = datas
			this.ServeJSON()
			return
		}
		datas["code"] = 0
		datas["msg"] = infomsg
		datas["data"] = map[string]string{"url": "/static/img/" + fileName + ext}
		this.Data["json"] = datas
		this.ServeJSON()
		return
	}

}

/*
根据 ID 查询数据库 article 表,返回 ORM 对象以及文章对象
 */
func selectArticleData(id int) (o orm.Ormer, article models.Article) {
	o = orm.NewOrm()
	article = models.Article{Id: id}
	err := o.Read(&article)
	if err != nil {
		beego.Info("未查询到文章,ID: ", id)
	} else {
		beego.Info("查询文章详情成功,ID: ", id)
	}
	// 返回数据为: 对象及结果
	return o, article
}
