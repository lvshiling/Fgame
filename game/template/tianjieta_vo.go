/*此类自动生成,请勿修改*/
package template

/*天劫塔配置*/
type TianJieTaTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//地图id
	MapId int32 `json:"map_id"`

	//等级
	Level int32 `json:"level"`

	//boss怪id
	BossId int32 `json:"boss_id"`

	//推荐战力
	SuggestedFightPower int32 `json:"suggested_fight_power"`

	//通关经验
	RewExp int32 `json:"rew_exp"`

	//通关经验点
	RewExpPoint int32 `json:"rew_exp_point"`

	//通关银两
	RewSilver int32 `json:"rew_silver"`

	//奖励元宝数量
	RewGold int32 `json:"rew_gold"`

	//绑元数量
	RewBindGold int32 `json:"rew_bind_gold"`

	//道具id
	RewItem int32 `json:"rew_item"`

	//道具数量
	RewCount int32 `json:"rew_count"`

	//夫妻助战获得亲密度
	RewQinMiDu int32 `json:"rew_qinmidu"`

	//属性加成
	Attr int32 `json:"attr"`

	//技能
	SkillId int32 `json:"skill_id"`

	//物品奖励id预览
	RewardPreviewId string `json:"reward_preview_id"`

	//物品奖励数量预览
	RewardPreviewNum string `json:"reward_preview_num"`
}
