/*此类自动生成,请勿修改*/
package template

/*神铸铸灵配置*/
type GodCastingCastingSpiritTemplateVO struct {

	//id
	Id int `json:"id"`

	//装备部位
	SubType int32 `json:"sub_type"`

	//铸灵类型
	ZhulingType int32 `json:"zhuling_type"`

	//需要的神铸等级
	NeedShenzhuLevel int32 `json:"need_shenzhu_level"`

	//铸灵消耗的物品ID
	UseItemId int32 `json:"use_item_id"`

	//关联的铸灵等级表起始ID
	ZhulingLevelBegin int32 `json:"zhuling_level_begin"`
}
