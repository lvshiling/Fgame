/*此类自动生成,请勿修改*/
package template

/*衣橱培养配置*/
type YiChuPeiYangTemplateVO struct {

	//id
	Id int `json:"id"`

	//后续id
	NextId int32 `json:"next_id"`

	//食丹等级
	Level int32 `json:"level"`

	//名字
	Name string `json:"name"`

	//使用物品
	UseItem int32 `json:"use_item"`

	//物品数量
	ItemCount int32 `json:"item_count"`

	//成功几率（万分比）
	UpstarWfb int32 `json:"upstar_wfb"`

	//最小次数
	TimesMin int32 `json:"times_min"`

	//最大次数
	TimesMax int32 `json:"times_max"`

	//每次随机加的最小祝福
	AddMin int32 `json:"add_min"`

	//每次随机加的最大祝福
	AddMax int32 `json:"add_max"`

	//每次随机加的最大祝福
	ZhufuMax int32 `json:"zhufu_max"`

	//生命加成（固定值）
	Hp int64 `json:"hp"`

	//攻击加成（固定）
	Attack int64 `json:"attack"`

	//防御加成（固定值）
	Defence int64 `json:"defence"`

	//技能id
	SkillId int32 `json:"skill_id"`

	//技能id2
	SkillId2 int32 `json:"skill_id2"`

	//套装属性万分比
	Percent int32 `json:"percent"`
}
