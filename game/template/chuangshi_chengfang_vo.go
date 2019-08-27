/*此类自动生成,请勿修改*/
package template

/*创世城池建设模板配置*/
type ChuangShiChengFangTemplateVO struct {

	//id
	Id int `json:"id"`

	//建筑名字
	Name string `json:"name"`

	//类型
	Type int32 `json:"type"`

	//升级关联id
	LevelBeginId int32 `json:"level_begin_id"`

	//升级物品id
	LevelItemId int32 `json:"level_item_id"`
}
