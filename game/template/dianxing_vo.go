/*此类自动生成,请勿修改*/
package template

/*点星配置*/
type DianXingTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//等级
	Level int32 `json:"level"`

	//星谱类型
	XingPuType int32 `json:"type"`

	//星谱名字
	Name string `json:"name"`

	//点星成功率
	UpdateWfb int32 `json:"update_wfb"`

	//点星所需银两数量
	UseSilver int32 `json:"use_silver"`

	//点星所需星尘数量
	UseXingChen int32 `json:"use_xingchen"`

	//点星所需物品id
	UseItem int32 `json:"use_item"`

	//点星所需物品数量
	ItemCount int32 `json:"item_count"`

	//最小次数
	TimesMin int32 `json:"times_min"`

	//最大次数
	TimesMax int32 `json:"times_max"`

	//每次点星增加的进度最小值
	AddMin int32 `json:"add_min"`

	//每次点星增加的进度最大值
	AddMax int32 `json:"add_max"`

	//前端显示的进度值
	ZhufuMax int32 `json:"zhufu_max"`

	//该等级增加的生命
	Hp int32 `json:"hp"`

	//该等级增加的攻击
	Attack int32 `json:"attack"`

	//该等级增加的防御
	Defence int32 `json:"defence"`

	//过天是否清空祝福值 0不清 1清
	IsClear int32 `json:"is_clear"`

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
}
