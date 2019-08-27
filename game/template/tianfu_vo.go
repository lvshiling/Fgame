/*此类自动生成,请勿修改*/
package template

/*天赋配置*/
type TianFuTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//天赋名称
	Name string `json:"name"`

	//天赋升级起始id
	LevelBegin int32 `json:"level_begin"`

	//激活该天赋需要的前置id
	ParentId int32 `json:"parent_id"`
}
