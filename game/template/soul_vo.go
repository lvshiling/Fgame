/*此类自动生成,请勿修改*/
package template

/*帝魂配置*/
type SoulTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//帝魂名称
	Name string `json:"name"`

	//帝魂品质
	SoulQuality int32 `json:"soul_quality"`

	//帝魂种类
	Type int32 `json:"type"`

	//帝魂类型(标签)
	SoulType int32 `json:"soul_type"`

	//帝魂等级
	Level int32 `json:"level"`

	//激活所需要的物品id
	NeedItemId string `json:"need_item_id"`

	//激活所需要的物品数量
	NeedItemCount string `json:"need_item_count"`

	//升级所需帝魂经验
	UplevelExp int32 `json:"uplevel_exp"`

	//可以吞噬的物品
	DevourId int32 `json:"devour_id"`

	//该等级附带的属性
	AttrId int32 `json:"attr_id"`

	//界面的标识显示
	Sign int32 `json:"sign"`

	//前置需激活的帝魂类型
	NeedSoul int32 `json:"need_soul"`

	//前置所需的帝魂等级
	NeedSoulLevel int32 `json:"need_soul_level"`
}
