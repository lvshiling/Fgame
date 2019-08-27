package logic

import (
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/drop/drop"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"math"
)

//掉落
func CustomItemDrop(sb scene.Scene, pos coretypes.Position, ownerId int64, itemId int32, num int32, stack int32, protectedTime int32, existTime int32) (flag bool) {
	tempIndex := int32(0)
	//添加掉落
	numPerStack := num / stack
	remain := num % stack
	bindType := itemtypes.ItemBindTypeUnBind

	if numPerStack > 0 {
		for i := 0; i < int(stack)-1; i++ {
			nextIndex, destPos := getDropItemPosition(sb.MapTemplate(), pos, tempIndex)
			dropItem := drop.CreateDropItem(scenetypes.DropOwnerTypePlayer, ownerId, itemId, 0, numPerStack, bindType, destPos, protectedTime, existTime)
			sb.AddSceneObject(dropItem)
			tempIndex = nextIndex

		}
	}
	_, destPos := getDropItemPosition(sb.MapTemplate(), pos, tempIndex)
	dropItem := drop.CreateDropItem(scenetypes.DropOwnerTypePlayer, ownerId, itemId, 0, numPerStack+remain, bindType, destPos, protectedTime, existTime)
	sb.AddSceneObject(dropItem)

	return true
}

//掉落
func CustomDrop(sb scene.Scene, n scene.NPC, pos coretypes.Position, ownerId int64, dropIdList []int32, numTimes int32) (flag bool) {
	if len(dropIdList) != 0 {
		for i := 0; i < int(numTimes); i++ {
			tempIndex := int32(0)
			for _, dropId := range dropIdList {
				dropTemplate := droptemplate.GetDropTemplateService().GetDropFromGroup(dropId)
				if dropTemplate == nil {
					continue
				}
				num := dropTemplate.RandomNum()
				stack := dropTemplate.RandomStack()
				level := dropTemplate.RandomGoldEquipLevel()
				bindType := dropTemplate.GetBindType()
				upstar := dropTemplate.RandomGoldEquipUpstarLevel()
				attrList, isRandom := dropTemplate.RandomGoldEquipAttr()
				if !isRandom {
					itemTemp := item.GetItemService().GetItem(int(dropTemplate.ItemId))
					if itemTemp == nil {
						continue
					}

					if itemTemp.IsGoldEquip() {
						attrList = itemTemp.GetGoldEquipTemplate().RandomGoldEquipAttr()
						isRandom = true
					}
				}
				//设置随机位置
				//添加掉落
				numPerStack := num / stack
				remain := num % stack

				if numPerStack > 0 {
					for i := 0; i < int(stack)-1; i++ {
						nextIndex, destPos := getDropItemPosition(sb.MapTemplate(), pos, tempIndex)
						dropItem := drop.CreateDropItemPropertyData(scenetypes.DropOwnerTypePlayer, ownerId, dropTemplate.ItemId, level, upstar, attrList, numPerStack, bindType, destPos, dropTemplate.ProtectedTime, dropTemplate.ExistTime)
						sb.AddSceneObject(dropItem)
						tempIndex = nextIndex
					}
				}
				nextIndex, destPos := getDropItemPosition(sb.MapTemplate(), pos, tempIndex)
				dropItem := drop.CreateDropItem(scenetypes.DropOwnerTypePlayer, ownerId, dropTemplate.ItemId, level, numPerStack+remain, bindType, destPos, dropTemplate.ProtectedTime, dropTemplate.ExistTime)
				sb.AddSceneObject(dropItem)
				tempIndex = nextIndex
				gameevent.Emit(sceneeventtypes.EventTypeBattleCustomDrop, n, dropItem)
			}
		}
	}
	return true
}

//掉落
func Drop(n scene.NPC, ownerType scenetypes.DropOwnerType, ownerId int64, numTimes int32) (flag bool) {
	sb := n.GetScene()
	if sb == nil {
		return
	}
	dropIdList := n.GetBiologyTemplate().GetDropIdList()
	if len(dropIdList) != 0 {
		for i := 0; i < int(numTimes); i++ {
			tempIndex := int32(0)
			for _, dropId := range dropIdList {
				dropTemplate := droptemplate.GetDropTemplateService().GetDropFromGroup(dropId)
				if dropTemplate == nil {
					continue
				}
				num := dropTemplate.RandomNum()
				stack := dropTemplate.RandomStack()
				level := dropTemplate.RandomGoldEquipLevel()
				bindType := dropTemplate.GetBindType()
				upstar := dropTemplate.RandomGoldEquipUpstarLevel()
				attrList, isRandom := dropTemplate.RandomGoldEquipAttr()
				if !isRandom {
					itemTemp := item.GetItemService().GetItem(int(dropTemplate.ItemId))
					if itemTemp == nil {
						continue
					}

					if itemTemp.IsGoldEquip() {
						attrList = itemTemp.GetGoldEquipTemplate().RandomGoldEquipAttr()
						isRandom = true
					}
				}
				//设置随机位置
				//添加掉落
				numPerStack := num / stack
				remain := num % stack

				if numPerStack > 0 {
					for i := 0; i < int(stack)-1; i++ {
						nextIndex, destPos := getDropItemPosition(sb.MapTemplate(), n.GetPosition(), tempIndex)
						dropItem := drop.CreateDropItemPropertyData(ownerType, ownerId, dropTemplate.ItemId, level, upstar, attrList, numPerStack, bindType, destPos, dropTemplate.ProtectedTime, dropTemplate.ExistTime)
						sb.AddSceneObject(dropItem)
						tempIndex = nextIndex
					}
				}
				nextIndex, destPos := getDropItemPosition(sb.MapTemplate(), n.GetPosition(), tempIndex)
				dropItem := drop.CreateDropItemPropertyData(ownerType, ownerId, dropTemplate.ItemId, level, upstar, attrList, numPerStack+remain, bindType, destPos, dropTemplate.ProtectedTime, dropTemplate.ExistTime)
				sb.AddSceneObject(dropItem)
				tempIndex = nextIndex
				gameevent.Emit(sceneeventtypes.EventTypeBattleObjectDrop, n, dropItem)
			}
		}
	}
	return true
}

const (
	maxIndex     = 9
	elapseRadius = 1
	maxCycle     = 100
)

func getDropItemPosition(mapTemplate *gametemplate.MapTemplate, centerPosition coretypes.Position, index int32) (nextIndex int32, pos coretypes.Position) {
	nextIndex = index
	//以防卡死
	for i := 0; i < maxCycle; i++ {
		elapase := nextIndex/maxIndex + 1
		tempIndex := index % maxIndex
		radius := float64(elapseRadius * elapase)
		//获取
		angle := float64(tempIndex)/float64(maxIndex)*math.Pi*2 - math.Pi/2
		destPos := coretypes.Position{
			X: centerPosition.X + math.Cos(angle)*radius,
			Y: centerPosition.Y,
			Z: centerPosition.Z + math.Sin(angle)*radius,
		}

		if mapTemplate.GetMap().IsMask(destPos.X, destPos.Z) {
			destPos.Y = mapTemplate.GetMap().GetHeight(destPos.X, destPos.Z)
			return nextIndex + 1, destPos
		}
		nextIndex += 1
	}
	return nextIndex, centerPosition
}
