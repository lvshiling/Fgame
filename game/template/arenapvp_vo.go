/*此类自动生成,请勿修改*/
package template

/*竞技场pvp配置*/
type ArenapvpTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一场id
	NextId int32 `json:"next_id"`

	//场次类型
	Type int32 `json:"changci"`

	//总人数
	PVPPlayerCount int32 `json:"player_all_count"`

	//场数
	PVPCount int32 `json:"huichang_count"`

	//地图id
	MapId int32 `json:"map_id"`

	//出生点1
	BirthPos1 string `json:"birth_pos1"`

	//出生点2
	BirthPos2 string `json:"birth_pos2"`

	//观众出生点
	BirthGuanzhong string `json:"birth_guanzhong"`

	//复活次数
	RebornCountMax int32 `json:"reborn_count_max"`

	//开始时间
	BeginTime string `json:"begin_time"`

	//结束时间
	EndTime string `json:"end_time"`

	//倒计时
	ZhanDouTime int64 `json:"zhandou_time"`

	//竞猜消耗绑元
	JingchaiUseBindgold int32 `json:"jingchai_use_bindgold"`

	//竞猜正确绑元
	JingcaiWinBindgold int32 `json:"jingcai_win_bindgold"`

	//竞猜正确奖励id
	JingcaiWinGetItemId string `json:"jingcai_win_get_item_id"`

	//竞猜正确奖励数量
	JingcaiWinGetItemCount string `json:"jingcai_win_get_item_count"`

	//竞猜错误奖励id
	JingcaiLoseGetItemId string `json:"jingcai_lose_get_item_id"`

	//竞猜错误奖励数量
	JingcaiLoseGetItemCount string `json:"jingcai_lose_get_item_count"`
}
