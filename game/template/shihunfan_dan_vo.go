/*此类自动生成,请勿修改*/
package template

/*噬魂幡培养丹配置*/
type ShiHunFanDanTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//等级
	Level int32 `json:"level"`

	//升阶成功率
	UpdateWfb int32 `json:"update_wfb"`

	//升级所需物品
	UseItem int32 `json:"use_item"`

	//升阶所需物品数量
	ItemCount int32 `json:"item_count"`

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
