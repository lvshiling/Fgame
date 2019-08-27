/*此类自动生成,请勿修改*/
package template

/*屠龙常量配置*/
type TuLongConstantTemplateVO struct {

	//id
	Id int `json:"id"`

	//奖励预览
	rew_preview_id string `json:"rew_preview_id"`

	//奖励预览的数量
	RewPreviewCount string `json:"rew_preview_count"`

	//大龙蛋id
	BigEgg int32 `json:"big_egg"`

	//小龙蛋id
	SmallEgg int64 `json:"small_egg"`

	//地图id
	MapId int32 `json:"map_id"`

	//大BOSS刷新时间(活动开始后多少时间,单位:毫秒)
	BossTime int32 `json:"boss_time"`

	//采集小龙蛋的时间(毫秒)
	CaiJiTime int32 `json:"caiji_time"`

	//小龙蛋的刷新时间(毫秒)
	SmallEggShuaXin int32 `json:"small_egg_shuaxin"`

	//场景人数限制
	PlayerLimitCount int32 `json:"player_limit_count"`
}
