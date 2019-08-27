/*此类自动生成,请勿修改*/
package template

/*buff配置*/
type BuffTemplateVO struct {

	//id
	Id int `json:"id"`

	//名称
	Name string `json:"name"`

	//组
	Group int32 `json:"group"`

	//等级
	Lev int32 `json:"lev"`

	//buff类型
	BuffType int32 `json:"buff_type"`

	//效果类型
	TypeEffect int32 `json:"type_effect"`

	//打断类型
	TypeRemove int32 `json:"type_remove"`

	//限制类型
	TypeLimit int64 `json:"type_limit"`

	//触发类型
	TypeTouch int32 `json:"type_touch"`

	//免疫类型
	ImmuneType int32 `json:"immune_type"`

	//状态生效后给予的状态id
	SubId string `json:"sub_id"`

	//状态生效后给予的状态概率
	SubRate string `json:"sub_rate"`

	//永久
	IsForever int32 `json:"is_forever"`

	//状态持续时间
	TimeDuration int64 `json:"time_duration"`

	//第一次触发时间当buff非获得生效时使用
	FirstTouch int64 `json:"first_touch"`

	//状态触发频率
	Frequency int64 `json:"frequency"`

	//是否覆盖
	Replace int32 `json:"replace"`

	//状态叠加类型
	StackType int32 `json:"stack_type"`

	//状态叠加次数
	StackMax int32 `json:"stack_max"`

	//改变最大生命的固定值
	LifemaxAdd int32 `json:"lifemax_add"`

	//改变最大生命的上限万分比
	LifemaxPercent int32 `json:"lifemax_percent"`

	//改变最大体力的固定值
	TpmaxAdd int32 `json:"tpmax_add"`

	//改变最大体力的上限万分比
	TpmaxPercent int32 `json:"tpmax_percent"`

	//改变攻击力的固定值
	AttackAdd int32 `json:"attack_add"`

	//改变攻击力的万分比
	AttackPercent int32 `json:"attack_percent"`

	//改变防御的固定值
	DefenseAdd int32 `json:"defense_add"`

	//改变防御的万分比
	DefensePercent int32 `json:"defense_percent"`

	//改变暴击的固定值
	CriticalAdd int32 `json:"critical_add"`

	//改变暴击的万分比
	CriticalPercent int32 `json:"critical_percent"`

	//改变坚韧的固定值
	ToughAdd int32 `json:"tough_add"`

	//改变坚韧的万分比
	ToughPercent int32 `json:"tough_percent"`

	//改变格挡的固定值
	BlockAdd int32 `json:"block_add"`

	//改变格挡的万分比
	BlockPercent int32 `json:"block_percent"`

	//改变破格的固定值
	AbnormalityAdd int32 `json:"abnormality_add"`

	//改变破格的万分比
	AbnormalityPercent int32 `json:"abnormality_percent"`

	//改变命中的固定值
	HitAdd int32 `json:"hit_add"`

	//改变命中的万分比
	HitPercent int32 `json:"hit_percent"`

	//改变闪避的固定值
	DodgeAdd int32 `json:"dodge_add"`

	//改变闪避的万分比
	DodgePercent int32 `json:"dodge_percent"`

	//改变混元的固定值
	HunyuanAttAdd int32 `json:"hunyuan_att_add"`

	//改变混元的万分比
	HunyuanAttPercent int32 `json:"hunyuan_att_percent"`

	//改变混元防御的固定值
	HunyuanDefAdd int32 `json:"hunyuan_def_add"`

	//改变混元防御万分比
	HunyuanDefPercent int32 `json:"hunyuan_def_percent"`

	//改变生命的固定值
	LifeAdd int32 `json:"life_add"`

	//改变生命的万分比
	LifePercent int32 `json:"life_percent"`

	//改变生命的万分比
	LifePercentMax int32 `json:"life_percent_max"`

	//改变体力的固定值
	TpAdd int32 `json:"tp_add"`

	//改变体力的万分比
	TpPercent int32 `json:"tp_percent"`

	//改变技能cd的万分比
	SpellCdPercent int32 `json:"spell_cd_percent"`

	//反弹值
	FantanAdd int32 `json:"fantan_add"`

	//反弹万分比
	FantanPercent int32 `json:"fantan_percent"`

	//护盾值
	HudunAdd int32 `json:"hudun_add"`

	//护盾万分比
	HudunPercent int32 `json:"hudun_add"`

	//改变速度的万分比
	SpeedMovePercent int32 `json:"speed_move_percent"`

	//额外增加伤害的固定值
	HarmBase int32 `json:"harm_base"`

	//额外增加伤害的万分比
	HarmPercent int32 `json:"harm_percent"`

	//减少伤害的固定值
	CuthurtBase int32 `json:"cuthurt_base"`

	//减少伤害的万分比
	CuthurtPercent int32 `json:"cuthurt_percent"`

	//改变暴击几率的万分比
	CritRatePercent int32 `json:"crit_rate_percent"`

	//改变暴击伤害的万分比
	CritHarmPercent int32 `json:"crit_harm_percent"`

	//改变命中几率的万分比
	HitRatePercent int32 `json:"hit_rate_percent"`

	//改变破格的万分比
	BlockRatePercent int32 `json:"block_rate_percent"`

	//改变闪避的万分比
	DodgeRatePercent int32 `json:"dodge_rate_percent"`

	//下线保存类型
	OfflineSaveType int32 `json:"offline_save_type"`

	//特殊效果类型
	EffectType int32 `json:"effect_type"`

	//特殊效果值
	EffectTypeBase int32 `json:"effect_type_base"`

	//特殊效果万分比
	EffectTypePercent int32 `json:"effect_type_persent"`

	//广播类型
	ShowType int32 `json:"show_type"`

	//每次触发buff时获得经验值
	GetExp int32 `json:"get_exp"`

	//每次触发buff时获得经验点
	GetExpPoint int32 `json:"get_exp_point"`

	//变形
	ModelId int32 `json:"model_id"`

	//描述
	Description string `json:"description"`

	//简称
	Alias string `json:"alias"`

	//添加经验
	AddExp int32 `json:"add_exp"`

	//状态与抗性相对应的类型
	RelativeFastness int32 `json:"relative_fastness"`

	//buff飘字
	BuffPiaozi int32 `json:"buff_piaozi"`

	//子技能
	SkillId int32 `json:"skill_id"`

	//前置技能
	ParentBuffId string `json:"parent_buff_id"`
}
