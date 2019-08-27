/*此类自动生成,请勿修改*/
package template

/*创世守卫配置*/
type ChuangShiWarGuardTemplateVO struct {

	//id
	Id int `json:"id"`

	//名称
	Name string `json:"name"`

	//生物id
	BiologyId int32 `json:"biology_id"`

	//地图id
	Map int32 `json:"map"`

	//守卫id
	SceneId int32 `json:"scene_id"`

	//默认召唤方式
	GuardSummon int32 `json:"guard_summon"`

	//需要银两
	NeedSilver int32 `json:"need_silver"`

	//需要元宝
	NeedGold int32 `json:"need_gold"`

	//需要绑元
	NeedBindGold int32 `json:"need_bind_gold"`

	//需要物品id
	NeedItemId string `json:"need_item_id"`

	//需要物品数量
	NeedItemCount string `json:"need_item_count"`
}
