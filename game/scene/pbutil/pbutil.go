package pbutil

import (
	scenepb "fgame/fgame/common/codec/pb/scene"
	uipb "fgame/fgame/common/codec/pb/ui"
	coretypes "fgame/fgame/core/types"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	commonpbutil "fgame/fgame/game/common/pbutil"
	lingtongdevtemplate "fgame/fgame/game/lingtongdev/template"
	lingtongdevtypes "fgame/fgame/game/lingtongdev/types"
	pktypes "fgame/fgame/game/pk/types"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"math"
)

var (
	scPing = &scenepb.SCPing{}
)

//场景ping
func BuildPing() *scenepb.SCPing {
	return scPing
}

//进入场景
func BuildEnterScene(player scene.Player) *scenepb.SCEnterScene {
	enterScene := &scenepb.SCEnterScene{}
	mapId := player.GetScene().MapId()
	enterScene.MapId = &mapId

	enterScene.Pos = BuildPosition(player.GetPosition())
	enterScene.PlayerData = buildScenePlayerData(player)
	return enterScene
}

//进入战斗
func BuildObjectBattle(pl scene.Player, isBattle bool) *scenepb.SCObjectBattle {
	scObjectBattle := &scenepb.SCObjectBattle{}
	scObjectBattle.BattleState = &isBattle
	return scObjectBattle
}

//进入视野
func BuildEnterScope(pl scene.Player) (enterScope *scenepb.SCObjectEnterScope, monsterList []scene.NPC) {
	enterNeighbors := pl.GetEnterNeighborsAndClear()
	if len(enterNeighbors) == 0 {
		return nil, nil
	}

	enterScope = &scenepb.SCObjectEnterScope{}
	for _, neighbor := range enterNeighbors {
		switch obj := neighbor.(type) {
		case scene.Player:
			{
				// log.WithFields(
				// 	log.Fields{
				// 		"玩家":   obj.GetId(),
				// 		"goId": runtimeutils.Goid(),
				// 	}).Infoln("玩家进入视野")
				enterScope.PlayerDataList = append(enterScope.PlayerDataList, buildScenePlayerData(obj))
				break
			}
		case scene.LingTong:
			{

				enterScope.PlayerDataList = append(enterScope.PlayerDataList, buildSceneLingTongData(obj))
				break
			}
		case scene.NPC:
			{
				if obj.IsDead() {
					break
				}
				isEnemy := pl.IsEnemy(obj)
				enterScope.MonsterDataList = append(enterScope.MonsterDataList, buildSceneNPCData(obj, isEnemy))
				monsterList = append(monsterList, obj)
				break
			}
		case scene.DropItem:
			{
				canGet := pl.IfCanGetDropItem(obj)
				enterScope.ItemDataList = append(enterScope.ItemDataList, buildSceneItemData(obj, canGet))
				break
			}
		}
	}
	return enterScope, monsterList
}

func BuildPosition(pos coretypes.Position) *scenepb.Position {
	posX := float32(pos.X)
	posY := float32(pos.Y)
	posZ := float32(pos.Z)
	tpos := &scenepb.Position{}
	tpos.PosX = &posX
	tpos.PosY = &posY
	tpos.PosZ = &posZ
	return tpos
}

func BuildSCPlayerFeiXieTransfer() *uipb.SCPlayerFeiXieTransfer {
	playerFeiXieTransfer := &uipb.SCPlayerFeiXieTransfer{}
	return playerFeiXieTransfer
}

func buildScenePlayerData(pl scene.Player) *scenepb.ScenePlayerData {
	spd := &scenepb.ScenePlayerData{}
	uid := pl.GetId()
	name := pl.GetName()
	angle := float32(pl.GetAngle())
	role := int32(pl.GetRole())
	sex := int32(pl.GetSex())
	spd.Uid = &uid
	spd.Name = &name
	spd.Pos = BuildPosition(pl.GetPosition())
	spd.Angle = &angle
	spd.Job = &role
	spd.Sex = &sex
	spd.PlayerShowData = buildScenePlayerShowData(pl)
	level := pl.GetLevel()
	spd.Level = &level

	spd.BuffList = BuildObjectBuffList(pl)
	ownerId := int64(0)
	spd.OwnerId = &ownerId
	return spd
}

