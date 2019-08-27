/*此类自动生成,请勿修改*/
package template

/*竞技场常量模板配置*/
type ArenaConstantTemplateVO struct {

	//id
	Id int `json:"id"`

	//pk地图
	MapId int32 `json:"map_id"`

	//青龙地图
	MapId1 int32 `json:"map_id1"`

	//白虎地图
	MapId2 int32 `json:"map_id2"`

	//朱雀地图
	MapId3 int32 `json:"map_id3"`

	//玄武地图
	MapId4 int32 `json:"map_id4"`

	//位置1 x
	PosX1 float64 `json:"pos_x1"`

	//位置1 y
	PosY1 float64 `json:"pos_y1"`

	//位置1 z
	PosZ1 float64 `json:"pos_z1"`

	//位置2 x
	PosX2 float64 `json:"pos_x2"`

	//位置2 y
	PosY2 float64 `json:"pos_y2"`

	//位置2 z
	PosZ2 float64 `json:"pos_z2"`

	//复活次数
	RebornAmountMax int32 `json:"reborn_amount_max"`

	//复活系数
	RevivePar int32 `json:"revive_par"`

	//复活系数
	ReviveValue int32 `json:"revive_value"`

	//复活单id
	ReviveItem int32 `json:"revive_item"`

	//倒计时
	BattleTime int32 `json:"battle_time"`

	//获胜轮次
	TherionWinCount int32 `json:"therion_win_count"`

	//采集时间
	TreeGetTime int64 `json:"tree_get_time"`

	//掉落宝箱比例
	EquipBoxRate string `json:"equip_box_rate"`

	//掉落宝箱id
	EquipBoxId string `json:"equip_box_id"`

	//宝箱采集时间
	EquipBoxTime int64 `json:"equip_box_time"`

	//经验树id
	TreeId int32 `json:"tree_id"`

	//最大奖励次数
	WinCount int32 `json:"win_count"`

	//采集掉落宝箱id
	CaijiItem int32 `json:"caiji_item"`

	//采集掉落宝箱数量
	CaijiItemCount int32 `json:"caiji_item_count"`

	//四圣兽队伍数量
	TeamCount int32 `json:"team_count"`

	//经验
	TreeExp int32 `json:"tree_exp"`

	//经验点
	TreeExpPoint int32 `json:"tree_exp_point"`

	//宝箱预览
	PreBoxItem string `json:"pre_box_item"`

	//属性区间
	AttrMin int32 `json:"attr_min"`

	//属性区间
	AttrMax int32 `json:"attr_max"`

	//复活最小
	ReviveMin int32 `json:"revive_min"`

	//复活最大
	ReviveMax int32 `json:"revive_max"`

	//假人日志刷新时间下限(单位毫秒)
	RiZhiTimeMin int32 `json:"rizhi_time_min"`

	//假人日志刷新时间上限(单位毫秒)
	RiZhiTimeMax int32 `json:"rizhi_time_max"`

	//假人日志刷新时间上限(单位毫秒)
	RiZhiMax int32 `json:"rizhi_max"`

	//活动开始时间
	BeginTime string `json:"begin_time"`

	//活动结束时间
	EndTime string `json:"end_time"`

	//每天获取积分上限
	JiFenMaxDay int32 `json:"jifen_max_day"`

	//连胜称号id
	TitleId int32 `json:"title_id"`

	//排行榜人数
	RankCount int32 `json:"rank_count"`

	//第一奖励物品id
	RankFirstRew int32 `json:"rank_first"`
}
