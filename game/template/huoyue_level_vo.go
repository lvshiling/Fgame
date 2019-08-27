/*此类自动生成,请勿修改*/
package template

/*活跃度等级奖励配置*/
type HuoYueLevelTemplateVO struct {

	//id
	Id int `json:"id"`

	//关联任务表的ID
	QuestId int32 `json:"quest_id"`

	//最小等级
	LevelMin int32 `json:"level_min"`

	//最大等级
	LevelMax int32 `json:"level_max"`

	//完成任务后增加的活跃值
	HuoYue int32 `json:"huoyue"`

	//完成任务后增加的血炼值
	XueLian int32 `json:"xuelian"`
}
