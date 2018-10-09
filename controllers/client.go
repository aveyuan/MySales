package controllers

import (
	"sales-project/models"
)

type ClientController struct {
	BaseController
}

func (this *ClientController)Add()  {
	this.IsLogin()
	if this.IsPost(){
		name := this.Ctx.Request.PostForm.Get("name")
		phone := this.Ctx.Request.PostForm.Get("phone")
		address := this.Ctx.Request.PostForm.Get("address")
		postid := this.Ctx.Request.PostForm.Get("postid")
		tagid,_ := this.GetInt("tag")
		tag := &models.Tag{Id:tagid}
		remarks := this.Ctx.Request.PostForm.Get("remarks")
		client := &models.Client{Name:name,Tag:tag,Phone:phone,Address:address,Postid:postid,Remarks:remarks,Createtime:this.GetDateTime()}
		if err := client.Add();err !=nil{
			this.Ctx.WriteString("添加失败")
		}
		this.Redirect(this.URLFor(".List"),302)

	}else {
		tag := &models.Tag{}
		this.Xsrf()
		this.Data["pagetitle"]="新增客户"
		this.Data["tag"]=tag.List()
		this.Layout="public/layout.html"
		this.TplName="client/add.html"
	}
}

func (this *ClientController)List()  {
	this.IsLogin()
	limit,err := this.GetInt("limit")
	if err !=nil{
		limit=10
	}
	page,err := this.GetInt("page")
	if err != nil{
		page=1
	}
	key := this.GetString("key")
	if key == ""{
		key = "*"
	}
	//如果提交的方式是搜索来的，必须定向到第一页
	if this.IsPost(){
		limit=10
		page=1
	}
	client := &models.Client{}
	clients,snum := client.ListLimit(limit,page,key)
	//有没有更好的方式，在不重组数据的情况下，直接通过orm获得所需要的值？这点太不方便了。
	list := make([]map[string]interface{},len(clients))
	for k,v := range clients{
		row := make(map[string]interface{})
		//考虑到版本升级的问题，加以判断
		if v.Tag==nil || v.Tag.Id==0{
			tag := &models.Tag{Name:""}
			row["tag"]=tag
		}else {
			tag := v.Tag.Idtag()
			row["tag"]=tag
		}
		row["id"]=v.Id
		row["name"]=v.Name
		row["phone"]=v.Phone
		row["address"]=v.Address
		row["postid"]=v.Postid
		row["remarks"]=v.Remarks
		list[k]=row
	}
	this.Data["clients"]=list
	this.Data["pagetitle"]="用户列表"
	//为了区分全搜索还是局部搜索要再次判断key
	if key == "*"{
		this.Data["key"]=""
	}else{
		this.Data["key"]=key
	}
	this.Data["pagecount"]=len(snum)
	this.Data["pagelimit"]=limit
	this.Data["page"]=page
	this.Xsrf()
	this.Layout="public/layout.html"
	this.TplName="client/list.html"
}

func (this *ClientController)Update()  {
	this.IsLogin()
	if this.IsPost(){
		id,err := this.GetInt("id")
			if err!=nil{
				this.Ctx.WriteString("获ID有误")
				this.StopRun()
			}
		tagid,err := this.GetInt("tag")
		tag := &models.Tag{Id:tagid}
		name := this.GetString("name")
		address := this.GetString("address")
		postid := this.GetString("postid")
		phone := this.GetString("phone")
		remarks := this.GetString("remarks")
		client := &models.Client{Id:id,Name:name,Address:address,Postid:postid,Tag:tag,Phone:phone,Remarks:remarks,Updatetime:this.GetDateTime()}
		if err := client.Update();err!=nil{
			this.Ctx.WriteString("更新失败")
		}else{
			this.Redirect(this.URLFor(".List"),302)
		}
	}else {
		id,err := this.GetInt("id")
		if err!=nil{
			this.Ctx.WriteString("数据有误")
		}
		client := &models.Client{Id:id}
		clients := client.IdClinet()
		//判断是会否有tag再处理
		if client.Tag !=nil{
			tag := client.TagGet()
			this.Data["tag"]=tag
		}else {
			tag := &models.Tag{Id:0,Name:""}
			cs := &models.Client{Tag:tag}
			this.Data["tag"]=cs
		}

		tag1 := &models.Tag{}
		tags := tag1.List()
		this.Xsrf()
		this.Data["client"]=clients
		this.Data["tags"]=tags

		this.Data["pagetitle"]="修改用户信息页面"
		this.Layout="public/layout.html"
		this.TplName="client/update.html"
	}
}

//不允许删除，功能暂时不用
//func (this *ClientController)Delete()  {
//	id,err := this.GetInt("id")
//	if err!=nil{
//		this.Ctx.WriteString("数据有误")
//	}
//	client := &models.Client{Id:id}
//	if err := client.Delete();err !=nil{
//		this.Ctx.WriteString("删除失败")
//	}else {
//		this.Redirect(this.URLFor(".List"),302)
//	}
//
//}