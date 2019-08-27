/*此类自动生成,请勿修改*/
package template

/*四神遗迹特殊怪配置*/
type FourGodSpecialTemplateVO struct {

	//id
	Id int `json:"id"`

	//类型 1为默认出生点 2为终点
	Type int32 `json:"type"`

	//地图id
	MapId int32 `json:"map_id"`

	//地图x坐标
	PosX float64 `json:"pos_x"`

	//地图y坐标
	PosY float64 `json:"pos_y"`

	//地图z坐标
	PosZ float64 `json:"pos_z"`
}
