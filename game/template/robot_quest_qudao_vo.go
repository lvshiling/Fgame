/*此类自动生成,请勿修改*/
package template

/*机器人渠道任务配置*/
type RobotQuestQudaoTemplateVO struct {

	//id
	Id int `json:"id"`

	//id
	QudaoId int `json:"qudao_id"`

	//地图id
	MapId int32 `json:"map_id"`

	//玩家数量
	PlayerLimitCount int32 `json:"player_limit_count"`

	//任务起始id
	QuestBeginId int32 `json:"quest_begin_id"`

	//任务结束id
	QuestOverId int32 `json:"quest_over_id"`

	//血量最小
	HpMin int64 `json:"hp_min"`

	//血量最大
	HpMax int64 `json:"hp_max"`

	//攻击最小
	AttackMin int64 `json:"attack_min"`

	//攻击最大
	AttackMax int64 `json:"attack_max"`

	//防御最小
	DefenceMin int64 `json:"defence_min"`

	//防御最大
	DefenceMax int64 `json:"defence_max"`

	//传送阵
	ChuansongzhenId int32 `json:"chuansongzhen_id"`

	//刷新时间
	RefreshTime int32 `json:"refresh_time"`
}
