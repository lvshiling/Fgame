/*此类自动生成,请勿修改*/
package template

/*创世城常量模板配置*/
type ChuangShiCityConstantTemplateVO struct {

	//id
	Id int `json:"id"`

	//地图id
	MapId int32 `json:"map_id"`

	//出生地点1
	BirthPos1 string `json:"birth_pos1"`

	//出生地点2
	BirthPos2 string `json:"birth_pos2"`

	//非城战期间进入
	IsJinRu int32 `json:"is_jinru"`

	//进入时间
	JinRuTime string `json:"jinru_time"`

	//人数限制
	JinruPlayerMax int32 `json:"jinru_player_max"`
}
