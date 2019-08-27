/*此类自动生成,请勿修改*/
package template

/*阵法激活模板配置*/
type ZhenFaJiHuoTemplateVO struct {

	//id
	Id int `json:"id"`

	//阵法类型
	Type int `json:"type"`

	//阵法总等级
	Name string `json:"name"`

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

	//需要的前置战法类型
	NeedZhenFaType int32 `json:"need_zhenfa_type"`

	//需要的前置战法等级
	NeedZhenFaLevel int32 `json:"need_zhenfa_level"`

	//生命上限
	Hp int64 `json:"hp"`

	//攻击值
	Attack int64 `json:"attack"`

	//防御值
	Defence int64 `json:"defence"`
}
