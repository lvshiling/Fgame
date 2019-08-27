package logic

import (
	"context"
	"encoding/json"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/game/global"
	playerinventory "fgame/fgame/game/inventory/player"
	onearenatemplate "fgame/fgame/game/onearena/template"
	onearenatypes "fgame/fgame/game/onearena/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/pkg/idutil"
	"fmt"

	onearenaentity "fgame/fgame/game/onearena/entity"
	playeronearena "fgame/fgame/game/onearena/player"
	onearenascene "fgame/fgame/game/onearena/scene"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	scenelogic "fgame/fgame/game/scene/logic"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"

	commonlog "fgame/fgame/common/log"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	inventorylogic "fgame/fgame/game/inventory/logic"
	onearenaeventtypes "fgame/fgame/game/onearena/event/types"

	log "github.com/Sirupsen/logrus"

	"fgame/fgame/game/onearena/pbutil"
)

type robbedRecord struct {
	RobName string
	Sucess  bool
	Level   onearenatypes.OneArenaLevelType
	Pos     int32
}

func newRobbedRecord(robName string, sucess bool, level onearenatypes.OneArenaLevelType, pos int32) *robbedRecord {
	d := &robbedRecord{
		RobName: robName,
		Sucess:  sucess,
		Level:   level,
		Pos:     pos,
	}
	return d
}

//玩家进入灵池抢夺
func PlayerEnterOneArena(pl player.Player, ownerId int64, ownerName string, level onearenatypes.OneArenaLevelType, pos int32) (flag bool) {
	oneArenaTemplate := onearenatemplate.GetOneArenaTemplateService().GetOneArenaTemplateByLevel(level, pos)
	if oneArenaTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"level":    level,
				"pos":      pos,
			}).Warn("onearena:处理跳转灵池,灵池不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(oneArenaTemplate.MapId)
	if mapTemplate == nil {
		return
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeLingChiFighting {
		return
	}
	sh := onearenascene.CreateOneArenaSceneData(pl, ownerId, ownerName, oneArenaTemplate)
	s := scene.CreateFuBenScene(oneArenaTemplate.MapId, sh)
	if s == nil {
		panic(fmt.Errorf("onearena:创建副本应该成功"))
	}
	scenelogic.PlayerEnterSingleFuBenScene(pl, s)
	flag = true
	return
}

//灵池被抢
func PeerRobbedRecord(oneArenaData *onearenaeventtypes.OneArenaData, robName string, sucess bool) (err error) {
	playerId := oneArenaData.GetPlayerId()
	level := oneArenaData.GetLevel()
	pos := oneArenaData.GetPos()
	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl != nil {
		robbed := newRobbedRecord(robName, sucess, level, pos)
		ctx := scene.WithPlayer(context.Background(), pl)
		playerRobbedRecordMsg := message.NewScheduleMessage(onPlayerRobbedRecord, ctx, robbed, nil)
		pl.Post(playerRobbedRecordMsg)
	} else {
		//写离线日志
		id, err := idutil.GetId()
		if err != nil {
			return err
		}
		status := onearenatypes.OneArenaRobbedStatusTypeSucess
		if !sucess {
			status = onearenatypes.OneArenaRobbedStatusTypeFail
		}
		now := global.GetGame().GetTimeService().Now()
		oneArenaRobbedEntity := &onearenaentity.PlayerOneArenaRobbedEntity{
			Id:         id,
			PlayerId:   playerId,
			RobName:    robName,
			RobTime:    now,
			Status:     int32(status),
			UpdateTime: now,
			CreateTime: now,
			DeleteTime: 0,
		}
		global.GetGame().GetGlobalUpdater().AddChangedObject(oneArenaRobbedEntity)
	}
	return
}

func onPlayerRobbedRecord(ctx context.Context, result interface{}, err error) error {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)
	robbed := result.(*robbedRecord)
	sucess := robbed.Sucess

	manager := tpl.GetPlayerDataManager(types.PlayerOneArenaDataManagerType).(*playeronearena.PlayerOneArenaDataManager)
	manager.RobbedRecord(robbed.RobName, sucess)
	if sucess {
		manager.ReplaceOneArena(robbed.Level, robbed.Pos)
		scOneArenaRobbedPush := pbutil.BuildSCOneArenaRobbedPush(robbed.RobName)
		tpl.SendMsg(scOneArenaRobbedPush)
	}
	return nil
}

//灵池产出鲲
func OneArenaOutputKun(playerId int64, level onearenatypes.OneArenaLevelType, pos int32, num int32) (err error) {
	oneArenaTempalte := onearenatemplate.GetOneArenaTemplateService().GetOneArenaTemplateByLevel(level, pos)
	if oneArenaTempalte == nil {
		return
	}
	dropMap := make(map[int32]int32)
	dropList := oneArenaTempalte.GetDropList()
	for i := int32(0); i < num; i++ {
		//固定物品掉落
		if len(dropList) != 0 {
			sweepDropMap := droptemplate.GetDropTemplateService().GetDropListItems(dropList)
			for itemId, itemNum := range sweepDropMap {
				_, ok := dropMap[itemId]
				if !ok {
					dropMap[itemId] = itemNum
				} else {
					dropMap[itemId] += itemNum
				}
			}
		}
	}

	if len(dropMap) == 0 {
		return
	}
	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl != nil {
		ctx := scene.WithPlayer(context.Background(), pl)
		playerAddKunMsg := message.NewScheduleMessage(onPlayerAddKun, ctx, dropMap, nil)
		pl.Post(playerAddKunMsg)
	} else {
		//玩家离线
		id, _ := idutil.GetId()
		itemInfoBytes, err := json.Marshal(dropMap)
		if err != nil {
			return err
		}
		now := global.GetGame().GetTimeService().Now()
		oneArenaKunEntity := &onearenaentity.PlayerOneArenaKunEntity{
			Id:         id,
			PlayerId:   playerId,
			KunInfo:    string(itemInfoBytes),
			UpdateTime: now,
			CreateTime: now,
			DeleteTime: 0,
		}
		global.GetGame().GetGlobalUpdater().AddChangedObject(oneArenaKunEntity)
	}
	return
}

func onPlayerAddKun(ctx context.Context, result interface{}, err error) (terr error) {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)
	dropMap := result.(map[int32]int32)
	PlayerAddKun(tpl, dropMap)
	return
}

func PlayerAddKun(pl player.Player, kunMap map[int32]int32) {
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	inventoryLogReason := commonlog.InventoryLogReasonOneArenaOutputKun
	reasonText := inventoryLogReason.String()

	flag := inventoryManager.HasEnoughSlots(kunMap)

	kunLimitMap := make(map[int32]int32)
	kunLimit := onearenatemplate.GetOneArenaTemplateService().GetOneArenaKunLimit()
	for itemId, num := range kunMap {
		if num > kunLimit {
			num = kunLimit
		}
		kunLimitMap[itemId] = num
	}
	if !flag {
		//写邮件
		emailTitle := lang.GetLangService().ReadLang(lang.OneArenaOutputTitle)
		emailContent := lang.GetLangService().ReadLang(lang.OneArenaOutputContent)
		emaillogic.AddEmail(pl, emailTitle, emailContent, kunLimitMap)
		return
	}

	flag = inventoryManager.BatchAdd(kunLimitMap, inventoryLogReason, reasonText)
	if !flag {
		panic(fmt.Errorf("onearena: OneArenaOutputKun BatchAdd should be ok"))
	}
	inventorylogic.SnapInventoryChanged(pl)
}
