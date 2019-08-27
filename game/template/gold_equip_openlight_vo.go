/*此类自动生成,请勿修改*/
package template

/*元神金装开光配置*/
type GoldEquipOpenLightTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//开光次数
	Times int32 `json:"times"`

	//成功率
	SuccessRate int32 `json:"success_rate"`

	//使用物品id
	UseItem string `json:"use_item"`

	//使用物品数量
	UseCount string `json:"use_count"`

	//最小次数
	TimesMin int32 `json:"times_min"`

	//最大次数
	TimesMax int32 `json:"times_max"`

	//属性加成
	AttrPercent int32 `json:"attr_percent"`

	//返还物品Id
	MeltingReturnId string `json:"melting_return_id"`

	//返还物品数量
	MeltingReturnCount string `json:"melting_return_count"`
}
