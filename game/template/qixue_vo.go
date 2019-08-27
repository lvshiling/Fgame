/*此类自动生成,请勿修改*/
package template

/*泣血枪配置*/
type QiXueTemplateVO struct {

	//id
	Id int `json:"id"`

	//泣血枪的等级
	Level int32 `json:"type"`

	//下一个id
	NextId int32 `json:"next_id"`

	//不同阶别下对应的星数
	Star int32 `json:"star"`

	//升级需要的杀戮心
	UseResources int32 `json:"use_resources"`

	//获得兵魂ID
	WeaponId int32 `json:"weapon_id"`

	//死亡时掉落星级下限
	GasMin int32 `json:"resources_min"`

	//死亡时掉落星级上限
	GasMax int32 `json:"resources_max"`

	//掉落星级概率
	GasPercent int32 `json:"resources_percent"`

	//每个星星代表的杀戮心值
	StarCount int32 `json:"star_count"`

	//生命
	Hp int32 `json:"hp"`

	//攻击
	Attack int32 `json:"attack"`

	//防御
	Defence int32 `json:"defence"`
}