func buildScenePlayerShowData(pl scene.Player) *scenepb.ScenePlayerShowData {
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	hp := pl.GetHP()
	hpMax := pl.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
	spd.Hp = &hp
	spd.HpMax = &hpMax
	titleId := pl.GetTitleId()

	spd.TitleId = &titleId
	weaponId := pl.GetWeaponId()
	spd.WeaponId = &weaponId
	weaponWaken := pl.GetWeaponState() != 0
	spd.WeaponWaken = &weaponWaken
	clothesId := pl.GetFashionId()
	spd.ClothesId = &clothesId
	rideId := int32(0)
	if !pl.IsMountHidden() {
		rideId = pl.GetMountId()
	}
	spd.RideId = &rideId
	wingId := pl.GetWingId()
	spd.WingId = &wingId
	rideLevel := int32(1)
	spd.RideLevel = &rideLevel
	bangPaiName := pl.GetAllianceName()
	spd.BangPaiName = &bangPaiName
	pkState := int32(pl.GetPkState())
	spd.PkState = &pkState
	s := pl.GetScene()
	pkValue := int32(0)
	if s.MapTemplate().IsWorld() {
		pkValue = pl.GetPkValue()
	}
	spd.PkValue = &pkValue
	pkCamp := int32(pl.GetPkCamp().Camp())
	spd.Camp = &pkCamp
	speed := int32(pl.GetBattleProperty(propertytypes.BattlePropertyTypeMoveSpeed))
	spd.Speed = &speed
	teamName := pl.GetTeamName()
	spd.TeamName = &teamName
	shenFa := pl.GetShenFaId()
	spd.ShenFaId = &shenFa
	lingYu := pl.GetLingYuId()
	spd.LingYuId = &lingYu
	fourGodKey := int32(0)
	if s.MapTemplate().GetMapType() == scenetypes.SceneTypeFourGodGate || s.MapTemplate().GetMapType() == scenetypes.SceneTypeFourGodWar {
		fourGodKey = pl.GetFourGodKey()
	}
	spd.FourGodKey = &fourGodKey
	realm := pl.GetRealm()
	spd.Realm = &realm
	spouse := pl.GetSpouse()
	spd.SpouseName = &spouse
	weddingStatus := pl.GetWeddingStatus()
	//特殊处理 结婚是2
	if !s.MapTemplate().IsMarry() && weddingStatus == 2 {
		weddingStatus = 0
	}
	spd.WeddingStatus = &weddingStatus
	vip := pl.GetVip()
	spd.Vip = &vip
	//客户端特殊处理
	ringType := pl.GetRingType()
	spd.RingType = &ringType
	petId := int32(0)
	spd.PetId = &petId
	faBaoId := pl.GetFaBaoId()
	spd.FaBaoId = &faBaoId
	xianTiId := pl.GetXianTiId()
	spd.XianTiId = &xianTiId
	flyPetId := pl.GetFlyPetId()
	spd.FlyPetId = &flyPetId
	shenYuKey := int32(0)
	if s.MapTemplate().GetMapType() == scenetypes.SceneTypeShenYu {
		shenYuKey = pl.GetShenYuKey()
	}
	spd.ShenYuKey = &shenYuKey

	if pl.GetFactionType() == scenetypes.FactionTypeModel {
		isModel := true
		spd.IsModel = &isModel
	}

	isRobot := pl.IsRobot()
	spd.IsRobot = &isRobot
	jieYiName := pl.GetSceneJieYiName()
	spd.JieYiName = &jieYiName
	camp := int32(pl.GetCamp())
	spd.ZhenYing = &camp
	guanZhi := int32(pl.GetGuanZhi())
	spd.GuanZhi = &guanZhi
	return spd
}

func buildSceneLingTongData(lingTong scene.LingTong) *scenepb.ScenePlayerData {
	spd := &scenepb.ScenePlayerData{}
	uid := lingTong.GetId()
	name := lingTong.GetName()
	angle := float32(lingTong.GetAngle())
	role := int32(0)
	sex := int32(0)
	spd.Uid = &uid
	spd.Name = &name
	spd.Pos = BuildPosition(lingTong.GetPosition())
	spd.Angle = &angle
	spd.Job = &role
	spd.Sex = &sex
	spd.PlayerShowData = buildSceneLingTongShowData(lingTong)
	level := int32(0)
	spd.Level = &level
	spd.BuffList = nil
	ownerId := int64(0)
	if lingTong.GetOwner() != nil {
		ownerId = lingTong.GetOwner().GetId()
	}
	spd.OwnerId = &ownerId
	return spd
}

