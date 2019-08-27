package types

type DamageType int32

const (
	//自动回血
	DamageTypeAutoRecovery DamageType = iota
	//普通攻击
	DamageTypeAttack
	//恢复
	DamageTypeRecovery
	//暴击
	DamageTypeCrit
	//反弹伤害
	DamageTypeFISH
	//普通攻击格挡
	DamageTypeAttackGeDang
	//暴击格挡
	DamageTypeCritGeDang
	//闪避
	DamageTypeDodge
	//混元伤害
	DamageTypeHunYuan
	//护盾
	DamageTypeHuDun
	//灼烧
	DamageTypeZhuoShao
	//免疫
	DamageTypeMianYi
	//无效
	DamageTypeWuXiao
	//压制
	DamageTypeYaZhi
	//无法压制
	DamageTypeWuFaYaZhi
	//中毒
	DamageTypeZhongDu
	//帝魂
	DamageTypeSoul
	//血池
	DamageTypeXueChi
	//幻兽
	DamageTypePet
	//灵童
	DamageTypeLingTong
	//吸血
	DamageTypeXiXue
	//预留1
	DamageTypeYuLiu1
	//预留2
	DamageTypeYuLiu2
	//预留3
	DamageTypeYuLiu3
	//预留4
	DamageTypeYuLiu4
	//预留5
	DamageTypeYuLiu5
)

type MoveType int32

const (
	MoveTypeNormal MoveType = 1 + iota
	MoveTypeHit
	MoveTypeATK
)

var (
	moveTypeMap = map[MoveType]string{
		MoveTypeNormal: "普通",
		MoveTypeHit:    "击退",
		MoveTypeATK:    "攻击",
	}
)

func (mt MoveType) String() string {
	return moveTypeMap[mt]
}

type OwnerType int32

const (
	//默认没有主人
	OwnerTypeNone OwnerType = iota
	//玩家
	OwnerTypePlayer
	//仙盟
	OwnerTypeAlliance
	//夫妻
	OwnerTypeMarry
	//阵营
	OwnerTypeCamp
)
