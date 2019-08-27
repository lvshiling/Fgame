/*此类自动生成,请勿修改*/
package template

/*阵法阵旗模板配置*/
type ZhenFaZhenQiTemplateVO struct {

	//id
	Id int `json:"id"`

	//后续id
	NextId int32 `json:"next_id"`

	//阵法类型
	Type int32 `json:"type"`

	//阵旗类型
	SubType int32 `json:"sub_type"`

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

	//生命上限
	Hp int64 `json:"hp"`

	//攻击值
	Attack int64 `json:"attack"`

	//防御值
	Defence int64 `json:"defence"`

	//需要的阵法等级
	NeedLevel int32 `json:"need_level"`
}
