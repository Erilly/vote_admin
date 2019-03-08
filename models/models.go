package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
	"crypto/md5"
	"strconv"
	"encoding/hex"
)


type Question struct {
	Id uint
	QuestionId string `orm:"index;type(char);size(8)" description:(问卷ID)`
	PageCut int8 `orm:"default(0)" description:(状态 0不分页 1分页)`
	Title string `orm:"size(255)" description:(问卷标题)`
	Description string `orm:"type(text)" description:(问卷描述)`
	Usertoken string `orm:"index;type(char);size(32)" description:(用户Token)`
	Status int8 `orm:"default(0)" description:(状态：0正常 1删除)`
	PublishStatus int8 `orm:"default(0)" description:(发布状态：0上线 1下线)`
	Ctime time.Time `orm:"auto_now_add;type(datetime)" description:(创建时间)`
	Mtime time.Time `orm:"auto_now;type(datetime)" description:(更新时间)`
}


type Selector struct{
	Id uint
	QuestionId string `orm:"index;type(char);size(8)" description:(问卷ID)`
	SelectorId string `orm:"index;type(char);size(32)" description:(问题ID)`
	Title string `orm:"size(255)" description:(问卷标题)`
	TemplateType int8 `orm:"default(1)" description:(模板类型)`
	Page int16 `orm:"default(1)" description:(模板类型)`
	Ctime time.Time `orm:"auto_now_add;type(datetime)" description:(创建时间)`
	Mtime time.Time `orm:"auto_now;type(datetime)" description:(更新时间)`
}

type Option struct{
	Id uint
	SelectorId string `orm:"index;type(char);size(32)" description:(问题ID)`
	OptionId string `orm:"index;type(char);size(32)" description:(选项ID)`
	Title string `orm:"size(255)" description:(问卷标题)`
	Pid uint `orm:"default(0)" description:(父级id)`
	Status int8 `orm:"default(0)" description:(状态：0正常 1删除)`
	Ctime time.Time `orm:"auto_now_add;type(datetime)" description:(创建时间)`
	Mtime time.Time `orm:"auto_now;type(datetime)" description:(更新时间)`
}


type AnswerLog struct{
	Id uint
	QuestionId string `orm:"index;type(char);size(8)" description:(问卷ID)`
	Content string `orm:"type(text)" description:(答案结果)`
	Token string `orm:"index;type(char);size(32)" description:(答题客户端标识)`
	UserId uint `orm:"default(0)" description:(答题用户id)`
	Ctime time.Time `orm:"auto_now_add;type(datetime)" description:(创建时间)`
	Mtime time.Time `orm:"auto_now;type(datetime)" description:(更新时间)`
}


type User struct{
	Id uint
	User string `orm:"index;type(char);size(32)" description:(用户名)`
	Token string `orm:"index;type(char);size(32)" description:(Token)`
	Status int8 `orm:"default(0)" description:(状态 0正常 1删除)`
	Ctime time.Time `orm:"auto_now_add;type(datetime)" description:(创建时间)`
	Mtime time.Time `orm:"auto_now;type(datetime)" description:(更新时间)`
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

func GetAllQuestion()([]*Question,error){
	o := orm.NewOrm()
	questions := make([]*Question,0)
	qs := o.QueryTable("vt_question")
	_,err:=qs.OrderBy("-ctime").All(&questions)

	return questions,err
}

func GetQuestion(questionID string)(*Question,error){
	o := orm.NewOrm()
	questions := new(Question)
	qs := o.QueryTable("vt_question")
	err:=qs.Filter("question_id",questionID).One(questions)
	return questions,err
}

func AddQuestion(title,content string) (int64,string){
	now:=time.Now()
	token:=getMd5Token(strconv.FormatInt(time.Now().UTC().UnixNano(), 10),8)

	o := orm.NewOrm()

	question:=&Question{
		Title:title,
		Description:content,
		QuestionId:token,
		Ctime:now,
		Mtime:now,
	}

	id,err:=o.Insert(question)
	if err==nil{

	}
	return id,token
}

/*
	获取加密参数
 */
func getMd5Token(s string, l int) string{
	h := md5.New()
	h.Write([]byte(s))
	token:=hex.EncodeToString(h.Sum(nil))

	return token[:l]
}