/*此类自动生成,请勿修改*/
package template

/*创世城池建设等级模板配置*/
type ChuangShiChengFangJianSheTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一级id
	NextId int32 `json:"next_id"`

	//类型
	Type int32 `json:"type"`

	//等级
	Level int32 `json:"level"`

	//升级经验
	NeedExp int32 `json:"need_exp"`

	//生命
	Hp int32 `json:"hp"`

	//攻击
	Attack int32 `json:"attack"`

	//防御
	Defence int32 `json:"defence"`

	//天气关联id
	TianqiId int32 `json:"tianqi_id"`

	//掉落下限
	FallMin int32 `json:"fall_min"`

	//掉落上限
	FallMax int32 `json:"fall_max"`

	//掉落概率
	FallPercent int32 `json:"fall_percent"`
}
