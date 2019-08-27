/*此类自动生成,请勿修改*/
package template

/*神铸铸灵升级配置*/
type GodCastingCastingSpiritLevelTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一id
	NextId int32 `json:"next_id"`

	//等级
	Level int32 `json:"level"`

	//hp
	Hp int64 `json:"hp"`

	//攻击力
	Attack int64 `json:"attack"`

	//防御力
	Defence int64 `json:"defence"`

	//铸灵消耗的物品数量
	UseItemCount int32 `json:"use_item_count"`

	//神铸成功率万分比
	UpdateWfb int32 `json:"update_wfb"`

	//铸灵每次升阶最小次数
	TimesMin int32 `json:"times_min"`

	//铸灵每次升阶最大次数
	TimesMax int32 `json:"times_max"`

	//铸灵每次升阶每次随机增加祝福值最小数目
	AddMin int32 `json:"add_min"`

	//铸灵每次升阶每次随机增加祝福值最大数目
	AddMax int32 `json:"add_max"`

	//祝福值最大
	ZhufuMax int32 `json:"zhufu_max"`
}
