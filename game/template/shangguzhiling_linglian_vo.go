/*此类自动生成,请勿修改*/
package template

/*上古之灵灵炼配置*/
type ShangguzhilingLinglianTemplateVO struct {

	//id
	Id int `json:"id"`

	//灵兽类型
	Type int32 `json:"type"`

	//部位类型
	SubType int32 `json:"sub_type"`

	//需要上古之灵的等级
	NeedSgzlLevel int32 `json:"need_sgzl_level"`

	//部位起始属性随机池
	ChushiAttrPoolBeginId int32 `json:"chushi_attr_pool_begin_id"`

	//部位属性随机池
	AttrPoolBeginId int32 `json:"attr_pool_begin_id"`
}
