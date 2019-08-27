/*此类自动生成,请勿修改*/
package template

/*戮仙刃配置*/
type MassacreTemplateVO struct {

	//id
	Id int `json:"id"`

	//戮仙刃阶别
	Type int32 `json:"type"`

	//下一个id
	NextId int32 `json:"next_id"`

	//不同阶别下对应的星数
	Star int32 `json:"star"`

	//升星成功率
	UpdatePercent int32 `json:"update_percent"`

	//升阶所需银两
	UseMoney int32 `json:"use_money"`

	//升阶所需杀气
	UseGas int32 `json:"use_gas"`

	//进阶所需物品id
	UseItem int32 `json:"use_item"`

	//进阶所需物品数量
	ItemCount int32 `json:"item_count"`

	//最小次数
	TimesMin int32 `json:"times_min"`

	//最大次数
	TimesMax int32 `json:"times_max"`

	//获得兵魂ID
	WeaponId int32 `json:"weapon_id"`

	//死亡时掉落星级下限
	GasMin int32 `json:"gas_min"`

	//死亡时掉落星级上限
	GasMax int32 `json:"gas_max"`

	//掉落星级概率
	GasPercent int32 `json:"gas_percent"`

	//每个星星代表的杀气值
	StarCount int32 `json:"star_count"`

	//生命
	Hp int32 `json:"hp"`

	//攻击
	Attack int32 `json:"attack"`

	//防御
	Defence int32 `json:"defence"`
}
