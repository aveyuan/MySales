package models

//合同
//给某些客户定义合同金额，在销售的时候自动按合同金额执行
type Contract struct {
	Id int
	Titile string //合同标题
	Content string //合同内容
	Products  string //合同产品以及价格
	Sales *Sales `orm:"rel(one)"` //与销售单一对一关联
}
