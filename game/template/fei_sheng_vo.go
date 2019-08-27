/*此类自动生成,请勿修改*/
package template

/*飞升模板配置*/
type FeiShengTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一级di
	NextId int `json:"next_id"`

	//飞升等级
	Level int32 `json:"level"`

	//所需功德
	GongDe int32 `json:"gongde"`

	//飞升概率
	Rate int32 `json:"rate"`

	//提升概率物品
	ItemId int32 `json:"item_id"`

	//提升概率物品数量
	ItemCount int32 `json:"item_count"`

	//增加的概率
	AddRate int32 `json:"add_rate"`

	//生命
	Hp int64 `json:"hp"`

	//攻击
	Attack int64 `json:"attack"`

	//防御
	Defence int64 `json:"defence"`

	//可配置的潜能
	QnAdd int32 `json:"qn_add"`

	//洗点所需元宝
	XidianGold int32 `json:"xidian_gold"`

	//功德的转换比例
	GongdeRatio int32 `json:"gongde_ratio"`

	//赠送经验的转换比例
	GiveExpRatio int32 `json:"give_exp_ratio"`
}
