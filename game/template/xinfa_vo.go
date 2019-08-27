/*此类自动生成,请勿修改*/
package template

/*心法配置*/
type XinFaTemplateVO struct {

	//id
	Id int `json:"id"`

	//后续id
	NextId int32 `json:"next_id"`

	//名字
	Name string `json:"name"`

	//类型
	Type int32 `json:"type"`

	//等级
	Level int32 `json:"level"`

	//描述
	Des string `json:"des"`

	//扩展描述
	DesInfo string `json:"des_info"`

	//升级所需银两
	NeedYinLiang int32 `json:"need_yinliang"`

	//激活所需物品ID
	NeedItemId int32 `json:"need_item_id"`

	//激活所需物品数量
	NeedItemNum int32 `json:"need_item_num"`

	//技能id
	SkillId int32 `json:"skill_id"`
}
