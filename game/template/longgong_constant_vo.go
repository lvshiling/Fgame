/*此类自动生成,请勿修改*/
package template

/*龙宫探宝常量配置*/
type LongGongConstantTemplateVO struct {

	//id
	Id int `json:"id"`

	//boss生物id
	BossId int32 `json:"boss_id"`

	//召唤boss需要的珍珠采集数
	BossNeedCaiJiCount int32 `json:"boss_need_caiji_count"`

	//boss被杀后出现的采集点生物id
	BossBeKillCaiJiBiologyId int32 `json:"boss_bekill_caiji_biology_id"`

	//boss采集点地图x坐标
	PosX float64 `json:"pos_x"`

	//boss采集点地图y坐标
	PosY float64 `json:"pos_y"`

	//boss采集点地图z坐标
	PosZ float64 `json:"pos_z"`

	//boss采集点玩家最多采集次数
	BossBeKillCaiJiPersonalCount int32 `json:"boss_bekill_caiji_personal_count"`

	//珍珠采集虚假时间
	XuJiaCaiJiAddTime int32 `json:"xujia_caiji_add_time"`

	//每次珍珠采集虚假数据增加的数量
	XuJiaCaiJiAddCount int32 `json:"xujia_caiji_add_count"`

	//地图id
	MapId int32 `json:"map_id"`
}
