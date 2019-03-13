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
	Articles []*Article `orm:"rel(m2m)"` // 一个用户对应多个文章,同时一个文章可以被多个用户查看
}

// 文章列表
type Article struct {
	Id int `orm:"auto"`
	Title string `orm:"size(20)"`// 标题
	Content string `orm:"size(500)"` // 内容
	Img string `orm:"size(50);null"`// 图片
	Time time.Time `orm:"type(datetime);auto_now_add"`// 发布时间
	Count int `orm:"default(0)"`// 阅读量
	ArticleType *ArticleType `orm:"rel(fk);null;on_delete(set_null)"` // 一个类型对应多个文章
	User []*User `orm:"reverse(many)"`  //
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
	// 注册数据库驱动
	orm.RegisterDriver(driverName,orm.DRMySQL)
	// 注册数据库
	orm.RegisterDataBase("default",driverName,dataSource)
	// 注册数据库表
	orm.RegisterModel(new(User),new(Article),new(ArticleType))
	// 同步数据库,第二个参数是否强制更改数据库表结构变化
	orm.RunSyncdb("default",false,true)
}
