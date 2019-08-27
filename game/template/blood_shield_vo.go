/*此类自动生成,请勿修改*/
package template

/*血盾配置*/
type BloodShieldTemplateVO struct {

	//id
	Id int `json:"id"`

	//血盾阶别
	Type int32 `json:"type"`

	//血盾名称
	Name string `json:"name"`

	//下一个id
	NextId int32 `json:"next_id"`

	//不同阶别下对应的星数
	Star int32 `json:"star"`

	//升星成功率
	UpdatePercent int32 `json:"update_percent"`

	//进级所需银两
	UseMoney int32 `json:"use_money"`

	//升级所需血炼值
	UseBlood int32 `json:"use_blood"`

	//进阶所需物品id
	UseItem string `json:"use_item"`

	//进阶所需物品数量
	ItemCount string `json:"item_count"`

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

	//升级时获得技能ID
	SpellId int32 `json:"spell_id"`

	//食用血盾碎片上限
	MedicinalLimit int32 `json:"medicinal_limit"`

	//系统公告
	Note int32 `json:"note"`
}
