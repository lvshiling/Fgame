/*此类自动生成,请勿修改*/
package template

/*灵童升星配置*/
type LingTongUpstarTemplateVO struct {

	//id
	Id int `json:"id"`

	//后续id
	NextId int32 `json:"next_id"`

	//等级
	Level int32 `json:"level"`

	//名称
	Name string `json:"name"`

	//升级成功率
	UpdateWfb int32 `json:"upstar_rate"`

	//升级所需物品
	UseItem int32 `json:"upstar_item_id"`

	//使用的物品数量
	ItemCount int32 `json:"upstar_item_count"`

	//最小次数
	TimesMin int32 `json:"times_min"`

	//最大次数
	TimesMax int32 `json:"times_max"`

	//每次随机加的最小祝福
	AddMin int32 `json:"add_min"`

	//每次随机加的最大祝福
	AddMax int32 `json:"add_max"`

	//前端显示的最大祝福值
	ZhufuMax int32 `json:"zhufu_max"`

	//该等级增加的生命
	Hp int64 `json:"hp"`

	//该等级增加的攻击
	Attack int64 `json:"attack"`

	//该等级增加的防御
	Defence int64 `json:"defence"`
}
