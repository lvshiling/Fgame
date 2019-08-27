/*此类自动生成,请勿修改*/
package template

/*活动怪物掉落模板配置*/
type TeShuDropTemplateVO struct {

	//id
	Id int `json:"id"`

	//活动id
	GroupId int32 `json:"group_id"`

	//最小区间
	MinCount int32 `json:"min_count"`

	//最大区间
	MaxCount int32 `json:"max_count"`

	//掉落id
	DropId int32 `json:"drop_id"`
}
