/*此类自动生成,请勿修改*/
package template

/*服务器生物生成配置*/
type BiologyPosTemplateVO struct {

	//id
	Id int `json:"id"`

	//怪物id
	BiologyId int32 `json:"biology_id"`

	//地图id
	MapId int32 `json:"map_id"`

	//下一id
	NextId int32 `json:"next_id"`

	//x坐标
	PosX float64 `json:"pos_x"`

	//y坐标
	PosY float64 `json:"pos_y"`

	//z坐标
	PosZ float64 `json:"pos_z"`
}
