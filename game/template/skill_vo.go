/*此类自动生成,请勿修改*/
package template

/*技能配置*/
type SkillTemplateVO struct {

	//id
	Id int `json:"id"`

	//名字
	Name string `json:"name"`

	//类型1
	FirstType int32 `json:"type_1st"`

	//类型2
	SecondType int32 `json:"type_2nd"`

	//类型3
	ThirdType int32 `json:"type_3rd"`

	//类型4
	FourthType int32 `json:"type_4rd"`

	//当type_2nd为3，5时，走以下2进制规则：1，被玩家攻击生效,2，被怪物攻击生效
	BeTriggered int32 `json:"be_triggered"`

	//当type_2nd为5时，此字段代表被敌人攻击时血量降低的触发点（万分比）
	HpTrigger int32 `json:"hp_trigger"`

	//表现buff
	BuffId int `json:"buff_id"`

	//受击表现buff
	BuffId2 int `json:"buff_id_2"`

	//给自己上buff
	BuffId3 int32 `json:"buff_id_3"`

	//给自己上buff概率
	BuffId3Rate int32 `json:"buff_id3_rate"`

	//
	TypeId int32 `json:"type_id"`

	//等级
	Lev int32 `json:"lev"`

	//触发几率
	Rate int32 `json:"rate"`

	//cd时间
	CdTime int32 `json:"cd_time"`

	//cd组
	CdGroup int32 `json:"cd_group"`

	//技能选择目标（二进制）1.自身2.友方 4.敌方
	TargetSelect int32 `json:"target_select"`

	//作用目标（二进制）1.玩家 2.怪物（NPC）
	TargetAction int32 `json:"target_action"`

	//是否无视安全区
	RuleBreak int32 `json:"rule_break"`

	//朝向
	FaceNeed int32 `json:"face_need"`

	//施法距离
	Distance int32 `json:"distance"`

	//作用范围类型
	AreaType int32 `json:"area_type"`

	//作用范围半径
	AreaRadius int32 `json:"area_radius"`

	//作用范围
	AreaRange int32 `json:"area_range"`

	//目标数量
	TargetCount int32 `json:"target_count"`

	//职业需要
	ProNeed int32 `json:"pro_need"`

	//等级需要
	LevNeed int32 `json:"lev_need"`

	//消耗血量
	CostHpValue int32 `json:"cost_hp_value"`

	//消耗血量万分比
	CostHpPersent int32 `json:"cost_hp_persent"`

	//消耗体力
	CostTpValue int32 `json:"cost_tp_value"`

	//基础伤害
	DamageValueBase int32 `json:"damage_value_base"`

	//伤害值
	DamageValue int32 `json:"damage_value"`

	//技能威力值
	SpellDamage int32 `json:"spell_damage"`

	//技能伤害
	SpellPower int32 `json:"spell_power"`

	//血量上线万分比
	DamagePersent int32 `json:"damage_persent"`

	//治疗值
	CureValue int32 `json:"cure_value"`

	//治疗万分比
	CurePersent int32 `json:"cure_persent"`

	//仇恨值
	HatredValue int32 `json:"hatred_value"`

	//仇恨万分比
	HatredPersent int32 `json:"hatred_persent"`

	//添加状态
	AddStatus string `json:"add_status"`

	//添加状态概率
	AddStatusRate string `json:"add_status_rate"`

	//添加属性
	AddAttrId int32 `json:"add_attr_id"`

	//仙盟属性光环
	UnionAttrId int32 `json:"union_attr_id"`

	//添加命中
	AddHit int32 `json:"add_hit"`

	//添加暴击
	AddCritical int32 `json:"add_critical"`

	//特殊效果对象
	SpecialTarget int32 `json:"special_target"`

	//特殊效果
	SpecialEffect int32 `json:"special_effect"`

	//特殊效果比例
	SpecialEffectRate int32 `json:"special_effect_rate"`

	//特殊效果值
	SpecialEffectValue int32 `json:"special_effect_value"`

	//特殊效果值2
	SpecialEffectValue2 int32 `json:"special_effect_value2"`

	//特殊效果值3
	SpecialEffectValue3 int32 `json:"special_effect_value3"`

	//战力添加
	AddPower int32 `json:"add_power"`

	//动态技能升级id
	SpellUpgradeBeginId int32 `json:"spell_upgrade_begin_id"`

	//天赋起始id
	TianFuBeginId int32 `json:"tianfu_begin_id"`

	//消耗物品
	ConsumeGoods int32 `json:"consume_goods"`

	//消耗物品数量
	ConsumeGoodsCount int32 `json:"consume_goods_count"`

	//消耗经验
	ConsumeExperience int32 `json:"consume_experience"`

	//消耗银两
	ConsumeMoney int32 `json:"consume_money"`

	//限制被动
	LimitTouch int32 `json:"limit_touch"`

	//男技能特效
	BoyShowId int32 `json:"boy_show_id"`

	//女技能特效
	GirlShowId int32 `json:"girl_show_id"`

	//男技能ico
	BoySpellIco int32 `json:"boy_spell_ico"`

	//女技能ico
	GirlSpellIco int32 `json:"girl_spell_ico"`

	//技能描述
	SpellDescription string `json:"spell_description"`

	//星盘界面仙魂描述
	SpellGetdescription string `json:"spell_getdescription"`

	//仙魂特效描述
	SpecialDes string `json:"special_des"`

	//仙魂每级的效果描述
	UplevelDes1 string `json:"uplevel_des1"`

	//仙魂每级的效果持续时间描述
	UplevelDes2 string `json:"uplevel_des2"`

	//技能获得描述
	SpellUiDescription string `json:"spell_ui_description"`

	//品质
	Quality int32 `json:"quality"`

	//多段
	MultiCount int32 `json:"multi_count"`

	//展示名字
	ShowName int32 `json:"show_name"`

	//自动拖入
	AutoSlot int32 `json:"auto_slot"`

	//挂机
	GuaJi int32 `json:"guaji"`

	//下一个技能
	SpellNext int32 `json:"spell_next"`

	//施法时间
	ActionTime int32 `json:"action_time"`

	//延迟时间1
	DelayTimes1 string `json:"delay_times_1"`

	//延迟时间2
	DelayTimes2 string `json:"delay_times_2"`

	//服务器延迟时间
	DelayTimesServer string `json:"delay_times_server"`

	//受到伤害最高上限
	BeHarmLimitHp int32 `json:"be_harm_limit_hp"`

	//复活buff
	RebornBuff int32 `json:"reborn_buff"`

	//复活技能
	RebornSkill int32 `json:"reborn_skill"`

	//复活技能概率
	RebornSkillRate int32 `json:"reborn_skill_rate"`
}
