/*此类自动生成,请勿修改*/
package template

/*生物配置*/
type BiologyTemplateVO struct {

	//id
	Id int `json:"id"`

	//名字
	Name string `json:"name"`

	//所在地图
	Remarks string `json:"remarks"`

	//等级
	Level int `json:"level"`

	//境界
	Jingjie int32 `json:"jingjie"`

	//生物类型
	SetType int `json:"set_type"`

	//生物显示类型
	Type int `json:"type"`

	//脚本类型
	ScriptType int `json:"script_type"`

	//阵营,关联阵营表
	Faction int `json:"faction"`

	//posX
	ThreatType int `json:"threat_type"`

	//战力
	Force int `json:"force"`

	//血量
	Hp int `json:"hp"`

	//攻击
	Attack int `json:"attack"`

	//防御
	Defence int `json:"defence"`

	//暴击
	Critical int `json:"critical"`

	//坚韧
	Tough int `json:"tough"`

	//格挡
	Block int `json:"block"`

	//破格
	Abnormality int `json:"abnormality"`

	//命中
	Hit int `json:"hit"`

	//闪避
	Dodge int `json:"dodge"`

	//混元伤害
	HunyuanAtt int `json:"hunyuan_att"`

	//混元防御
	HunyuanDef int `json:"hunyuan_def"`

	//强制命中
	ForceHit int `json:"force_hit"`

	//强制暴击
	ForceCritical int `json:"force_critical"`

	//移动速度
	SpeedMove int `json:"speed_move"`

	//脱离战斗是否回血
	AutoRecoverMaxhealth int `json:"auto_recover_maxhealth"`

	//自动回血速度
	AutoRecoverTime int32 `json:"auto_recover_time"`

	//自动回血万分比
	AutoRecoverNum int32 `json:"auto_recover_num"`

	//固定伤害
	BaseDamge int `json:"base_damge"`

	//奖励归属类型
	DropOwnerType int `json:"drop_owner_type"`

	//固定经验值奖励
	ExpBase int `json:"exp_base"`

	//经验点数奖励
	ExpPoint int `json:"exp_point"`

	//掉落包组合
	DropCombine string `json:"drop_combine"`

	//掉落类型
	DropType int32 `json:"drop_type"`

	//掉落值
	DropFlag int32 `json:"drop_flag"`

	//巡逻范围
	Patrolradius int `json:"patrolradius"`

	//视野范围
	Alertradius int `json:"alertradius"`

	//追击范围
	Randradius int `json:"randradius"`

	//攻击技能id
	AttackId int `json:"attack_id"`

	//技能id1
	SkillId1 int `json:"skill_id_1"`

	//技能id1概率
	SkillRate1 int `json:"skill_rate1"`

	//技能id2
	SkillId2 int `json:"skill_id_2"`

	//技能id2概率
	SkillRate2 int `json:"skill_rate2"`

	//模型id
	ModleId int `json:"modle_id"`

	//对话
	Dialogue string `json:"dialogue"`

	//对话
	Speak string `json:"speak"`

	//冒泡持续时间
	DurationTime int `json:"duration_time"`

	//冒泡间隔时间
	IntervalTime int `json:"interval_time"`

	//称号id
	TitleId int `json:"title_id"`

	//半身像id
	HalfBodyPic int `json:"half_body_pic"`

	//对话内容
	Conversation string `json:"conversation"`

	//显示血量
	IsShowBlood int `json:"is_show_blood"`

	//头顶显血
	Hpbar int `json:"hp_bar"`

	//头顶血条数量
	LifeNumber int `json:"life_number"`

	//是否播放休闲动画
	IsShowAction int `json:"is_show_action"`

	//NPC语音
	TalkSound string `json:"talk_sound"`

	//重生类型
	RebornType int `json:"reborn_type"`

	//采集时间
	CaiJiTime int32 `json:"caiji_time"`

	//重生时间
	RebornTime string `json:"reborn_time"`

	//消失时间
	XiaoshiTime int32 `json:"xiaoshi_time"`

	//状态
	BuffIds string `json:"buff_ids"`

	//传送阵id
	PortalId int32 `json:"portal_id"`

	//活动组ID，用于活动掉落
	ActivityGroupId string `json:"activity_group_id"`

	//特殊效果免疫
	IsJitui int32 `json:"is_jitui"`

	//碰撞长度
	Pengzhuang float64 `json:"pengzhuang"`

	//碰撞宽度
	PengzhuangKuan float64 `json:"pengzhuang_kuan"`

	//所需疲劳值
	NeedPilao int32 `json:"need_pilao"`

	//浊气值
	ZhuoQi int32 `json:"zhuoqi"`

	//采集点刷新时间
	CaiJiRecoverTime int64 `json:"caiji_recover_time"`

	//采集点可采集次数
	CaiJiLimitCount int32 `json:"caiji_limit_count"`

	//采集点被采集到次数为0时，场景上的采集点模型是否消失
	CaiJiIsXiaoShi int32 `json:"caiji_is_xiaoshi"`

	//采集物被采集选的掉落id
	CaiJiChooseDrop string `json:"caiji_choose_drop"`

	//密藏比例
	BossMizangRate int32 `json:"boss_mizang_rate"`

	//密藏id
	MiZangId int32 `json:"mizang_id"`
}
