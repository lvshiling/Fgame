/*此类自动生成,请勿修改*/
package template

/*结义道具配置*/
type JieYiTokenTemplateVO struct {

	//id
	Id int `json:"id"`

	//信物类型
	Type int32 `json:"type"`

	//所需物品id
	UseItemId int32 `json:"use_item_id"`

	//所需物品数量
	UseItemCount int32 `json:"use_item_count"`

	//信物强化等级起始id
	BeginId int `json:"xinwu_strengthen_begin_id"`

	//信物分享的万分比属性
	SharePercent int32 `json:"share_percent"`
}
