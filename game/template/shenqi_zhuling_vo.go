/*此类自动生成,请勿修改*/
package template

/*神器注灵配置*/
type ShenQiZhuLingTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//神器类型
	ShenQiType int32 `json:"type"`

	//器灵类型
	QiLingTyp int32 `json:"qiling_type"`

	//器灵部位
	QiLingSubTyp int32 `json:"qiling_sub_type"`

	//注灵等级
	Level int32 `json:"level"`

	//升级成功率
	UpdateWfb int32 `json:"update_wfb"`

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

	//生命
	Hp int32 `json:"hp"`

	//攻击
	Attack int32 `json:"attack"`

	//防御
	Defence int32 `json:"defence"`

	//对应器灵属性增加百分比
	Percent int32 `json:"percent"`

	//消耗灵气值
	NeedZhuLing int32 `json:"need_zhuling"`
}
