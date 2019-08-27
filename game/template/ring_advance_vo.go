/*此类自动生成,请勿修改*/
package template

/*特戒进阶配置*/
type RingAdvanceTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//阶数
	Advance int32 `json:"number"`

	//升阶成功率
	UpdateWfb int32 `json:"update_wfb"`

	//升阶所需物品id
	UseItem int32 `json:"use_item"`

	//升阶所需物品数量
	ItemCount int32 `json:"item_count"`

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

	//生命万分比
	HpPercent int32 `json:"hp_percent"`

	//攻击万分比
	AttackPercent int32 `json:"attack_percent"`

	//防御万分比
	DefPercent int64 `json:"def_percent"`

	//升阶后祝福值 0不清 1清
	IsClear int32 `json:"is_clear"`
}
