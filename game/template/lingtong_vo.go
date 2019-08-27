/*此类自动生成,请勿修改*/
package template

/*灵童配置*/
type LingTongTemplateVO struct {

	//id
	Id int `json:"id"`

	//灵童名称
	Name string `json:"name"`

	//激活类型
	JiHuoType int32 `json:"jihuo_type"`

	//激活灵童使用的物品id
	UseItemId int32 `json:"use_item_id"`

	//激活灵童使用的物品数量
	UseItemCount int32 `json:"use_item_count"`

	//灵童独立攻击力
	LingTongAttack int64 `json:"lingtong_attack"`

	//灵童独立暴击
	LingTongCritical int64 `json:"lingtong_critical"`

	//灵童独立命中值
	LingTongHit int64 `json:"lingtong_hit"`

	//灵童独立破格
	LingTongAbnormality int64 `json:"lingtong_abnormality"`

	//玩家角色攻击力继承到灵童身上万分比
	PlayerAttackPercent int64 `json:"player_attack_percent"`

	//玩家增加的生命上限
	Hp int64 `json:"hp"`

	//玩家增加的攻击值
	Attack int64 `json:"attack"`

	//玩家增加的防御值
	Defence int64 `json:"defence"`

	//灵童升级起始id
	LingTongShengJiId int32 `json:"lingtong_shengji_id"`

	//灵童培养起始id
	LingTongPeiYangId int32 `json:"lingtong_peiyang_id"`

	//灵童改名消耗物品
	NameItemId int32 `json:"name_item_id"`

	//灵童改名消耗物品数量
	NameItemCount int32 `json:"name_item_count"`

	//灵童初始时装关联时装表ID
	LingTongFashionId int32 `json:"lingtong_fashion_id"`

	//灵童初始武器关联灵童武器表id
	LingTongWeapon int32 `json:"lingtong_weapon"`

	//普通攻击的id,前端控制
	AttackId int32 `json:"attack_id"`

	//被动技能,服务器触发
	SkillId1 int32 `json:"skill_id_1"`

	//升星起始id
	LingTongUpstarBeginId int32 `json:"lingtong_upstar_begin_id"`

	//灵珠开启等级
	LingzhuOpenLevel int32 `json:"lingzhu_open_level"`
}
