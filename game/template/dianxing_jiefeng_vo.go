/*此类自动生成,请勿修改*/
package template

/*点星解封配置*/
type DianXingJieFengTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//等级
	Level int32 `json:"level"`

	//升阶成功率
	UpdateWfb int32 `json:"update_wfb"`

	//升级所需银两数量
	UseSilver int32 `json:"use_silver"`

	//升级所需物品id
	UseItem int32 `json:"use_item"`

	//升级所需物品数量
	ItemCount int32 `json:"item_count"`

	//最小次数
	TimesMin int32 `json:"times_min"`

	//最大次数
	TimesMax int32 `json:"times_max"`

	//每次解封增加的进度最小值
	AddMin int32 `json:"add_min"`

	//每次解封增加的进度最大值
	AddMax int32 `json:"add_max"`

	//前端显示的进度值
	ZhufuMax int32 `json:"zhufu_max"`

	//该等级增加的属性万分比
	AttrPercent int32 `json:"percent"`

	//该等级增加的生命
	Hp int32 `json:"hp"`

	//该等级增加的攻击
	Attack int32 `json:"attack"`

	//该等级增加的防御
	Defence int32 `json:"defence"`

	//坐骑属性万分比
	MountPercent int32 `json:"mount_percent"`

	//战翼属性万分比
	WingPercent int32 `json:"wing_percent"`

	//领域属性万分比
	FieldPercent int32 `json:"field_percent"`

	//身法属性万分比
	ShenFaPercent int32 `json:"shenfa_percent"`

	//护体盾属性万分比
	BodyShieldPercent int32 `json:"body_shield_percent"`

	//暗器属性万分比
	AnQiPercent int32 `json:"anqi_percent"`

	//法宝属性万分比
	FaBaoPercent int32 `json:"fabao_percent"`

	//仙体属性万分比
	XianTiPercent int32 `json:"xianti_percent"`

	//天魔体属性万分比
	TianMoTiPercent int32 `json:"tianmoti_percent"`

	//噬魂幡属性万分比
	ShiHunFanPercent int32 `json:"shihunfan_percent"`

	//灵骑属性万分比
	LingTongMountPercent int32 `json:"lingtong_zuoqi_percent"`

	//灵兵属性万分比
	LingTongWeaponPercent int32 `json:"lingtong_binghun_percent"`

	//灵翼属性万分比
	LingTongWingPercent int32 `json:"lingtong_zhanyi_percent"`

	//灵宝属性万分比
	LingTongFaBaoPercent int32 `json:"lingtong_fabao_percent"`

	//灵体属性万分比
	LingTongXianTiPercent int32 `json:"lingtong_xianti_percent"`

	//灵域属性万分比
	LingTongLingYuPercent int32 `json:"lingtong_lingyu_percent"`

	//灵身属性万分比
	LingTongShenFaPercent int32 `json:"lingtong_shenfa_percent"`
}
