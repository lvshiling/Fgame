/*此类自动生成,请勿修改*/
package template

/*buff动态配置*/
type BuffDongTaiTemplateVO struct {

	//id
	Id int `json:"id"`

	//后续id
	NextId int32 `json:"next_id"`

	//基础buff
	BuffId int32 `json:"buff_id"`

	//等级
	Lev int32 `json:"lev"`

	//buff持续时间
	TimeDuration int64 `json:"time_duration"`

	//改变移动速度万分比
	SpeedMovePercent int32 `json:"speed_move_percent"`

	//改变坚韧固定值
	ToughAdd int32 `json:"tough_add"`

	//改变坚韧万分比值
	ToughPercent int32 `json:"tough_percent"`
}
