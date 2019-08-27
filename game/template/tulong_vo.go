/*此类自动生成,请勿修改*/
package template

/*屠龙配置*/
type TuLongTemplateVO struct {

	//id
	Id int `json:"id"`

	//怪物类型 0为大BOSS 1为小BOSS
	Type int32 `json:"type"`

	//生物id
	BiologyId int32 `json:"biology_id"`

	//标识,用于服务端区分不同的怪物
	BiaoShi int32 `json:"biaoshi"`

	//仙盟成员获得的银两
	RewSilver int32 `json:"rew_silver"`

	//仙盟成员获得的元宝
	RewGold int32 `json:"rew_gold"`

	//仙盟成员获得的绑元
	RewBindGold int32 `json:"rew_bindgold"`

	//仙盟成员获得的经验
	RewExp int32 `json:"rew_exp"`

	//仙盟成员获得的物品id
	RewItemId string `json:"rew_item_id"`

	//仙盟成员获得的物品数量
	RewCount string `json:"rew_count"`
}
