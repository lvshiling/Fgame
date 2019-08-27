/*此类自动生成,请勿修改*/
package template

/*仙体皮肤升星配置*/
type XianTiUpstarTemplateVO struct {

	//id
	Id int `json:"id"`

	//后续id
	NextId int32 `json:"next_id"`

	//星级
	Level int32 `json:"level"`

	//名字
	Name string `json:"name"`

	//升星需要消耗的物品id
	UpstarItemId int32 `json:"upstar_item_id"`

	//升星消耗需要的物品数量
	UpstarItemCount int32 `json:"upstar_item_count"`

	//升星成功几率（万分比）
	UpstarRate int32 `json:"upstar_rate"`

	//最小次数
	TimesMin int32 `json:"times_min"`

	//最大次数
	TimesMax int32 `json:"times_max"`

	//每次随机加的最小祝福
	AddMin int32 `json:"add_min"`

	//每次随机加的最大祝福
	AddMax int32 `json:"add_max"`

	//每次随机加的最大祝福
	ZhufuMax int32 `json:"zhufu_max"`

	//生命加成（固定值）
	Hp int32 `json:"hp"`

	//攻击加成（固定）
	Attack int32 `json:"attack"`

	//防御加成（固定值）
	Defence int32 `json:"defence"`

	//仙体基础全属性万分比
	XianTiPercent int32 `json:"xianti_percent"`
}
