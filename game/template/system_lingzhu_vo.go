/*此类自动生成,请勿修改*/
package template

/*附加灵珠配置*/
type SystemLingZhuTemplateVO struct {

	//id
	Id int `json:"id"`

	//灵童等级
	LingtongId int32 `json:"lingtong_id"`

	//灵珠类型
	Type int32 `json:"type"`

	//升级消耗物品类型
	UseItemId int32 `json:"use_item_id"`

	//关联升级表起始Id
	LevelBegin int32 `json:"level_begin"`
}
