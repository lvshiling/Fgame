package scene

import (
	coretypes "fgame/fgame/core/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//检查是否能被位移
type CheckAttackedMoveHandler interface {
	CheckAttackedMove(attackObj BattleObject, defenceObj BattleObject, attackedPos coretypes.Position) bool
}

type CheckAttackedMoveHandlerFunc func(attackObj BattleObject, defenceObj BattleObject, attackedPos coretypes.Position) bool

func (h CheckAttackedMoveHandlerFunc) CheckAttackedMove(attackObj BattleObject, defenceObj BattleObject, attackedPos coretypes.Position) bool {
	return h(attackObj, defenceObj, attackedPos)
}

var (
	checkAttackedMoveHandlerMap = map[scenetypes.SceneType]CheckAttackedMoveHandler{}
)

func RegisterCheckAttackedMoveHandler(typ scenetypes.SceneType, h CheckAttackedMoveHandler) {
	_, ok := checkAttackedMoveHandlerMap[typ]
	if ok {
		panic(fmt.Errorf("重复注册%s检查攻击", typ.String()))
	}
	checkAttackedMoveHandlerMap[typ] = h
}

func GetCheckAttackedMoveHandler(typ scenetypes.SceneType) CheckAttackedMoveHandler {
	h, ok := checkAttackedMoveHandlerMap[typ]
	if !ok {
		return nil
	}
	return h
}

//检查是否能攻击
type NPCCheckAttackHandler interface {
	CheckAttack(attackObj BattleObject, defenceObj NPC) bool
}

type NPCCheckAttackHandlerFunc func(attackObj BattleObject, defenceObj NPC) bool

func (h NPCCheckAttackHandlerFunc) CheckAttack(attackObj BattleObject, defenceObj NPC) bool {
	return h(attackObj, defenceObj)
}

var (
	npcCheckAttackHandlerMap = map[scenetypes.BiologyScriptType]NPCCheckAttackHandler{}
)

func RegisterCheckNPCAttackHandler(typ scenetypes.BiologyScriptType, h NPCCheckAttackHandler) {
	_, ok := npcCheckAttackHandlerMap[typ]
	if ok {
		panic(fmt.Errorf("重复注册%s检查攻击", typ.String()))
	}
	npcCheckAttackHandlerMap[typ] = h
}

func GetCheckNPCAttackHandler(typ scenetypes.BiologyScriptType) NPCCheckAttackHandler {
	h, ok := npcCheckAttackHandlerMap[typ]
	if !ok {
		return nil
	}
	return h
}

//检查是否能位移
type CheckMoveHandler interface {
	CheckMove(p Player, destPos coretypes.Position) (bool, coretypes.Position)
}

type CheckMoveHandlerFunc func(p Player, destPos coretypes.Position) (bool, coretypes.Position)

func (h CheckMoveHandlerFunc) CheckMove(p Player, destPos coretypes.Position) (bool, coretypes.Position) {
	return h(p, destPos)
}

var (
	checkMoveHandlerMap = map[scenetypes.SceneType]CheckMoveHandler{}
)

func RegisterCheckMoveHandler(typ scenetypes.SceneType, h CheckMoveHandler) {
	_, ok := checkMoveHandlerMap[typ]
	if ok {
		panic(fmt.Errorf("重复注册%s检查攻击", typ.String()))
	}
	checkMoveHandlerMap[typ] = h
}

func GetCheckMoveHandler(typ scenetypes.SceneType) CheckMoveHandler {
	h, ok := checkMoveHandlerMap[typ]
	if !ok {
		return nil
	}
	return h
}

//检查能否攻击
type CheckAttackHandler interface {
	CheckAttack(attackObject BattleObject) bool
}

type CheckAttackHandlerFunc func(attackObject BattleObject) bool

func (h CheckAttackHandlerFunc) CheckAttack(attackObject BattleObject) bool {
	return h(attackObject)
}

var (
	checkAttackHandlerMap = map[scenetypes.SceneType]CheckAttackHandler{}
)

func RegisterCheckAttackHandler(typ scenetypes.SceneType, h CheckAttackHandler) {
	_, ok := checkAttackHandlerMap[typ]
	if ok {
		panic(fmt.Errorf("重复注册%s检查攻击", typ.String()))
	}
	checkAttackHandlerMap[typ] = h
}

func GetCheckAttackHandler(typ scenetypes.SceneType) CheckAttackHandler {
	h, ok := checkAttackHandlerMap[typ]
	if !ok {
		return nil
	}
	return h
}
