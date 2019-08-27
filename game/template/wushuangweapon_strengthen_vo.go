/*此类自动生成,请勿修改*/
package template

/*无双神器强化配置*/
type WushuangWeaponStrengthenTemplateVO struct {

	//id
	Id int `json:"id"`

	//强化等级
	Level int32 `json:"level"`

	//下一级强化等级
	NextId int32 `json:"next_id"`

	//强化生命值
	Hp int64 `json:"hp"`

	//强化攻击力
	Attack int64 `json:"attack"`

	//强化防御力
	Defence int64 `json:"defence"`

	//升级所需经验
	Experience int64 `json:"experience"`

	//突破需要的最低转数
	TupoZhuanshu int32 `json:"tupo_zhuanshu"`

	//突破成功率
	TupoRate int32 `json:"tupo_rate"`

	//是否需要突破
	IsTupo int32 `json:"is_tupo"`

	//突破需要物品数量
	TupoNeedCount int32 `json:"tupo_need_count"`

	//突破需要物品品质
	TupoNeedQuality int32 `json:"tupo_need_quality"`
}
