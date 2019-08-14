package models

import "github.com/astaxie/beego/orm"

//合同
//给某些客户定义合同金额，在销售的时候自动按合同金额执行
type Contract struct {
	Id int
	Titile string //合同标题
	Content string //合同内容
	Products  string //合同产品以及价格
	Sales *Sales `orm:"rel(one)"` //与销售单一对一关联
}

//添加联系人
func (this *Contract)Add()error  {
	o := orm.NewOrm()
	if _,err := o.Insert(this);err!=nil{
		return err
	}
	return nil
}

//联系人列表
func (this *Contract)List()[]*Contract  {
	var contract []*Contract
	o := orm.NewOrm()
	o.QueryTable(Product{}).All(&contract)
	return products
}

//通过id查询联系人
func (this *Contract)IdProduct()*Contract {
	o := orm.NewOrm()
	o.Read(this)
	return this
}