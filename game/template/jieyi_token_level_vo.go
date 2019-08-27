/*此类自动生成,请勿修改*/
package template

/*结义道具等级配置*/
type JieYiTokenLevelTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//信物等级
	Level int32 `json:"level"`

	//所需物品id
	UseItemId int32 `json:"strengthen_item_id"`

	//所需物品数量
	UseItemCount int32 `json:"strengthen_item_count"`

	//升级概率
	Rate int32 `json:"strengthen_rate"`

	//最小次数
	TimesMin int32 `json:"times_min"`

	//最大次数
	TimesMax int32 `json:"times_max"`

	//增加的最小祝福值
	AddMin int32 `json:"add_min"`

	//增加的最大祝福值
	AddMax int32 `json:"add_max"`

	//前端显示的祝福值
	ZhufuMax int32 `json:"zhufu_max"`

	//该类型增加的生命
	Hp int64 `json:"hp"`

	//该类型增加的攻击
	Attack int64 `json:"attack"`

	//该类型增加的防御
	Defence int64 `json:"defence"`
}
