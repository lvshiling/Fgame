package types

import (
	itemtypes "fgame/fgame/game/item/types"
)

type SynthesisType int32

const (
	HpGem             SynthesisType = iota + 1 //1生命宝石
	AttackGem                                  //2攻击宝石
	DefanceGem                                 //3防御宝石
	TuMoLing                                   //4屠魔令
	Secret                                     //5秘籍合成
	FashionChip                                //6时装碎片合成
	MountChip                                  //7坐骑碎片合成
	WingChip                                   //8战翼碎片合成
	WingResolve                                //9战翼相关拆解
	Jewelry                                    //10首饰合成
	HuFuChip                                   //11虎符合成
	AnqiResolve                                //12暗器相关拆解
	BodyShieldResolve                          //13护盾相关拆解
	ShenfaResolve                              //14身法相关拆解
	LingyuResolve                              //15领域相关拆解
	ShieldResolve                              //16盾刺相关拆解
	MountResolve                               //17坐骑相关拆解
	FeatherResolve                             //18仙羽相关拆解
	EquipmentBox                               //19装备箱合成
	SkillBox                                   //20技能箱合成
	YuanShenStone                              //21元神强化石
	DiHunCombine                               //22帝魂合成
	DiHunJueXing                               //23帝魂觉醒
	FaBaoResolve                               //24法宝相关拆解
	XianTiResolve                              //25仙体相关拆解
	FaBaoChip                                  //26法宝合成
	XianTiChip                                 //27仙体合成
)

func (t SynthesisType) Valid() bool {
	switch t {
	case HpGem,
		AttackGem,
		DefanceGem,
		TuMoLing,
		Secret,
		FashionChip,
		MountChip,
		WingChip,
		WingResolve,
		Jewelry,
		HuFuChip,
		AnqiResolve,
		BodyShieldResolve,
		ShenfaResolve,
		LingyuResolve,
		ShieldResolve,
		MountResolve,
		FeatherResolve,
		EquipmentBox,
		SkillBox,
		YuanShenStone,
		DiHunCombine,
		DiHunJueXing,
		FaBaoResolve,
		XianTiResolve,
		FaBaoChip,
		XianTiChip:
		return true
	default:
		return false
	}
}

var (
	luckyItemSubType = map[SynthesisType]itemtypes.ItemSubType{
		AttackGem:  itemtypes.ItemLuckySubTypeAttackLucky,
		HpGem:      itemtypes.ItemLuckySubTypeHpLucky,
		DefanceGem: itemtypes.ItemLuckySubTypeDefenceLucky,
	}
)

func (t SynthesisType) GetLuckyItemSubType() itemtypes.ItemSubType {
	return luckyItemSubType[t]
}
