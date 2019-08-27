/*此类自动生成,请勿修改*/
package template

/*天魔配置*/
type TianMoTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//天魔名称
	Name string `json:"name"`

	//阶数
	Number int32 `json:"number"`

	//激活类型
	ShengjieType int32 `json:"shengjie_type"`

	//激活条件
	ShengjieValue int32 `json:"shengjie_value1"`

	//升阶成功率
	UpdateWfb int32 `json:"update_wfb"`

	//升级所需元宝
	UseMoney int32 `json:"use_money"`

	//进阶所需物品id
	UseItem int32 `json:"use_item"`

	//进阶所需物品数量
	ItemCount int32 `json:"item_count"`

	//最小次数
	TimesMin int32 `json:"times_min"`

	//最大次数
	TimesMax int32 `json:"times_max"`

	//每次增加祝福值最小值
	AddMin int32 `json:"add_min"`

	//每次增加祝福值最大值
	AddMax int32 `json:"add_max"`

	//前端祝福值显示值
	ZhufuMax int32 `json:"zhufu_max"`

	//模型ID
	ModelId int32 `json:"model_id"`

	//进阶消耗银两
	UseYinliang int32 `json:"use_yinliangr"`

	//天魔体食丹等级上限
	ShidanLimit int32 `json:"shidan_limit"`

	//技能id
	SkillId int32 `json:"skill_id"`

	//外观类型
	WaiguanType int32 `json:"waiguan_type"`

	//外观id
	WaiguanValue int32 `json:"waiguan_value1"`

	//生命
	Hp int32 `json:"hp"`

	//攻击
	Attack int32 `json:"attack"`

	//防御
	Defence int32 `json:"defence"`

	//本阶进阶过天是否清空祝福值 0不清 1清
	IsClear int32 `json:"is_clear"`
}
