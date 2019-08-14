package models

//联系人/个人客户
type Concat struct {
	Id      int
	Name 	string //姓名
	Phone   string //电话
	Address string //地址
	Postid	string //邮编
	Remarks string //备注
	Tag	*Tag `orm:"null;rel(fk)"` //标签
	Createtime string //创建时间
	Updatetime string //更新时间
	Sales []*Sales `orm:"reverse(many)"` //客户订单

}

