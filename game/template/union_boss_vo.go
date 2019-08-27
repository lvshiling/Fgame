/*此类自动生成,请勿修改*/
package template

/*仙盟boss模板配置*/
type UnionBossTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//boss等级
	Level int32 `json:"level"`

	//升级所需经验
	Experience int32 `json:"experience"`

	//生物id
	BiologyId int32 `json:"biology_id"`

	//地图id
	MapId int32 `json:"map_id"`

	//奖励预览
	RewPre string `json:"rew_pre"`

	//盟主额外获得的银两奖励
	MzAwardSilver int32 `json:"mz_award_silver"`

	//盟主额外获得的元宝奖励
	MzAwardGold int32 `json:"mz_award_gold"`

	//盟主额外获得的绑元奖励
	MzAwardBindGold int32 `json:"mz_award_bindgold"`

	//盟主额外获得的物品奖励
	MzAwardItemId string `json:"mz_award_item_id"`

	//盟主额外获得的物品奖励数量
	MzAwardItemCount string `json:"mz_award_item_count"`

	//仙盟全体成员获得的银两数量
	CyAwardSilver int32 `json:"cy_award_silver"`

	//仙盟全体成员获得的元宝数量
	CyAwardGold int32 `json:"cy_award_gold"`

	//仙盟全体成员获得的绑元数量
	CyAwardBindGold int32 `json:"cy_award_bindgold"`

	//仙盟全体成员获得的物品奖励
	CyAwardItemId string `json:"cy_award_item_id"`

	//仙盟全体成员获得的物品奖励数量
	CyAwardItemCount string `json:"cy_award_item_count"`

	//坐标x
	PosX float64 `json:"pos_x"`

	//坐标y
	PosY float64 `json:"pos_y"`

	//坐标z
	PosZ float64 `json:"pos_z"`
}
