/*此类自动生成,请勿修改*/
package template

/*领域配置*/
type LingyuTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//领域名称
	Name string `json:"name"`

	//身法阶数
	Number int32 `json:"number"`

	//领域类型
	Type int32 `json:"type"`

	//进阶成功率
	UpdateWfb int32 `json:"update_wfb"`

	//进级所需元宝
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

	//升星开始id
	FieldUpstarBeginId int32 `json:"field_upstar_begin_id"`

	//激活技能
	Skill int32 `json:"skill"`

	//模型ID
	ModelId int32 `json:"model_id"`

	//进阶消耗银两
	UseYinliang int32 `json:"use_yinliangr"`

	//使用幻化丹id
	ShidanId1 int32 `json:"shidan_id_1"`

	//使用幻化丹数量
	ShidanLimit int32 `json:"shidan_limit"`

	//领域图标
	Icon int32 `json:"icon"`

	//幻化条件类型1
	MagicConditionType1 int32 `json:"magic_condition_type1"`

	//根据类型,写入对应参数1
	MagicConditionParameter1 string `json:"magic_condition_parameter1"`

	//幻化条件类型2
	MagicConditionType2 int32 `json:"magic_condition_type2"`

	//根据类型,写入对应参数2
	MagicConditionParameter2 string `json:"magic_condition_parameter2"`

	//幻化条件类型3
	MagicConditionType3 int32 `json:"magic_condition_type3"`

	//根据类型,写入对应参数3
	MagicConditionParameter3 string `json:"magic_condition_parameter3"`

	//本阶进阶过天是否清空祝福值 0不清 1清
	IsClear int32 `json:"is_clear"`

	//激活类型
	ShengjieType int32 `json:"shengjie_type"`

	//激活条件
	ShengjieValue int32 `json:"shengjie_value1"`

	//hp
	Hp int32 `json:"hp"`

	//该等级增加的攻击
	Attack int32 `json:"attack"`

	//该等级增加的防御
	Defence int32 `json:"defence"`
}
