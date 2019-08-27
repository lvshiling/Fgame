/*此类自动生成,请勿修改*/
package template

/*法宝配置*/
type FaBaoTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//名字
	Type int32 `json:"type"`

	//名字
	Name string `json:"name"`

	//阶数
	Number int32 `json:"number"`

	//顺序显示
	Pos int32 `json:"pos"`

	//升阶成功率
	UpdateWfb int32 `json:"update_wfb"`

	//升阶所需元宝
	UseMoney int32 `json:"use_money"`

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

	//前端显示的进度值
	ZhufuMax int32 `json:"zhufu_max"`

	//升星起始id
	FaBaoUpstarBeginId int32 `json:"fabao_upstar_begin_id"`

	//模型ID
	ModelId int32 `json:"model_id"`

	//所需银两
	UseYinliang int32 `json:"use_yinliang"`

	//食丹等级上限
	ShidanLimit int32 `json:"shidan_limit"`

	//法宝
	Icon int32 `json:"icon"`

	//幻化条件类型1
	MagicConditionType1 int32 `json:"magic_condition_type1"`

	//幻化条件类型参数1
	MagicConditionParameter1 string `json:"magic_condition_parameter1"`

	//幻化条件类型2
	MagicConditionType2 int32 `json:"magic_condition_type2"`

	//幻化条件类型参数2
	MagicConditionParameter2 string `json:"magic_condition_parameter2"`

	//幻化条件类型3
	MagicConditionType3 int32 `json:"magic_condition_type3"`

	//幻化条件类型参数3
	MagicConditionParameter3 string `json:"magic_condition_parameter3"`

	//ui
	UiScale string `json:"ui_scale"`

	//该等级增加的生命
	Hp int32 `json:"hp"`

	//该等级增加的攻击
	Attack int32 `json:"attack"`

	//该等级增加的防御
	Defence int32 `json:"defence"`

	//过天是否清空祝福值
	IsClear int32 `json:"is_clear"`
}
