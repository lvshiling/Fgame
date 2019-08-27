/*此类自动生成,请勿修改*/
package template

/*竞技场机器人模板配置*/
type ThreeRobatTemplateVO struct {

	//id
	Id int `json:"id"`

	//层数下限
	LevelMin int32 `json:"level_min"`

	//层数上限
	LevelMax int32 `json:"level_max"`

	//机器人下限
	RobotMin int32 `json:"robot_min"`

	//机器人上限
	RobotMax int32 `json:"robot_max"`

	//复活最小
	RebornMin int32 `json:"reborn_min"`

	//复活最大
	RebornMax int32 `json:"reborn_max"`

	//时间最小
	TimeMin int32 `json:"time_min"`

	//时间最大
	TimeMax int32 `json:"time_max"`
}
