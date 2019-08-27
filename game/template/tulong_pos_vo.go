/*此类自动生成,请勿修改*/
package template

/*屠龙出生位置配置*/
type TuLongPosTemplateVO struct {

	//id
	Id int `json:"id"`

	//地图id
	MapId int32 `json:"map_id"`

	//坐标点类型 0为BOSS出生 1为玩家出生点
	Type int32 `json:"type"`

	//用于区分刷新出大BOSS的点 使刷新出大BOSS的点不会有玩家出生
	BiaoShi int32 `json:"biaoshi"`

	//坐标x
	PosX float64 `json:"pos_x"`

	//坐标y
	PosY float64 `json:"pos_y"`

	//坐标z
	PosZ float64 `json:"pos_z"`
}
