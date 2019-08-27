/*此类自动生成,请勿修改*/
package template

/*上古之灵升级配置*/
type ShangguzhilingJinjieTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一Id
	NextId int32 `json:"next_id"`

	//阶级
	Number int32 `json:"number"`

	//使用的物品id
	UseItem int32 `json:"use_item"`

	//使用的物品数量
	ItemCount int32 `json:"item_count"`

	//成功率万分比
	UpdateWfb int32 `json:"update_wfb"`

	//每次升阶最小次数
	TimesMin int32 `json:"times_min"`

	//每次升阶最大次数
	TimesMax int32 `json:"times_max"`

	//每次升阶每次随机增加祝福值最小数目
	AddMin int32 `json:"add_min"`

	//每次升阶每次随机增加祝福值最大数目
	AddMax int32 `json:"add_max"`

	//祝福值最大
	ZhufuMax int32 `json:"zhufu_max"`

	//hp
	Hp int32 `json:"hp"`

	//攻击力
	Attack int32 `json:"attack"`

	//防御力
	Defence int32 `json:"defence"`

	//百分比
	Percent int32 `json:"percent"`
}
