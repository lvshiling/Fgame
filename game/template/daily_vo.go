/*此类自动生成,请勿修改*/
package template

/*日环任务模板配置*/
type DailyTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//第X个任务
	Times int32 `json:"times_min"`

	//第X个任务
	TimesMax int32 `json:"times_max"`

	//最低等级
	LevelMin int32 `json:"level_min"`

	//最高等级
	LevelMax int32 `json:"level_max"`

	//关联任务表的ID,为该条数据的条件
	QuestId int32 `json:"quest_id"`

	//随机到该任务的几率
	Percent int32 `json:"percent"`

	//完成任务奖励的银两
	RewSilver int32 `json:"rew_silver"`

	//完成任务奖励的绑元
	RewBindGold int32 `json:"rew_bind_gold"`

	//完成任务奖励的元宝
	RewGold int32 `json:"rew_gold"`

	//奖励经验值
	RewExp int32 `json:"rew_xp"`

	//奖励经验点数
	RewExpPoint int32 `json:"rew_exp_point"`

	//奖励物品
	RewItemId string `json:"rew_item_id"`

	//奖励物品数量
	RewItemCount string `json:"rew_item_count"`

	//日环抽奖
	DropId int32 `json:"drop_id"`
}
