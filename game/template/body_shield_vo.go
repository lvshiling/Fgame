/*此类自动生成,请勿修改*/
package template

/*护体盾配置*/
type BodyShieldTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//护体盾名称
	Name string `json:"name"`

	//阶数
	Number int32 `json:"number"`

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

	//护体盾属性
	Attr int32 `json:"attr"`

	//模型ID
	ModelId int32 `json:"model_id"`

	//进阶消耗银两
	UseYinliang int32 `json:"use_yinliangr"`

	//金甲丹食丹等级上限
	ShidanLimit int32 `json:"shidan_limit"`

	//本阶进阶过天是否清空祝福值 0不清 1清
	IsClear int32 `json:"is_clear"`
}
