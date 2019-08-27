/*此类自动生成,请勿修改*/
package template

/*婚戒培养配置*/
type MarryRingTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//戒指类型
	Type int32 `json:"type"`

	//戒指等级
	Level int32 `json:"level"`

	//升阶成功率
	UpdateWfb int32 `json:"update_wfb"`

	//升阶所需银两
	UseSilver int32 `json:"use_silver"`

	//升阶所需元宝
	UseGold int32 `json:"use_gold"`

	//升阶所需绑元
	UseBindGold int32 `json:"use_bindgold"`

	//升级所需物品
	UseItem int32 `json:"use_item"`

	//升阶所需物品数量
	ItemCount int32 `json:"item_count"`

	//最小次数
	TimesMin int32 `json:"times_min"`

	//最大次数
	TimesMax int32 `json:"times_max"`

	//每次培养增加的进度最小值
	AddMin int32 `json:"add_min"`

	//每次培养增加的进度最大值
	AddMax int32 `json:"add_max"`

	//属性加成
	Attr int32 `json:"attr"`

	//前端显示的进度值
	NeedRate int32 `json:"need_rate"`

	//模型ID
	ModelId int32 `json:"model_id"`
}
