/*此类自动生成,请勿修改*/
package template

/*灵童_灵兵培养配置*/
type LingTongWeaponPeiYangTemplateVO struct {

	//id
	Id int `json:"id"`

	//后续id
	NextId int32 `json:"next_id"`

	//等级
	Level int32 `json:"level"`

	//升级成功率
	UpdateWfb int32 `json:"update_wfb"`

	//升级所需物品
	UseItem int32 `json:"use_item"`

	//使用的物品数量
	ItemCount int32 `json:"item_count"`

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

	//灵童攻击力
	LingTongAttack int64 `json:"lingtong_attack"`

	//灵童独立暴击
	LingTongCritical int64 `json:"lingtong_critical"`

	//灵童独立命中值
	LingTongHit int64 `json:"lingtong_hit"`

	//灵童独立破格
	LingTongAbnormality int64 `json:"lingtong_abnormality"`
}
