/*此类自动生成,请勿修改*/
package template

/*附加使用觉醒丹配置*/
type SystemAwakeLevelTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//觉醒等级
	Level int32 `json:"level"`

	//觉醒成功率
	UpdateWfb int32 `json:"update_wfb"`

	//觉醒使用银两
	UseSilver int64 `json:"use_silver"`

	//觉醒使用物品数量
	UseItemCount int32 `json:"item_count"`

	//觉醒最小次数
	TimesMin int32 `json:"times_min"`

	//觉醒最大次数
	TimesMax int32 `json:"times_max"`

	//觉醒最小祝福
	AddMin int32 `json:"add_min"`

	//觉醒最大祝福
	AddMax int32 `json:"add_max"`

	//最大祝福值
	ZhufuMax int32 `json:"zhufu_max"`

	//该等级增加的生命
	Hp int32 `json:"hp"`

	//该等级增加的攻击
	Attack int32 `json:"attack"`

	//该等级增加的防御
	Defence int32 `json:"defence"`
}
