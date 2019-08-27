/*此类自动生成,请勿修改*/
package template

/*打宝塔模板配置*/
type TowerTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//最低等级
	LevelMin int32 `json:"level_min"`

	//最高等级
	LevelMax int32 `json:"level_max"`

	//地图id
	MapId int32 `json:"map_id"`

	//配置0非boss层
	BossId int32 `json:"boss_id"`

	//Boss刷新x
	BossPosX float64 `json:"boss_pos_x"`

	//Boss刷新y
	BossPosY float64 `json:"boss_pos_y"`

	//Boss刷新z
	BossPosZ float64 `json:"boss_pos_z"`

	//需要的物品id
	ZhifeiItem int32 `json:"zhifei_item"`

	//需要的物品数量
	ZhifeiItemCount int32 `json:"zhifei_item_count"`

	//奖励预览
	RewardPreviewId string `json:"reward_preview_id"`

	//小怪id
	XujiaId string `json:"xujia_id"`

	//小怪掉落id
	XujiaItem string `json:"xujia_item"`
}
