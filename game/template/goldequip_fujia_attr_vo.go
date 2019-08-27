/*此类自动生成,请勿修改*/
package template

/*元神金装附加属性池配置*/
type GoldEquipFuJiaAttrTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int `json:"next_id"`

	//权重
	Rate int32 `json:"rate"`

	//生命
	Hp int32 `json:"hp"`

	//攻击
	Attack int32 `json:"attack"`

	//防御
	Defence int32 `json:"defence"`

	//加成该装备基础属性百分比
	EquipPercent int32 `json:"equip_percent"`

	//每15秒回复生命的值
	Huixie int32 `json:"15s_huixie"`

	//对BOSS类型（策划类型）增伤
	BossZengshang int32 `json:"boss_zengshang"`

	//对BOSS类型（策划类型）减伤
	BossJianshang int32 `json:"boss_jianshang"`

	//击杀怪物后增加的buffId
	KillMonsterBuff int32 `json:"kill_monster_buff"`
}