func buildSceneLingTongShowData(lingTong scene.LingTong) *scenepb.ScenePlayerShowData {
	spd := &scenepb.ScenePlayerShowData{}
	uid := lingTong.GetId()
	spd.Uid = &uid
	isLingTong := int32(1)
	spd.LingTong = &isLingTong
	hp := int64(1)
	hpMax := int64(1)
	spd.Hp = &hp
	spd.HpMax = &hpMax
	titleId := lingTong.GetLingTongTitleId()

	spd.TitleId = &titleId
	weaponId := lingTong.GetLingTongWeaponId()
	if weaponId == 0 {
		weaponId = 1
	}
	spd.WeaponId = &weaponId
	weaponWaken := lingTong.GetLingTongWeaponId() != 0
	spd.WeaponWaken = &weaponWaken
	clothesId := lingTong.GetLingTongFashionId()
	if clothesId == 0 {
		clothesId = 1
	}
	spd.ClothesId = &clothesId
	rideId := int32(0)
	if !lingTong.IsLingTongMountHidden() {
		rideId = lingTong.GetLingTongMountId()
		rideId = getLingTongMountId(rideId)
	}
	spd.RideId = &rideId
	wingId := lingTong.GetLingTongWingId()
	spd.WingId = &wingId
	rideLevel := int32(1)
	spd.RideLevel = &rideLevel
	bangPaiName := ""
	spd.BangPaiName = &bangPaiName
	pkState := int32(0)
	spd.PkState = &pkState
	pkValue := int32(0)
	spd.PkValue = &pkValue
	pkCamp := int32(0)
	spd.Camp = &pkCamp
	speed := int32(math.Ceil(float64(lingTong.GetOwner().GetBattleProperty(propertytypes.BattlePropertyTypeMoveSpeed)) * 1.5))
	spd.Speed = &speed
	teamName := ""
	spd.TeamName = &teamName
	shenFa := lingTong.GetLingTongShenFaId()
	spd.ShenFaId = &shenFa
	lingYu := lingTong.GetLingTongLingYuId()
	spd.LingYuId = &lingYu
	fourGodKey := int32(0)
	spd.FourGodKey = &fourGodKey
	realm := int32(0)
	spd.Realm = &realm
	spouse := ""
	spd.SpouseName = &spouse
	weddingStatus := int32(0)
	spd.WeddingStatus = &weddingStatus
	vip := int32(0)
	spd.Vip = &vip
	//客户端特殊处理
	ringType := int32(0)
	spd.RingType = &ringType
	petId := int32(0)
	spd.PetId = &petId
	faBaoId := lingTong.GetLingTongFaBaoId()
	spd.FaBaoId = &faBaoId
	xianTiId := lingTong.GetLingTongXianTiId()
	spd.XianTiId = &xianTiId
	return spd
}

func buildSceneNPCData(npc scene.NPC, isEnemy bool) *scenepb.SceneMonsterData {
	spd := &scenepb.SceneMonsterData{}
	uid := npc.GetId()
	angle := float32(npc.GetAngle())
	spd.Uid = &uid
	spd.CurPos = BuildPosition(npc.GetPosition())
	spd.Angle = &angle
	hp := npc.GetHP()
	spd.Hp = &hp
	maxHp := npc.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
	spd.HpMax = &maxHp
	tempId := int32(npc.GetTempId())
	spd.TempID = &tempId
	spd.BuffList = BuildObjectBuffList(npc)
	spd.Enemy = &isEnemy
	name := npc.GetName()
	spd.Name = &name
	return spd
}

func buildSceneItemData(dropItem scene.DropItem, state bool) *scenepb.SceneItemData {
	d := &scenepb.SceneItemData{}
	uid := dropItem.GetId()
	d.Uid = &uid
	itemId := dropItem.GetItemId()
	d.ItemId = &itemId
	itemNum := dropItem.GetItemNum()
	d.ItemNum = &itemNum
	level := dropItem.GetLevel()
	if level > 0 {
		d.Level = &level
	}
	d.CurPos = BuildPosition(dropItem.GetPosition())
	d.State = &state
	return d
}

func BuildExitScope(pl scene.Player) *scenepb.SCObjectExitScope {
	exitNeighbors := pl.GetLeaveNeighborsAndClear()
	if len(exitNeighbors) == 0 {
		return nil
	}
	exitScope := &scenepb.SCObjectExitScope{}
	for _, neighbor := range exitNeighbors {
		so, ok := neighbor.(scene.SceneObject)
		if !ok {
			continue
		}

		switch bo := so.(type) {
		case scene.NPC:
			// log.WithFields(
			// 	log.Fields{
			// 		"玩家":    pl.GetId(),
			// 		"退出的生物": bo.GetId(),
			// 		"类型":    bo.GetBiologyTemplate().GetBiologyType(),
			// 		"goId":  runtimeutils.Goid(),
			// 	}).Infoln("生物退出视野前")
			if bo.IsDead() && bo.GetBiologyTemplate().GetBiologyType().DeadIgnore() {
				continue
			}
			// log.WithFields(
			// 	log.Fields{
			// 		"玩家":    pl.GetId(),
			// 		"退出的生物": bo.GetId(),
			// 		"goId":  runtimeutils.Goid(),
			// 	}).Infoln("生物退出视野")
			break
		case scene.Player:
			{

				// log.WithFields(
				// 	log.Fields{
				// 		"玩家":    pl.GetId(),
				// 		"退出的玩家": bo.GetId(),
				// 		"goId":  runtimeutils.Goid(),
				// 	}).Infoln("玩家退出视野")
				break
			}
		case scene.LingTong:
			{
				// log.WithFields(
				// 	log.Fields{
				// 		"玩家":    pl.GetId(),
				// 		"退出的灵童": bo.GetId(),
				// 		"goId":  runtimeutils.Goid(),
				// 	}).Infoln("灵童退出视野")
				break
			}
		}

		exitScope.ExitObjectList = append(exitScope.ExitObjectList, buildExitObject(so))

	}
	return exitScope
}

