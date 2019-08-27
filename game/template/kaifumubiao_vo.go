/*此类自动生成,请勿修改*/
package template

/*开服目标模板配置*/
type KaiFuMuBiaoTemplateVO struct {

	//id
	Id int `json:"id"`

	//标签
	Name string `json:"name"`

	//下一个id
	NextId int32 `json:"next_id"`

	//开服第几天解锁
	KaiFuTime int32 `json:"kaifu_time"`

	//任务组起始id
	GroupBeginId int32 `json:"group_begin_id"`

	//银两奖励
	RewardSilver int64 `json:"reward_silver"`

	//完成任务组目标奖励的物品
	ItemId string `json:"item_id"`

	//完成任务组目标奖励的物品数量
	ItemCount string `json:"item_count"`

	//完成该任务组内多少个任务才能够领取任务组奖励
	FinishQuestCount int32 `json:"finish_quest_count"`
}
