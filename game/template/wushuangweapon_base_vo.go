/*此类自动生成,请勿修改*/
package template

/*无双神器基本配置*/
type WushuangWeaponBaseTemplateVO struct {

	//id
	Id int `json:"id"`

	//基础生命值
	Hp int64 `json:"hp"`

	//基础攻击力
	Attack int64 `json:"attack"`

	//基础防御力
	Defence int64 `json:"defence"`

	//可强化的最大等级
	StrengthenBeginId int32 `json:"strengthen_begin_id"`

	//外观激活需要的
	WaiguanJihuoLevel int32 `json:"waiguan_jihuo_level"`

	//外观类型
	WaiguanType int32 `json:"waiguan_type"`

	//外观Id
	WaiguanId int32 `json:"waiguan_id"`
}
