/*此类自动生成,请勿修改*/
package template

/*天机牌配置*/
type TianJiPaiTemplateVO struct {

	//id
	Id int `json:"id"`

	//任务类型
	PoolType int32 `json:"pool_type"`

	//最小等级
	LevelMin int32 `json:"level_min"`

	//最大等级
	LevelMax int32 `json:"level_max"`

	//关联到功能开启表
	ModuleOpenedId int32 `json:"module_opened_id"`

	//关联任务表的ID
	QuestId int32 `json:"quest_id"`

	//最小星级
	StarMin int32 `json:"star_min"`

	//最大星级
	StarMax int32 `json:"star_max"`

	//完成任务奖励的银两
	RewSilver int32 `json:"rew_silver"`

	//间隔多少次走掉落包1
	DropCount1 int32 `json:"drop_count1"`

	//掉落包1
	DropId1 int32 `json:"drop_id1"`

	//间隔多少次走掉落包2
	DropCount2 int32 `json:"drop_count2"`

	//掉落包2
	DropId2 int32 `json:"drop_id2"`

	//间隔多少次走掉落包3
	DropCount3 int32 `json:"drop_count3"`

	//掉落包3
	DropId3 int32 `json:"drop_id3"`

	//间隔多少次走掉落包4
	DropCount4 int32 `json:"drop_count4"`

	//掉落包4
	DropId4 int32 `json:"drop_id4"`

	//前五次掉落
	SpeDrop string `json:"spe_drop"`

	//备注
	Remarks string `json:"remarks"`
}
