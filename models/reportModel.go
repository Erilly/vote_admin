package models

import (
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
	Option []*Opt
}

type Opt struct {
	Id int
	Title string
	Pid int
	Status int8
	Ctime time.Time
	Mtime time.Time
	Lables string
	Data string
}

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

		Select:=new(Select)
		Select.Id = selector.Id
		Select.Title = selector.Title
		Select.TemplateType = selector.TemplateType
		Select.Page = selector.Page
		Select.Ctime = selector.Ctime
		Select.Mtime = selector.Mtime

		for _,option:=range selector.Option{
			Opt:=new(Opt)
			Opt.Id = option.Id
			Opt.Title = option.Title
			Opt.Pid = option.Pid
			Opt.Status = option.Status
			Opt.Ctime = option.Ctime
			Opt.Mtime = option.Mtime
			Opt.Lables = ""
			Opt.Data = ""

			Select.Option = append( Select.Option, Opt )
		}

		Quest.Selector = append(Quest.Selector, Select)
	}

	return Quest
}