func buildExitObject(so scene.SceneObject) *scenepb.ExitObject {
	eo := &scenepb.ExitObject{}
	uId := so.GetId()
	objType := int32(so.GetSceneObjectType())
	eo.Uid = &uId
	eo.ObjecType = &objType
	return eo
}

func buildMoveData(obj scene.SceneObject, tpos coretypes.Position, moveSpeed float64, angle float64, moveType scenetypes.MoveType, moveFlag bool) *scenepb.ObjectMoveData {
	moveData := &scenepb.ObjectMoveData{}
	angleF := float32(angle)
	moveData.Angle = &angleF
	moveSpeedF := float32(moveSpeed)
	moveData.MoveSpeed = &moveSpeedF

	moveData.Pos = BuildPosition(tpos)
	moveTypeInt := int32(moveType)
	moveData.MoveType = &moveTypeInt
	clientId := obj.GetId()
	moveData.Uid = &clientId
	objType := int32(obj.GetSceneObjectType())
	moveData.ObjecType = &objType
	moveData.Flag = &moveFlag
	return moveData
}

func BuildSCObjectMove(obj scene.SceneObject, pos coretypes.Position, moveSpeed float64, angle float64, moveType scenetypes.MoveType, moveFlag bool) *scenepb.SCObjectMove {
	moveData := buildMoveData(obj, pos, moveSpeed, angle, moveType, moveFlag)
	scObjectMove := &scenepb.SCObjectMove{}
	scObjectMove.MoveData = moveData
	return scObjectMove
}

func BuildSCObjectFixPosition(obj scene.SceneObject, pos coretypes.Position) *scenepb.SCObjectFixPosition {

	scObjectFixPosition := &scenepb.SCObjectFixPosition{}

	uId := obj.GetId()
	objType := int32(obj.GetSceneObjectType())
	scObjectFixPosition.Uid = &uId
	scObjectFixPosition.ObjecType = &objType
	scObjectFixPosition.Pos = BuildPosition(pos)
	return scObjectFixPosition
}

func buildAttackData(obj scene.SceneObject, pos coretypes.Position, angle float64, skillId int32) *scenepb.ObjectAttackData {
	attackData := &scenepb.ObjectAttackData{}
	angleF := float32(angle)
	attackData.Angle = &angleF

	attackData.SkillId = &skillId

	attackData.Pos = BuildPosition(pos)
	cId := obj.GetId()
	attackData.Uid = &cId
	objType := int32(obj.GetSceneObjectType())
	attackData.ObjecType = &objType

	return attackData
}
func buildPetAttackData(obj scene.Player, objType int32, pos coretypes.Position, angle float64, skillId int32) *scenepb.ObjectAttackData {
	attackData := &scenepb.ObjectAttackData{}
	angleF := float32(angle)
	attackData.Angle = &angleF

	attackData.SkillId = &skillId

	attackData.Pos = BuildPosition(pos)
	cId := obj.GetId()
	attackData.Uid = &cId

	attackData.ObjecType = &objType

	return attackData
}

func BuildSCObjectAttack(obj scene.SceneObject, pos coretypes.Position, angle float64, skillId int32) *scenepb.SCObjectAttack {
	attackData := buildAttackData(obj, pos, angle, skillId)
	scObjectAttack := &scenepb.SCObjectAttack{}
	scObjectAttack.AttackData = attackData
	return scObjectAttack
}

func BuildSCPetAttack(pl scene.Player, objType int32, pos coretypes.Position, angle float64, skillId int32) *scenepb.SCObjectAttack {
	attackData := buildPetAttackData(pl, objType, pos, angle, skillId)
	scObjectAttack := &scenepb.SCObjectAttack{}
	scObjectAttack.AttackData = attackData
	return scObjectAttack
}

//伤害包
func BuildSCObjectDamage(bo scene.BattleObject, damageType scenetypes.DamageType, damageValue int64, skillId int32, attackId int64) *scenepb.SCObjectDamage {
	objectDamage := &scenepb.ObjectDamage{}
	hp := bo.GetHP()
	uId := bo.GetId()

	objectDamage.Hp = &hp
	objectDamage.Uid = &uId

	objType := int32(bo.GetSceneObjectType())
	objectDamage.ObjecType = &objType
	damageTypeInt := int32(damageType)
	objectDamage.DamageType = &damageTypeInt
	objectDamage.DamageVal = &damageValue
	objectDamage.AttackId = &attackId
	scObjectDamage := &scenepb.SCObjectDamage{}

	scObjectDamage.SkillId = &skillId
	scObjectDamage.ObjectDamageList = append(scObjectDamage.ObjectDamageList, objectDamage)
	return scObjectDamage

}

