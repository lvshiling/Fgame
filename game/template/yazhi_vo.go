/*此类自动生成,请勿修改*/
package template

/*收益率模板配置*/
type YaZhiTemplateVO struct {

	//id
	Id int `json:"id"`

	//最低等级
	LevelMin int32 `json:"level_min"`

	//最高等级
	LevelMax int32 `json:"level_max"`

	//经验衰减万分比
	ExpPercent int32 `json:"exp_percent"`

	//掉落率衰落率
	DropPercent int32 `json:"drop_percent"`
}
