package models

import (
	"strconv"
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

const (
	//题型
	SINGLE_SELECTOTR = 1
	MULTI_SELECTOTR = 2
	SCORE_SELECTOTR = 3
	SCORE_MATRIX_SELECTOTR = 4
	FILL_SELECTOTR = 5
	FILL_MATRIX_SELECTOTR = 6
)

type Question struct {
	Id int
	PageCut int8 `orm:"default(0)" description:"状态 0不分页 1分页"`
	Title string `orm:"size(255)" description:"问卷标题"`
	Description string `orm:"type(text)" description:"问卷描述"`
	Usertoken string `orm:"index;type(char);size(32)" description:"用户Token"`
	Status int8 `orm:"default(0)" description:"状态：0正常 1删除"`
	PublishStatus int8 `orm:"default(0)" description:"发布状态：0上线 1下线"`
	Ctime time.Time `orm:"auto_now_add;type(datetime)" description:"创建时间"`
	Mtime time.Time `orm:"auto_now;type(datetime)" description:"更新时间"`
	Selector []*Selector `orm:"reverse(many)"`
}


type Selector struct{
	Id int
	Title string `orm:"size(255)" description:"问卷标题"`
	TemplateType int8 `orm:"default(1)" description:"模板类型"`
	Page int16 `orm:"default(1)" description:"模板类型"`
	Ctime time.Time `orm:"auto_now_add;type(datetime)" description:"创建时间"`
	Mtime time.Time `orm:"auto_now;type(datetime)" description:"更新时间"`
	Option []*Option `orm:"reverse(many)"`
	Question  *Question  `orm:"rel(fk)"`
}

type Option struct{
	Id int
	Title string `orm:"size(255)" description:"问卷标题"`
	Pid int `orm:"default(0)" description:"父级id"`
	Status int8 `orm:"default(0)" description:"状态：0正常 1删除"`
	Ctime time.Time `orm:"auto_now_add;type(datetime)" description:"创建时间"`
	Mtime time.Time `orm:"auto_now;type(datetime)" description:"更新时间"`
	Selector  *Selector  `orm:"rel(fk)"`
}

type AnswerLog struct{
	Id int
	QuestionId string `orm:"index;type(char);size(8)" description:"问卷ID"`
	Content string `orm:"type(text)" description:"答案结果"`
	Token string `orm:"index;type(char);size(32)" description:"答题客户端标识"`
	UserId int `orm:"default(0)" description:"答题用户id"`
	Ctime time.Time `orm:"auto_now_add;type(datetime)" description:"创建时间"`
	Mtime time.Time `orm:"auto_now;type(datetime)" description:"更新时间"`
}


type User struct{
	Id int
	User string `orm:"index;type(char);size(32)" description:"用户名"`
	Token string `orm:"index;type(char);size(32)" description:"Token"`
	Status int8 `orm:"default(0)" description:"状态 0正常 1删除"`
	Ctime time.Time `orm:"auto_now_add;type(datetime)" description:"创建时间"`
	Mtime time.Time `orm:"auto_now;type(datetime)" description:"更新时间"`
}

func RegisterDB() {

	database_link := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		beego.AppConfig.String("user"),
		beego.AppConfig.String("pass"),
		beego.AppConfig.String("host"),
		beego.AppConfig.String("port"),
		beego.AppConfig.String("database"))

	//注册数据模型
	orm.RegisterModelWithPrefix( beego.AppConfig.String("prefix") ,
		new(Question),
		new(Selector),
		new(Option),
		new(AnswerLog),
		new(User))

	//指定数据库类型
	orm.RegisterDriver( beego.AppConfig.String("driver"), orm.DRMySQL)
	//连接数据库
	orm.RegisterDataBase("default", beego.AppConfig.String("driver"), database_link)
}

func (this *Question) TableName() string {
	return "question"
}

func (this *Selector) TableName() string {
	return "selector"
}

func (this *Option) TableName() string {
	return "option"
}

func (this *AnswerLog) TableName() string {
	return "answerLog"
}

func (this *User) TableName() string {
	return "user"
}


/*
	获取加密参数
 */
func getMd5Token(l int,s string) string{

	if s=="" {
		s = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	}
	h := md5.New()
	h.Write([]byte(s))
	token:=hex.EncodeToString(h.Sum(nil))

	return token[:l]
}