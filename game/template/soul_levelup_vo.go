/*此类自动生成,请勿修改*/
package template

/*帝魂帝魂强化配置*/
type SoulLevelUpTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//帝魂类型
	Type int32 `json:"type"`

	//等级
	Level int32 `json:"level"`

	//强化成功率万分比
	UpdateWfb int32 `json:"update_wfb"`

	//每次需要消耗的银两
	UseSilver int32 `json:"use_silver"`

	//升级成功的最小次数
	TimesMin int32 `json:"times_min"`

	//升级成功的最大次数
	TimesMax int32 `json:"times_max"`

	//每次增加进度条的最小值
	AddMin int32 `json:"add_min"`

	//每次增加进度条的最大值
	AddMax int32 `json:"add_max"`

	//进度条总值
	ZhufuMax int32 `json:"zhufu_max"`

	//该等级加成的生命
	Hp int32 `json:"hp"`

	//该等级加成的攻击
	Attack int32 `json:"attack"`

	//该等级加成的防御
	Defence int32 `json:"defence"`
}
