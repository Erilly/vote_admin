package models

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

type Quest struct {
	Id int
	PageCut int8
	Title string
	Description string
	Usertoken string
	Status int8
	PublishStatus int8
	Ctime time.Time
	Mtime time.Time
	Selector []*Select
}
type Select struct {
	Id int
	Title string
	TemplateType int8
	Page int16
	Ctime time.Time
	Mtime time.Time
	Lables string
	Color string
	Data string
	Sum string
	Totle int
	Option []*Opt
}

type Opt struct {
	Id int
	Title string
	Pid int
	Status int8
	Count int
	Color string
	Ctime time.Time
	Mtime time.Time
}

type FillData struct {
	Pagers string
	Data []interface{}
}

var paginatorMap map[string]interface{}

var color = []string{
	"#ff0000",
	"#26B99A",
	"#337ab7",
	"#f3a746",
	"#E76359",
	"#60b0ce",
	"#810082",
	"#7b1f36",
	"#E74C3C",
	"#3498DB",
	"#761f36",
	"#9B59B6",
	"#BDC3C7"}

func GetReport(question Question)(*Quest){

	Quest:=new(Quest)

	Quest.Id = question.Id
	Quest.PageCut = question.PageCut
	Quest.Title = question.Title
	Quest.Description = question.Description
	Quest.Usertoken = question.Usertoken
	Quest.Status = question.Status
	Quest.PublishStatus = question.PublishStatus
	Quest.Ctime = question.Ctime
	Quest.Mtime = question.Mtime

	for _,selector := range question.Selector{
		var lables []string
		var data []int
		var sum []int
		var start1 []int
		var start2 []int
		var start3 []int
		var start4 []int
		var start5 []int

		Select:=new(Select)
		Select.Id = selector.Id
		Select.Title = selector.Title
		Select.TemplateType = selector.TemplateType
		Select.Page = selector.Page
		Select.Ctime = selector.Ctime
		Select.Mtime = selector.Mtime

		for k,option:=range selector.Option{
			Opt:=new(Opt)
			Opt.Id = option.Id
			Opt.Title = option.Title
			Opt.Pid = option.Pid
			Opt.Status = option.Status
			Opt.Ctime = option.Ctime
			Opt.Mtime = option.Mtime

			switch selector.TemplateType{
				case SINGLE_SELECTOTR:
					lables=append(lables,option.Title)

					count,_:=GetDatabase().C(MONGO_COLLECTION_VOTE_LOG).Find(
						&bson.M{"question_id": question.Id,"selector_"+strconv.Itoa(selector.Id):strconv.Itoa(option.Id)}).Count()

					Opt.Count = count
					Opt.Color = color[k]

					data=append(data,count)

				case MULTI_SELECTOTR:

					count,_:=GetDatabase().C(MONGO_COLLECTION_VOTE_LOG).Find(
						&bson.M{"question_id": question.Id,"selector_"+strconv.Itoa(selector.Id):strconv.Itoa(option.Id)}).Count()

					Opt.Count = count
					data=append(data,count)

				case SCORE_SELECTOTR:
					lables=append(lables,option.Title)
					count1,_:=GetDatabase().C(MONGO_COLLECTION_VOTE_LOG).Find(
						bson.M{"question_id": question.Id,"selector_"+strconv.Itoa(selector.Id):bson.M{"opt_"+strconv.Itoa(option.Id):"1"}}).Count()

					start1=append(start1,count1)

					count2,_:=GetDatabase().C(MONGO_COLLECTION_VOTE_LOG).Find(
						bson.M{"question_id": question.Id,"selector_"+strconv.Itoa(selector.Id):bson.M{"opt_"+strconv.Itoa(option.Id):"2"}}).Count()

					start2=append(start2,count2)

					count3,_:=GetDatabase().C(MONGO_COLLECTION_VOTE_LOG).Find(
						bson.M{"question_id": question.Id,"selector_"+strconv.Itoa(selector.Id):bson.M{"opt_"+strconv.Itoa(option.Id):"3"}}).Count()

					start3=append(start3,count3)

					count4,_:=GetDatabase().C(MONGO_COLLECTION_VOTE_LOG).Find(
						bson.M{"question_id": question.Id,"selector_"+strconv.Itoa(selector.Id):bson.M{"opt_"+strconv.Itoa(option.Id):"4"}}).Count()

					start4=append(start4,count4)

					count5,_:=GetDatabase().C(MONGO_COLLECTION_VOTE_LOG).Find(
						bson.M{"question_id": question.Id,"selector_"+strconv.Itoa(selector.Id):bson.M{"opt_"+strconv.Itoa(option.Id):"5"}}).Count()

					start5=append(start5,count5)

					Opt.Count = count1 + count2*2 + count3*3 + count4*4 + count5*5
					sum=append(sum,Opt.Count)

				case FILL_SELECTOTR:
			}

			Select.Option = append( Select.Option, Opt )
		}

		var b int
		for _,value := range data{
			b += value
		}
		if b >0 {
			Select.Totle = b
		}

		l,_:=json.Marshal(lables)
		Select.Lables = string(l)

		if selector.TemplateType == SCORE_SELECTOTR {
			s, _ := json.Marshal(sum)
			Select.Sum = string(s)

			score := make(map[string][]int)

			score["start1"] = start1
			score["start2"] = start2
			score["start3"] = start3
			score["start4"] = start4
			score["start5"] = start5

			d, _ := json.Marshal(score)
			Select.Data = string(d)

		}else if selector.TemplateType == FILL_SELECTOTR {
			result := FillReport(question.Id,Select.Id,5,1)

			d,_:=json.Marshal(result)
			Select.Data = string(d)

		}else{

			d,_:=json.Marshal(data)
			Select.Data = string(d)

		}

		c,_:=json.Marshal(color)
		Select.Color = string(c)

		Quest.Selector = append(Quest.Selector, Select)
	}

	return Quest
}

func FillReport(qid,sid,limit,page int) *FillData {
	var result []interface{}

	if limit < 1{
		limit = 10
	}

	if page < 1{
		page =1
	}

	offset:=(page-1)*limit

	sql:=GetDatabase().C(MONGO_COLLECTION_VOTE_LOG).
		Find(bson.M{"question_id": qid}).
		Select(bson.M{"selector_"+strconv.Itoa(sid):1})

	total,_:=sql.Count()
	sql.Limit(limit).Skip(offset).All(&result)

	pagerMap:=Paginator(page,limit,int64(total))

	FillData:=new(FillData)
	FillData.Pagers = FormatePagerHtml(pagerMap,qid,sid)
	FillData.Data = result

	return FillData
}