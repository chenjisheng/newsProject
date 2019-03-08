package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 表的设计
// 用户表
type User struct {
	Id int
	UserName string
	Password string
	Articles []*Article `orm:"rel(m2m)"`
}

// 文章列表 和 文章类型是一对多
type Article struct {
	Id int `orm:"auto"`
	Title string `orm:"size(20)"`// 标题
	Content string `orm:"size(500)"` // 内容
	Img string `orm:"size(50);null"`// 图片
	Time time.Time `orm:"type(datetime);auto_now_add"`// 发布时间
	Count int `orm:"default(0)"`// 阅读量
	ArticleType *ArticleType `orm:"rel(fk)"`// 文章类型
	User []*User `orm:"reverse(many)"`
}

// 文章类型
type ArticleType struct {
	Id int
	TypeName string `orm:"size(20)"` // 文章类型
	Articles []*Article `orm:"reverse(many)"`
}
func init() {
	driverName := "mysql"
	mysqlUser := beego.AppConfig.String("mysqluser")
	mysqlPassword := beego.AppConfig.String("mysqlpass")
	mysqlUrls := beego.AppConfig.String("mysqlurls")
	mysqlPort := beego.AppConfig.String("mysqlport")
	mysqlDbName := beego.AppConfig.String("mysqldb")
	// username:passwd@tcp(ip:port)/dbname?charset=utf8
	dataSource := mysqlUser + ":" + mysqlPassword + "@" + "tcp(" + mysqlUrls + ":" + mysqlPort + ")/" + mysqlDbName + "?charset=utf8"
	orm.RegisterDriver(driverName,orm.DRMySQL)
	orm.RegisterDataBase("default",driverName,dataSource)
	orm.RegisterModel(new(User),new(Article),new(ArticleType))
	orm.RunSyncdb("default",false,true)
}
