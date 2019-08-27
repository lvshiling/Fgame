/*此类自动生成,请勿修改*/
package template

/*附加系统觉醒丹阶别配置*/
type SystemAwakeTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//系统类型
	SysType int32 `json:"type"`

	//系统阶数
	Number int32 `json:"number"`

	//该等级增加的生命
	Hp int32 `json:"hp"`

	//该等级增加的攻击
	Attack int32 `json:"attack"`

	//该等级增加的防御
	Defence int32 `json:"defence"`

	//升级表起始id
	BeginId int `json:"level_begin_id"`
}
