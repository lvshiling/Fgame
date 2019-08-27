/*此类自动生成,请勿修改*/
package template

/*命格命宫配置*/
type MingGeMingGongTemplateVO struct {

	//id
	Id int `json:"id"`

	//命理类型
	Type int32 `json:"type"`

	//下一级id
	NextId int32 `json:"next_id"`

	//父命理id
	ParentId int32 `json:"parent_id"`

	//需要父命理达到的战力
	NeedParentZhanLi int32 `json:"need_parent_zhanli"`

	//需要玩家的等级
	NeedLevel int32 `json:"need_level"`

	//需要玩家的转数
	NeedZhuanShu int32 `json:"need_zhuanshu"`

	//起始id
	BeginId int32 `json:"begin_id"`
}