//复活包
func BuildSCPlayerRelive(pl scene.Player) *scenepb.SCPlayerRelive {
	scPlayerRelive := &scenepb.SCPlayerRelive{}
	scPlayerRelive.ReliveData = buildObjectRelive(pl)
	return scPlayerRelive
}

//复活包
func buildObjectRelive(bo scene.BattleObject) *scenepb.ObjectReliveData {

	objectReliveData := &scenepb.ObjectReliveData{}
	mapId := bo.GetScene().MapId()
	objectReliveData.SceneId = &mapId
	uId := bo.GetId()
	objectReliveData.Uid = &uId
	objectReliveData.Pos = BuildPosition(bo.GetPosition())
	hp := bo.GetHP()
	objectReliveData.Hp = &hp
	objectReliveData.HpMax = &hp
	return objectReliveData
}

//复活包
func BuildSCUIPlayerRelive(pl scene.Player) *uipb.SCPlayerRelive {
	scPlayerRelive := &uipb.SCPlayerRelive{}

	return scPlayerRelive
}

//pk状态改变
func BuildScenePlayerPkStateSwitched(pl scene.Player, state pktypes.PkState, camp pktypes.PkCamp) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	pkState := int32(state)
	spd.PkState = &pkState
	campInt := camp.Camp()
	spd.Camp = &campInt
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//pk值改变
func BuildScenePlayerPkValueChanged(pl scene.Player, val int32) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	spd.PkValue = &val
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//max hp改变
func BuildScenePlayerMaxHPChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	hp := pl.GetHP()
	hpMax := pl.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
	spd.Hp = &hp
	spd.HpMax = &hpMax
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//称号变化
func BuildScenePlayerTitleChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	titleId := pl.GetTitleId()
	spd.TitleId = &titleId
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//战翼变化
func BuildScenePlayerWingChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	wingId := pl.GetWingId()
	spd.WingId = &wingId
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//坐骑变化
func BuildScenePlayerMountChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	mountId := int32(0)
	if !pl.IsMountHidden() {
		mountId = pl.GetMountId()
	}
	spd.RideId = &mountId
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//时装变化
func BuildScenePlayerFashionChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	fashionId := pl.GetFashionId()
	spd.ClothesId = &fashionId
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//冰魂变化
func BuildScenePlayerWeaponChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	weaponId := pl.GetWeaponId()
	spd.WeaponId = &weaponId
	weaponWaken := pl.GetWeaponState() != 0
	spd.WeaponWaken = &weaponWaken
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//速度变化
func BuildScenePlayerSpeedChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	speed := int32(pl.GetBattleProperty(propertytypes.BattlePropertyTypeMoveSpeed))
	spd.Speed = &speed
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//帮会变化
func BuildScenePlayerAllianceChanged(pl scene.Player, val string) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	spd.BangPaiName = &val
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//帮会变化
func BuildScenePlayerTeamChanged(pl scene.Player, teamName string) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	spd.TeamName = &teamName
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//四神钥匙改变
func BuildScenePlayerFourGodKeyChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	fourGodKey := pl.GetFourGodKey()
	spd.FourGodKey = &fourGodKey
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//身法变化
func BuildScenePlayerShenfaChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	shenfaId := pl.GetShenFaId()
	spd.ShenFaId = &shenfaId
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//领域变化
func BuildScenePlayerLingyuChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	lingyuId := pl.GetLingYuId()
	spd.LingYuId = &lingyuId
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//境界改变
func BuildScenePlayerRealmChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	realm := pl.GetRealm()
	spd.Realm = &realm
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//战翼变化
func BuildScenePlayerSpouseChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	spouse := pl.GetSpouse()
	spd.SpouseName = &spouse
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//婚礼状态改变
func BuildScenePlayerWeddingStatusChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	s := pl.GetScene()
	weddingStatus := pl.GetWeddingStatus()
	//特殊处理 结婚是2
	if !s.MapTemplate().IsMarry() && weddingStatus == 2 {
		weddingStatus = 0
	}
	spd.WeddingStatus = &weddingStatus
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//模型改变
func BuildScenePlayerModelChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	model := pl.GetModel()

	spd.ModelId = &model
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//坐骑变化
func BuildScenePlayerVipChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	vip := pl.GetVip()
	spd.Vip = &vip
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//婚戒变化
func BuildScenePlayerRingChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	ringType := pl.GetRingType()

	spd.RingType = &ringType
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//法宝变化
func BuildScenePlayerFaBaoChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	fabaoId := pl.GetFaBaoId()
	spd.FaBaoId = &fabaoId
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//宠物变化
func BuildScenePlayerPetChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	petId := int32(0)
	spd.PetId = &petId
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//宠物变化
func BuildScenePlayerFlyPetChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	flyPetId := pl.GetFlyPetId()
	spd.FlyPetId = &flyPetId
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//仙体变化
func BuildScenePlayerXianTiChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	xiantiId := pl.GetXianTiId()
	spd.XianTiId = &xiantiId
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//神域钥匙变化
func BuildScenePlayerShenYuChanged(pl scene.Player) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	lingTong := int32(0)
	keyNum := pl.GetShenYuKey()

	spd.Uid = &uid
	spd.LingTong = &lingTong
	spd.ShenYuKey = &keyNum
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//帮会变化
func BuildScenePlayerJieYiChanged(pl scene.Player, val string) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	spd.JieYiName = &val
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//阵营变化
func BuildScenePlayerCampChanged(pl scene.Player, camp chuangshitypes.ChuangShiCampType, guanZhi chuangshitypes.ChuangShiGuanZhi) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	campInt := int32(camp)
	spd.ZhenYing = &campInt
	guanZhiInt := int32(guanZhi)
	spd.GuanZhi = &guanZhiInt
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//官职变化
func BuildScenePlayerGuanZhiChanged(pl scene.Player, camp chuangshitypes.ChuangShiCampType, guanZhi chuangshitypes.ChuangShiGuanZhi) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(0)
	spd.LingTong = &lingTong
	campInt := int32(camp)
	spd.ZhenYing = &campInt
	guanZhiInt := int32(guanZhi)
	spd.GuanZhi = &guanZhiInt
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//物品获得
func BuildSCItemGet(itemIds []int64) *scenepb.SCItemGet {
	scItemGet := &scenepb.SCItemGet{}
	scItemGet.ItemId = itemIds
	return scItemGet
}

