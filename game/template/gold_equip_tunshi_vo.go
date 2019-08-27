/*此类自动生成,请勿修改*/
package template

/*元神等级配置*/
type GoldYuanTemplateVO struct {

	//id
	Id int `json:"id"`

	//元神等级
	Level int32 `json:"level"`

	//下级id
	NextId int32 `json:"next_id"`

	//经验
	Exp int32 `json:"exp"`

	//生命
	AddHp int32 `json:"hp"`

	//攻击
	AddAttack int32 `json:"attack"`

	//防御
	AddDefect int32 `json:"defence"`
}
