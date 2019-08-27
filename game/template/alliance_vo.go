/*此类自动生成,请勿修改*/
package template

/*仙盟模板配置*/
type AllianceTemplateVO struct {

	//id
	Id int `json:"id"`

	//仙盟版本
	UnionType int32 `json:"union_type"`

	//仙盟等级
	UnionLevel int32 `json:"union_level"`

	//升级所需建设度
	UnionBuild int64 `json:"union_build"`

	//下一级id
	NextLevelId int32 `json:"next_level_id"`

	//升级所需的前置等级
	UnionParentLevel int32 `json:"union_parent_level"`

	//联盟最高人数
	UnionMax int32 `json:"union_max"`

	//副盟主数量
	UnionPost2 int32 `json:"union_post2"`

	//堂主数量
	UnionPost3 int32 `json:"union_post3"`

	//精英数量
	UnionPost4 int32 `json:"union_post4"`

	//仙盟仓库数量
	UnionStorage int32 `json:"union_storage"`
}