func BuildExitScene(playerId int64) *scenepb.SCExitScene {
	scExitScene := &scenepb.SCExitScene{}
	scExitScene.PlayerId = &playerId
	return scExitScene
}

var (
	fuBenExit = &uipb.SCFuBenExit{}
)

//副本退出
func BuildSCFuBenExit() *uipb.SCFuBenExit {
	return fuBenExit
}

//跳转npc
func BuildSCGoToNPC(npcId int32) *uipb.SCGoToNPC {
	scGoToNPC := &uipb.SCGoToNPC{}
	scGoToNPC.NpcId = &npcId
	return scGoToNPC
}

//跳转npc
func BuildSCMonsterCampChanged(n scene.NPC, enemy bool) *scenepb.SCMonsterCampChanged {
	scMonsterCampChanged := &scenepb.SCMonsterCampChanged{}
	uId := n.GetId()
	scMonsterCampChanged.Uid = &uId
	objType := int32(n.GetSceneObjectType())
	scMonsterCampChanged.ObjecType = &objType
	scMonsterCampChanged.Enemy = &enemy
	return scMonsterCampChanged
}

//物品归属变化
func BuildSCItemOwnerChanged(di scene.DropItem, state bool) *scenepb.SCItemOwnerChanged {
	scItemOwnerChanged := &scenepb.SCItemOwnerChanged{}
	id := di.GetId()
	scItemOwnerChanged.ItemId = &id
	scItemOwnerChanged.State = &state
	return scItemOwnerChanged
}

//被击杀
func BuildSCPlayerKilled(name string) *uipb.SCPlayerKilled {
	scPlayerKilled := &uipb.SCPlayerKilled{}
	scPlayerKilled.KillName = &name
	return scPlayerKilled
}

//被击杀
func BuildSCPlayerAttacked(id int64) *uipb.SCPlayerAttacked {
	scPlayerAttacked := &uipb.SCPlayerAttacked{}
	scPlayerAttacked.AttackId = &id
	return scPlayerAttacked
}

//击杀玩家
func BuildSCPlayerKill(pl scene.Player) *uipb.SCPlayerKill {
	scPlayerKill := &uipb.SCPlayerKill{}
	playerId := pl.GetId()
	name := pl.GetName()
	role := int32(pl.GetRole())
	sex := int32(pl.GetSex())
	force := pl.GetForce()
	scPlayerKill.PlayerId = &playerId
	scPlayerKill.Role = &role
	scPlayerKill.Sex = &sex
	scPlayerKill.Name = &name
	scPlayerKill.Force = &force
	return scPlayerKill
}

var (
	scSceneHeartBeat = &scenepb.SCSceneHeartBeat{}
)

//心跳
func BuildSCSceneHeartBeat() *scenepb.SCSceneHeartBeat {

	return scSceneHeartBeat
}

