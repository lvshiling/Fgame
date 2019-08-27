/*此类自动生成,请勿修改*/
package template

/*灵童_灵身通灵配置*/
type LingTongShenFaTongLingTemplateVO struct {

	//id
	Id int `json:"id"`

	//后续id
	NextId int32 `json:"next_id"`

	//等级
	Level int32 `json:"level"`

	//升级成功率
	UpdateWfb int32 `json:"update_wfb"`

	//升级所需物品
	UseItem int32 `json:"use_item"`

	//使用的物品数量
	ItemCount int32 `json:"item_count"`

	//最小次数
	TimesMin int32 `json:"times_min"`

	//最大次数
	TimesMax int32 `json:"times_max"`

	//每次随机加的最小祝福
	AddMin int32 `json:"add_min"`

	//每次随机加的最大祝福
	AddMax int32 `json:"add_max"`

	//前端显示的最大祝福值
	ZhufuMax int32 `json:"zhufu_max"`

	//灵身基础全属性万分比
	LingTongShenFaPercent int32 `json:"add_percent"`
}
