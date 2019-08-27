package scene_test

// import (
// 	. "fgame/fgame/game/game"
// 	"fgame/fgame/game/game/types"
// 	"math"
// 	"testing"
// )

// type TestBattleObject struct {
// 	MaxHP             int32
// 	Attack            int32
// 	Defend            int32
// 	DamageAttack      int32
// 	DamageDefend      int32
// 	DamageAttackValue int32
// 	DamageDefendValue int32
// }

// func (tbo *TestBattleObject) GetBattleProperty(typ types.BattlePropertyType) int32 {
// 	switch typ {
// 	case types.BattlePropertyTypeAttack:
// 		return tbo.Attack
// 	case types.BattlePropertyTypeDefend:
// 		return tbo.Defend
// 	case types.BattlePropertyTypeMaxHP:
// 		return tbo.MaxHP
// 	case types.BattlePropertyTypeDamageAdd:
// 		return tbo.DamageAttackValue
// 	case types.BattlePropertyTypeDamageAddPercent:
// 		return tbo.DamageAttack
// 	case types.BattlePropertyTypeDamageDefend:
// 		return tbo.DamageDefendValue
// 	case types.BattlePropertyTypeDamageDefendPercent:
// 		return tbo.DamageDefend
// 	}
// 	return 0
// }

// func (tbo *TestBattleObject) GetBattleFinalProperty(typ types.BattleFinalPropertyType) int32 {
// 	switch typ {
// 	case types.BattleFinalPropertyTypeDamageAttack:
// 		return tbo.DamageAttack
// 	case types.BattleFinalPropertyTypeDamageDefend:
// 		return tbo.DamageDefend
// 	}
// 	return 0
// }

// type testObj struct {
// 	Id        int
// 	AttackObj BattleObject
// 	DefendObj BattleObject
// 	//技能伤害加成
// 	SkillDamageAttack float64
// 	//技能额外伤害
// 	SkillExtraDamage float64

// 	//计算最终伤害
// 	GetDamage int64
// }

// var (
// 	testObjs = []*testObj{
// 		&testObj{
// 			Id: 1,
// 			AttackObj: &TestBattleObject{
// 				MaxHP:             0,
// 				Attack:            2000,
// 				Defend:            0,
// 				DamageAttack:      1000,
// 				DamageDefend:      0,
// 				DamageAttackValue: 1000,
// 				DamageDefendValue: 0,
// 			},
// 			DefendObj: &TestBattleObject{
// 				MaxHP:             15000,
// 				Attack:            0,
// 				Defend:            100,
// 				DamageAttack:      0,
// 				DamageDefend:      10000,
// 				DamageAttackValue: 0,
// 				DamageDefendValue: 10000,
// 			},
// 			SkillDamageAttack: 0.2,
// 			SkillExtraDamage:  1000,
// 			GetDamage:         -9000,
// 		},
// 		&testObj{
// 			Id: 10,
// 			AttackObj: &TestBattleObject{
// 				MaxHP:             0,
// 				Attack:            1100,
// 				Defend:            0,
// 				DamageAttack:      1000,
// 				DamageDefend:      0,
// 				DamageAttackValue: 0,
// 				DamageDefendValue: 0,
// 			},
// 			DefendObj: &TestBattleObject{
// 				MaxHP:             15000,
// 				Attack:            0,
// 				Defend:            1000,
// 				DamageAttack:      0,
// 				DamageDefend:      0,
// 				DamageAttackValue: 0,
// 				DamageDefendValue: 0,
// 			},
// 			SkillDamageAttack: 1.1,
// 			SkillExtraDamage:  100,
// 			GetDamage:         458,
// 		},
// 	}
// )

// func TestGetBasicDamage(t *testing.T) {
// 	for _, testObj := range testObjs {
// 		damage := GetBasicDamage(testObj.AttackObj, testObj.DefendObj, testObj.SkillDamageAttack, testObj.SkillExtraDamage)
// 		if int64(math.Ceil(damage)) == testObj.GetDamage {
// 			continue
// 		}
// 		t.Fatalf("id %d expect %d,but get %d", testObj.Id, testObj.GetDamage, int64(math.Ceil(damage)))
// 	}
// }
