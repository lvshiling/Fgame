/*此类自动生成,请勿修改*/
package template

/*神魔常量配置*/
type ShenMoConstantTemplateVO struct {

	//id
	Id int `json:"id"`

	//物资生物id
	WuZiBiologyId int32 `json:"wuzi_biology_id"`

	//大旗生物id
	DaQiBiologyId int32 `json:"daqi_biology_id"`

	//物资采集获得的功勋
	WuZiGetGongXun int32 `json:"wuzi_get_gongxun"`

	//物资采集获得的仙盟积分
	WuZiGeiJiFen int32 `json:"wuzi_gei_jifen"`

	//采集大旗获得的功勋
	DaQiGetGongXun int32 `json:"daqi_get_gongxun"`

	//采集大旗获得的仙盟积分
	DaQiGetJiFen int32 `json:"daqi_get_jifen"`

	//活动人数上限
	PlayerLimitCount int32 `json:"player_limit_count"`

	//同一玩家的击杀cd(毫秒)
	KillCd int32 `json:"kill_cd"`

	//击杀同一玩家多次后进入该玩家的cd
	KillCountCd int32 `json:"kill_count_cd"`

	//复活后的buffid
	RebornBuff int32 `json:"reborn_buff"`

	//地图id
	MapId int32 `json:"map_id"`
}
