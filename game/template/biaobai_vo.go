/*此类自动生成,请勿修改*/
package template

/*表白培养配置*/
type MarryDevelopTemplateVO struct {

	//id
	Id int `json:"id"`

	//表白系统等级
	Level int32 `json:"level"`

	//下一级id
	NextId int32 `json:"next_id"`

	//经验值
	Experience int32 `json:"experience"`

	//配偶提供百分比
	Percent int32 `json:"percent"`

	//该等级增加的生命
	AddHp int32 `json:"hp"`

	//该等级增加的攻击
	AddAttack int32 `json:"attack"`

	//该等级增加的防御
	AddDefence int32 `json:"defence"`
}