//场景玩家信息
func BuildSCScenePlayerData(pl scene.Player) *uipb.SCScenePlayerDataChanged {
	tp := pl.GetTP()
	scScenePlayerDataChanged := &uipb.SCScenePlayerDataChanged{}
	scScenePlayerDataChanged.Tp = &tp
	return scScenePlayerDataChanged
}

//体力
func BuildSCScenePlayerTPChanged(pl scene.Player) *uipb.SCScenePlayerDataChanged {
	tp := pl.GetTP()
	scScenePlayerDataChanged := &uipb.SCScenePlayerDataChanged{}
	scScenePlayerDataChanged.Tp = &tp
	return scScenePlayerDataChanged
}

//技能使用
func BuildSCScenePlayerSkillUse(pl scene.Player, skillId int32) *uipb.SCScenePlayerSkillUse {
	scScenePlayerSkillUse := &uipb.SCScenePlayerSkillUse{}
	scScenePlayerSkillUse.SkillId = &skillId
	return scScenePlayerSkillUse
}

//客户端特殊处理
//血量变化
func BuildSceneLingTongHpChanged(pl scene.LingTong) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(1)
	spd.LingTong = &lingTong
	hp := int64(1)
	spd.Hp = &hp
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//称号变化
func BuildSceneLingTongTitleChanged(pl scene.LingTong) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(1)
	spd.LingTong = &lingTong
	titleId := pl.GetLingTongTitleId()
	spd.TitleId = &titleId
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//战翼变化
func BuildSceneLingTongWingChanged(pl scene.LingTong) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(1)
	spd.LingTong = &lingTong
	wingId := pl.GetLingTongWingId()
	spd.WingId = &wingId
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//战翼变化
func BuildSceneLingTongSpeedChanged(pl scene.LingTong) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(1)
	spd.LingTong = &lingTong
	speed := int32(pl.GetOwner().GetBattleProperty(propertytypes.BattlePropertyTypeMoveSpeed))
	spd.Speed = &speed
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//坐骑变化
func BuildSceneLingTongMountChanged(pl scene.LingTong) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(1)
	spd.LingTong = &lingTong
	mountId := int32(0)
	if !pl.IsLingTongMountHidden() {
		mountId = pl.GetLingTongMountId()
		//客户端需要特殊处理特殊处理
		mountId = getLingTongMountId(mountId)
	}
	spd.RideId = &mountId
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//时装变化
func BuildSceneLingTongFashionChanged(pl scene.LingTong) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(1)
	spd.LingTong = &lingTong
	fashionId := pl.GetLingTongFashionId()
	spd.ClothesId = &fashionId
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//冰魂变化
func BuildSceneLingTongWeaponChanged(pl scene.LingTong) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(1)
	spd.LingTong = &lingTong
	weaponId := pl.GetLingTongWeaponId()
	spd.WeaponId = &weaponId
	weaponWaken := pl.GetLingTongWeaponState() != 0
	spd.WeaponWaken = &weaponWaken
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//身法变化
func BuildSceneLingTongShenfaChanged(pl scene.LingTong) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(1)
	spd.LingTong = &lingTong
	shenfaId := pl.GetLingTongShenFaId()
	spd.ShenFaId = &shenfaId
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//领域变化
func BuildSceneLingTongLingyuChanged(pl scene.LingTong) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(1)
	spd.LingTong = &lingTong
	lingyuId := pl.GetLingTongLingYuId()
	spd.LingYuId = &lingyuId
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//法宝变化
func BuildSceneLingTongFaBaoChanged(pl scene.LingTong) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(1)
	spd.LingTong = &lingTong
	fabaoId := pl.GetLingTongFaBaoId()
	spd.FaBaoId = &fabaoId
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//仙体变化
func BuildSceneLingTongXianTiChanged(pl scene.LingTong) *scenepb.SCPlayerDataChanged {
	scPlayerDataChanged := &scenepb.SCPlayerDataChanged{}
	spd := &scenepb.ScenePlayerShowData{}
	uid := pl.GetId()
	spd.Uid = &uid
	lingTong := int32(1)
	spd.LingTong = &lingTong
	xiantiId := pl.GetLingTongXianTiId()
	spd.XianTiId = &xiantiId
	scPlayerDataChanged.PlayerData = spd
	return scPlayerDataChanged
}

//复活包
func BuildSCLingTongRelive(pl scene.LingTong) *scenepb.SCPlayerRelive {
	scPlayerRelive := &scenepb.SCPlayerRelive{}
	scPlayerRelive.ReliveData = buildLingTongRelive(pl)
	return scPlayerRelive
}

