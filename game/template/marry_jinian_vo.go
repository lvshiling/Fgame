/*此类自动生成,请勿修改*/
package template

/*结婚纪念*/
type MarryJiNianTemplateVO struct {

	//id
	Id int `json:"id"`

	//类型：1低级，2中级，3高级
	Type int32 `json:"type"`

	//对应婚礼到达次数
	NeedNum int32 `json:"need_num"`

	//到达次数后赠送的物品
	ZhuheItem string `json:"zhuhe_item"`

	//到达次数后赠送的物品数量
	ZhuheItemCount string `json:"zhuhe_item_count"`

	//对应称号Id
	TitleId int32 `json:"title_id"`
}
