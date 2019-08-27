/*此类自动生成,请勿修改*/
package template

/*阵旗仙火模板配置*/
type ZhenFaXianHuoTemplateVO struct {

	//id
	Id int `json:"id"`

	//后续id
	NextId int32 `json:"next_id"`

	//阵法类型
	Type int32 `json:"type"`

	//等级
	Level int32 `json:"level"`

	//成功率
	UpdateWfb int32 `json:"update_wfb"`

	//消耗的物品id
	UseItem int32 `json:"use_item"`

	//消耗的物品数量
	ItemCount int32 `json:"item_count"`

	//最小次数
	TimesMin int32 `json:"times_min"`

	//最大次数
	TimesMax int32 `json:"times_max"`

	//每次增加祝福值最小值
	AddMin int32 `json:"add_min"`

	//每次增加祝福值最大值
	AddMax int32 `json:"add_max"`

	//前端祝福值显示值
	ZhuFuMax int32 `json:"zhufu_max"`

	//失败回退概率
	ReturnRate int32 `json:"return_rate"`

	//升星失败回退到多少级
	ReturnLevelId int32 `json:"return_level_id"`

	//生命上限
	Hp int64 `json:"hp"`

	//攻击值
	Attack int64 `json:"attack"`

	//防御值
	Defence int64 `json:"defence"`

	//该等级使用的保护符物品id
	BaoHuFuId int32 `json:"baohufu_id"`

	//消耗的保护符数量
	BaoHuFuCount int32 `json:"baohufu_count"`

	//本阵法增加的百分比
	Percent int32 `json:"percent"`
}
