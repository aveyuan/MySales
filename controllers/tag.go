package controllers

import "sales-project/models"

type TagController struct {
	BaseController
}

func (this *TagController) List() {
	tag := &models.Tag{}
	this.Data["tag"] = tag.List()
	this.Data["page"] = "标签列表"
	this.Layout = "public/layout.html"
	this.TplName = "tag/list.html"
}

func (this *TagController) Add() {
	if this.IsPost() {
		tagname := this.GetString("name")
		tag := &models.Tag{Name: tagname}
		if err := tag.Add(); err != nil {
			this.Ctx.WriteString("添加失败")
		}
		this.Redirect(this.URLFor(".List"), 302)
	} else {
		this.Xsrf()
		this.Data["page"] = "添加标签"
		this.Layout = "public/layout.html"
		this.TplName = "tag/add.html"
	}
}

func (this *TagController) Update() {
	if this.IsPost() {
		id, err := this.GetInt("id")
		if err != nil {
			this.Ctx.WriteString("获ID有误")
			this.StopRun()
		}
		name := this.GetString("name")
		tag := &models.Tag{Id: id, Name: name}
		if err := tag.Update(); err != nil {
			this.Ctx.WriteString("更新失败")
		} else {
			this.Redirect(this.URLFor(".List"), 302)
		}
	} else {
		id, err := this.GetInt("id")
		if err != nil {
			this.Ctx.WriteString("数据有误")
		}
		tag := &models.Tag{Id: id}
		this.Xsrf()
		this.Data["tag"] = tag.Idtag()
		this.Data["pagetitle"] = "修改标签信息页面"
		this.Layout = "public/layout.html"
		this.TplName = "tag/update.html"
	}
}
