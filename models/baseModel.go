package models

import (
	"math"
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
	SINGLE_SELECTOTR = 1	//单选
	MULTI_SELECTOTR  = 2	//多选
	SCORE_SELECTOTR  = 3	//打分
	FILL_SELECTOTR   = 4	//填空
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
	return "answer_log"
}

func (this *User) TableName() string {
	return "user"
}


/**
 *获取加密参数
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

/**
 *分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
 */
func Paginator(page, limit int, nums int64) map[string]interface{} {

	var firstpage int //前一页地址
	var lastpage int  //后一页地址
	//根据nums总数，和limit每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(limit))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		firstpage = page - 1
		lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		firstpage = page - 1
		lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		firstpage = int(math.Max(float64(1), float64(page-1)))
		lastpage = page + 1
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = firstpage
	paginatorMap["lastpage"] = lastpage
	paginatorMap["currpage"] = page
	paginatorMap["totals"] = nums

	return paginatorMap
}