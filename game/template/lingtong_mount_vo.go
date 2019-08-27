/*此类自动生成,请勿修改*/
package template

/*灵骑配置*/
type LingTongMountTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//坐骑名称
	Name string `json:"name"`

	//幻化类型
	Type int32 `json:"type"`

	//阶数
	Number int32 `json:"number"`

	//升阶成功率
	UpdateWfb int32 `json:"update_wfb"`

	//升级所需元宝
	UseMoney int64 `json:"use_money"`

	//进阶所需物品id
	UseItem int32 `json:"use_item"`

	//进阶所需物品数量
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
	ZhufuMax int32 `json:"zhufu_max"`

	//坐骑属性
	Attr int32 `json:"attr"`

	//坐骑升星起始id
	LingTongMountUpstarBeginId int32 `json:"upstar_begin_id"`

	//模型ID
	ModelId int32 `json:"model_id"`

	//进阶消耗银两
	UseYinliang int64 `json:"use_yinliangr"`

	//幻化丹食丹等级上限
	ShidanLimit int32 `json:"shidan_limit"`

	//培养丹食丹等级上限
	CulturingDanLimit int32 `json:"culturing_dan_limit"`

	//灵骑头像图标
	Icon int32 `json:"icon"`

	//幻化条件类型1
	MagicConditionType1 int32 `json:"magic_condition_type1"`

	//根据类型,写入对应参数1
	MagicConditionParameter1 string `json:"magic_condition_parameter1"`

	//幻化条件类型2
	MagicConditionType2 int32 `json:"magic_condition_type2"`

	//根据类型,写入对应参数2
	MagicConditionParameter2 string `json:"magic_condition_parameter2"`

	//幻化条件类型3
	MagicConditionType3 int32 `json:"magic_condition_type3"`

	//根据类型,写入对应参数3
	MagicConditionParameter3 string `json:"magic_condition_parameter3"`

	//本阶进阶过天是否清空祝福值 0不清 1清
	IsClear int32 `json:"is_clear"`

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

	//客户端需要特殊处理,坐骑id
	MountId int32 `json:"mount_id"`
}
