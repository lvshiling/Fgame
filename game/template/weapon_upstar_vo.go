/*此类自动生成,请勿修改*/
package template

/*兵魂升星配置*/
type WeaponUpstarTemplateVO struct {

	//id
	Id int `json:"id"`

	//下个id
	NextId int32 `json:"next_id"`

	//星级
	Level int32 `json:"level"`

	//名字
	Name string `json:"name"`

	//升星成功率
	UpstarRate int32 `json:"upstar_rate"`

	//升星物品ID(开天)
	UpstarItemId int32 `json:"upstar_item_id"`

	//升星物品数量(开天)
	UpstarItemCount int32 `json:"upstar_item_count"`

	//升星物品ID(奕剑)
	UpstarItemId2 int32 `json:"upstar_item_id2"`

	//升星物品数量(奕剑)
	UpstarItemCount2 int32 `json:"upstar_item_count2"`

	//升星物品ID(破月)
	UpstarItemId3 int32 `json:"upstar_item_id3"`

	//升星物品数量(破月)
	UpstarItemCount3 int32 `json:"upstar_item_count3"`

	//最小次数
	TimesMin int32 `json:"times_min"`

	//最大次数
	TimesMax int32 `json:"times_max"`

	//每次培养增加的进度最小值
	AddMin int32 `json:"add_min"`

	//每次培养增加的进度最大值
	AddMax int32 `json:"add_max"`

	//前端显示的进度值
	ZhufuMax int32 `json:"zhufu_max"`

	//该等级增加的生命
	Hp int32 `json:"hp"`

	//该等级增加的攻击
	Attack int32 `json:"attack"`

	//该等级增加的防御
	Defence int32 `json:"defence"`
}
