/*此类自动生成,请勿修改*/
package template

/*称号配置*/
type TitleTemplateVO struct {

	//id
	Id int `json:"id"`

	//称号名称
	Name string `json:"name"`

	//称号类型
	Type int32 `json:"type"`

	//称号子类型
	SubType int32 `json:"sub_type"`

	//排序
	Sort int32 `json:"sort"`

	//需要物品id
	NeedItemId int32 `json:"need_item_id"`

	//需要物品数量
	NeedItemCount int32 `json:"need_item_count"`

	//有效性
	Time int64 `json:"time"`

	//属性
	Attr int32 `json:"attr"`

	//称号升星起始id
	UpStarBeginId int32 `json:"title_upstar_begin_id"`

	//获取途径
	Access string `json:"access"`

	//资源
	Resource string `json:"resource"`
}
