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

// 显示文章内容,加载更新页面
// uri = ArticleUpdate?id=11
func (this *ArticleController) ShowArticledetailUpdate() {
	id := this.GetString("id")
	newId, _ := strconv.Atoi(id)
	_, article := selectData(newId)
	this.Data["article"] = article
	this.Layout = "base.html"
	this.TplName = "updateArticle.html"
	return
}

// 提交更新
// uri = ArticleUpdate?id=33
func (this *ArticleController) HandUpdate() {
	path_ := this.Ctx.Request.URL.Path
	query := this.Ctx.Request.URL.RawQuery
	url := path_ + "?" + query
	beego.Info("当前请求的URL: ", url)
	id := this.GetString("id")
	newId, _ := strconv.Atoi(id)
	o, article := selectData(newId)
	articleName := this.GetString("articleName")
	articleContent := this.GetString("articleContent")
	articleImg := this.GetString("articleImg")
	// 更新新值
	article.Title = articleName
	article.Content = articleContent
	article.Img = articleImg
	_, err := o.Update(&article)
	if err != nil {
		beego.Info("更新失败", err)
	}
	this.Redirect("/ShowMenu", 302)
}

// 删除文章
// uri = ArticleDelete/11
func (this *ArticleController) HandDelete() {
	id := this.GetString(":id")
	newId, _ := strconv.Atoi(id)
	o, article := selectData(newId)
	imgName := "static/img/"+filepath.Base(article.Img)
	_,err := os.Stat(imgName)
	if err != nil {
		beego.Info("文件不存在",err)
	} else{
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

// 显示文章详情
func (this *ArticleController) ShowArticleDetail() {
	id := this.GetString(":id")
	newId, _ := strconv.Atoi(id)
	o, article := selectData(newId)
	//article := models.Article{Id:newId}
	article.Count += 1
	o.Update(&article, "Count")
	this.Data["article"] = article
	this.Layout = "base.html"
	this.TplName = "articleDetail.html"
}

// 显示主页
func (this *ArticleController) ShowMenu(){
	this.Layout = "base.html"
	this.TplName = "articleList.html"
}
// 根据前端请求分页
func (this *ArticleController) ShowArticleList() {
	page,err := this.GetInt("page")
	limit,err := this.GetInt("limit")
	beego.Info("PAGE: ",page,"LIMIT: ",limit)
	// 查询数据
	// 将数据传递给视图
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	var articles []models.Article
	//_, err = qs.All(&articles) // select * from article;
	qs.Limit(limit,limit*(page-1)).All(&articles)  // 1. pagesize 2. start 数据库限制查询;
	counts,err := qs.Count()
	if err != nil {
		beego.Info("查询总数错误.")
	}
	//datas =
	var datas = map[string]interface{}{}
	datas["code"] = 0
	datas["msg"] = ""
	datas["count"] = counts
	datas["data"] = articles
	this.Data["json"] = datas
	this.ServeJSON()
	return
}

// 显示添加文章
func (this *ArticleController) ShowAddArticle() {
	this.Layout = "base.html"
	this.TplName = "addArticle.html"
}

// 添加文章
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
	_, err := o.Insert(&article)
	if err != nil {
		beego.Info("插入数据失败")
		this.Redirect("/AddArticle", 302)
		return
	}
	this.Redirect("/ShowMenu", 302)
}

// 上传图片
func (this *ArticleController) HandUploadImg(){
	file,header,err:= this.Ctx.Request.FormFile("file")
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
		datas["data"] = map[string]string{"url":"/static/img/"+fileName+ext}
		this.Data["json"] = datas
		this.ServeJSON()
		return
	}

}
// 根据ID 查询数据库,返回 ORM 对象
func selectData(id int) (o orm.Ormer, article models.Article) {
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