//复活包
func buildLingTongRelive(bo scene.BattleObject) *scenepb.ObjectReliveData {

	objectReliveData := &scenepb.ObjectReliveData{}
	mapId := bo.GetScene().MapId()
	objectReliveData.SceneId = &mapId
	uId := bo.GetId()
	objectReliveData.Uid = &uId
	objectReliveData.Pos = BuildPosition(bo.GetPosition())
	hp := int64(1)
	objectReliveData.Hp = &hp
	objectReliveData.HpMax = &hp
	return objectReliveData
}

func getLingTongMountId(mountId int32) int32 {
	if mountId != 0 {
		lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTemplate(lingtongdevtypes.LingTongDevSysTypeLingQi, int(mountId))
		lingTongMountTemplate, ok := lingTongDevTemplate.(*gametemplate.LingTongMountTemplate)
		if ok {
			mountId = lingTongMountTemplate.MountId
		}
	}
	return mountId
}

var (
	scPlayerEnterPVP = &uipb.SCPlayerEnterPVP{}
)

//进入战斗
func BuildPlayerEnterPVP(pl scene.Player) *uipb.SCPlayerEnterPVP {
	return scPlayerEnterPVP
}

var (
	scPlayerExitPVP = &uipb.SCPlayerExitPVP{}
)

//进入战斗
func BuildPlayerExitPVP(pl scene.Player) *uipb.SCPlayerExitPVP {
	return scPlayerExitPVP
}

//建立排行列表
func BuildSCSceneRankChanged(p scene.Player, r *scene.SceneRank) *uipb.SCSceneRankChanged {
	scSceneRankChanged := &uipb.SCSceneRankChanged{}
	scSceneRankChanged.RankInfo = BuildSceneRankInfo(p, r)
	return scSceneRankChanged
}

func BuildSceneRankInfo(p scene.Player, r *scene.SceneRank) *uipb.SceneRankInfo {
	rankType := r.GetRankType()
	playerInfoList := r.GetRankList()
	selfRank, selfValue := r.GetPlayerRank(p.GetId())
	rankInfo := &uipb.SceneRankInfo{}
	rankTypeInt := rankType.GetRankType()
	rankInfo.RankType = &rankTypeInt
	rankInfo.PlayerList = buildSceneRankPlayerInfoList(playerInfoList)
	rankInfo.SelfRankInfo = buildSceneSelfRank(selfRank, selfValue)

	return rankInfo
}

func BuildGeneralCollectInfoListByList(npcList []scene.NPC) (infoList []*uipb.GeneralCollectInfo) {
	for _, npc := range npcList {
		infoList = append(infoList, BuildGeneralCollectInfo(npc))
	}
	return infoList
}

func BuildGeneralCollectInfoList(npcMap map[int64]scene.NPC) (infoList []*uipb.GeneralCollectInfo) {
	for _, npc := range npcMap {
		infoList = append(infoList, BuildGeneralCollectInfo(npc))
	}
	return infoList
}

func BuildGeneralCollectInfo(npc scene.NPC) *uipb.GeneralCollectInfo {
	typ := int32(npc.GetBiologyTemplate().GetBiologyScriptType())
	status := npc.IsDead()
	statusTime := int64(0)
	if status {
		statusTime = npc.GetDeadTime()
	}
	pos := npc.GetPosition()
	biologyId := int32(npc.GetBiologyTemplate().Id)
	npcId := npc.GetId()

	bio := &uipb.GeneralCollectInfo{}
	bio.NcpId = &npcId
	bio.Typ = &typ
	bio.IsDead = &status
	bio.DeadTime = &statusTime
	bio.Pos = commonpbutil.BuildPos(pos)
	bio.BiologyId = &biologyId

	return bio
}

func buildSceneSelfRank(selfRank int32, selfValue int64) (selfRankInfo *uipb.SceneSelfRankInfo) {
	selfRankInfo = &uipb.SceneSelfRankInfo{}
	selfRankInfo.Rank = &selfRank
	selfRankInfo.Value = &selfValue

	return selfRankInfo
}

func buildSceneRankPlayerInfoList(playerInfoList []*scene.SceneRankPlayerInfo) (rankPlayerInfoList []*uipb.SceneRankPlayerInfo) {

	for _, playerInfo := range playerInfoList {
		rankPlayerInfo := buildSceneRankPlayerInfo(playerInfo)
		rankPlayerInfoList = append(rankPlayerInfoList, rankPlayerInfo)
	}

	return rankPlayerInfoList
}

func buildSceneRankPlayerInfo(playerInfo *scene.SceneRankPlayerInfo) *uipb.SceneRankPlayerInfo {
	playerId := playerInfo.GetPlayerId()
	playerName := playerInfo.GetPlayerName()
	val := playerInfo.GetValue()
	sceneRankPlayerInfo := &uipb.SceneRankPlayerInfo{}
	sceneRankPlayerInfo.PlayerId = &playerId
	sceneRankPlayerInfo.PlayerName = &playerName
	sceneRankPlayerInfo.Value = &val
	return sceneRankPlayerInfo
}
