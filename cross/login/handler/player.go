package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	coretypes "fgame/fgame/core/types"
	lingtongpbutil "fgame/fgame/cross/lingtong/pbutil"
	"fgame/fgame/cross/login/login"
	crossplayerlogic "fgame/fgame/cross/player/logic"
	playerpbutil "fgame/fgame/cross/player/pbutil"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	alliancecommon "fgame/fgame/game/alliance/common"
	"fgame/fgame/game/battle/battle"
	battlecommon "fgame/fgame/game/battle/common"
	buffcommon "fgame/fgame/game/buff/common"
	crosstypes "fgame/fgame/game/cross/types"
	densewatcommon "fgame/fgame/game/densewat/common"
	jieyicommon "fgame/fgame/game/jieyi/common"
	"fgame/fgame/game/lingtong/lingtong"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	pkcommon "fgame/fgame/game/pk/common"
	playercommon "fgame/fgame/game/player/common"
	relivecommon "fgame/fgame/game/relive/common"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"
	shenmocommon "fgame/fgame/game/shenmo/common"
	skillcommon "fgame/fgame/game/skill/common"
	teamcommon "fgame/fgame/game/team/common"
	xuechicommon "fgame/fgame/game/xuechi/common"
	"fgame/fgame/pkg/idutil"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_PLAYER_DATA_TYPE), dispatch.HandlerFunc(handlePlayerData))
}

//处理玩家信息推送
func handlePlayerData(s session.Session, msg interface{}) error {
	log.Info("login:处理玩家信息推送")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player().(*player.Player)
	siPlayerData := msg.(*crosspb.SIPlayerData)
	playerCommonObj := playerpbutil.ConvertFromPlayerBasicData(siPlayerData.GetPlayerBasicData())
	playerShowObj := playerpbutil.ConvertFromPlayerShowData(siPlayerData.GetPlayerShowData())
	skillList := playerpbutil.ConvertFromSkillDataList(siPlayerData.GetSkillList())
	pkObject := playerpbutil.ConvertFromPkData(siPlayerData.GetPkData())
	battlePropertyObj := playerpbutil.ConvertFromBattleProperty(siPlayerData.GetBattlePropertyData())
	basePropertyObj := playerpbutil.ConvertFromBaseProperty(siPlayerData.GetBasicPropertyData())
	teamObj := playerpbutil.ConvertFromTeamData(siPlayerData.GetTeamData())
	allianceObj := playerpbutil.ConvertFromAllianceData(siPlayerData.GetAllianceData())
	crossType := crosstypes.CrossType(siPlayerData.GetCrossData().GetCrossType())
	crossArgs := siPlayerData.GetCrossData().GetArgs()
	playerArenaObj := playerpbutil.ConvertFromArenaData(siPlayerData.GetArenaData())
	arenapvpObj := playerpbutil.ConvertFromArenapvpData(siPlayerData.GetArenapvpData())
	xueChiObj := playerpbutil.ConvertFromXueChiData(siPlayerData.GetXueChiData())
	reliveObj := playerpbutil.ConvertFromReliveData(siPlayerData.GetReliveData())
	battleObj := playerpbutil.ConvertFromBattleData(siPlayerData.GetBattleData())
	denseWatObj := playerpbutil.ConvertFromDenseWatData(siPlayerData.GetDenseWatData())
	shenMoObj := playerpbutil.ConvertFromShenMoData(siPlayerData.GetShenMoData())
	pkDataList := playerpbutil.ConvertFromActivityPkDataList(siPlayerData.GetActivityPkDataList())
	rankDataList := playerpbutil.ConvertFromActivityRankDataList(siPlayerData.GetActivityRankDataList())
	buffDataList := playerpbutil.ConvertFromBuffDataList(siPlayerData.GetBuffList())
	jieYiObj := playerpbutil.ConvertFromJieYiData(siPlayerData.GetJieYiData())
	// chuangShiObj := playerpbutil.ConvertCommonPlayerChuangShiObject(siPlayerData.GetChuangShiData())
	bossReliveDataList := playerpbutil.ConvertFromBossReliveDataList(siPlayerData.GetBossReliveDataList())
	teShuSkillDataList := playerpbutil.ConvertFromTeShuSkillDataList(siPlayerData.GetTeShuSkillDataList())

	power := siPlayerData.GetPower()
	var lingTong scene.LingTong
	if siPlayerData.LingTongData != nil {
		id, _ := idutil.GetId()
		pos := coretypes.Position{}
		angle := float64(0.0)
		lingTongId := siPlayerData.GetLingTongData().GetLingTongId()
		lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
		name := siPlayerData.GetLingTongData().GetName()
		if lingTongTemplate != nil {
			lingTongShowData := lingtongpbutil.ConvertFromLingTongShowData(siPlayerData.GetLingTongData().GetLingTongShowData())
			lingTongPropertyData := playerpbutil.ConvertFromBattleProperty(siPlayerData.GetLingTongData().GetBattlePropertyData())
			lingTong = lingtong.CreateLingTong(pl, id, name, pos, angle, lingTongTemplate, lingTongShowData, lingTongPropertyData)
		}

	}

	err := loadPlayerData(pl,
		playerCommonObj,
		playerShowObj,
		pkObject,
		skillList,
		buffDataList,
		basePropertyObj,
		battlePropertyObj,
		teamObj,
		allianceObj,
		crossType,
		crossArgs,
		playerArenaObj,
		arenapvpObj,
		xueChiObj,
		reliveObj,
		battleObj,
		denseWatObj,
		shenMoObj,
		power,
		lingTong,
		pkDataList,
		rankDataList,
		jieYiObj,
		// chuangShiObj,
		bossReliveDataList,
		teShuSkillDataList,
	)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": gcs.Player().GetId(),
			}).Error("login:玩家推送消息,失败")
		return err
	}

	log.Info("login:处理玩家推送消息完成")
	return nil
}

//登陆
func loadPlayerData(pl *player.Player,
	commonObj playercommon.PlayerCommonObject,
	showObj *battle.PlayerShowObject,
	pkObject pkcommon.PlayerPkObject,
	skillList []skillcommon.SkillObject,
	buffDataList []buffcommon.BuffObject,
	basePropertyObj map[int32]int64,
	battlePropertyObj map[int32]int64,
	teamObj teamcommon.PlayerTeamObject,
	allianceObj alliancecommon.PlayerAllianceObject,
	crossType crosstypes.CrossType,
	crossArgs []string,
	arenaObj *battle.PlayerArenaObject,
	arenapvpObj *battle.PlayerArenapvpObject,
	xueChiObj *xuechicommon.PlayerXueChiObject,
	reliveObj *relivecommon.PlayerReliveObject,
	battleObj *battlecommon.PlayerBattleObject,
	denseWatObj *densewatcommon.PlayerDenseWatObject,
	shenMoObj *shenmocommon.PlayerShenMoObject,
	power int64,
	lingTong scene.LingTong,
	killDataList []*scene.PlayerActvitiyKillData,
	rankDataList []*scene.PlayerActvitiyRankData,
	jieYiObj jieyicommon.PlayerJieYiObject,
	// chuangShiObj chuangshidata.CommonPlayerChuangShiObject,
	bossReliveDataList []*scene.PlayerBossReliveData,
	teShuSkillDataList []*scene.TeshuSkillObject,
) (err error) {

	//加载各个组件
	showServerId := true
	switch crossType {
	case crosstypes.CrossTypeArena,
		crosstypes.CrossTypeTeamCopy:
		showServerId = false
		break
	}
	flag := pl.Load(
		commonObj,
		showObj,
		pkObject,
		basePropertyObj,
		battlePropertyObj,
		skillList,
		buffDataList,
		teamObj,
		allianceObj,
		arenaObj,
		arenapvpObj,
		xueChiObj,
		reliveObj,
		battleObj,
		denseWatObj,
		shenMoObj,
		crossType,
		power,
		lingTong,
		killDataList,
		rankDataList,
		jieYiObj,
		// chuangShiObj,
		bossReliveDataList,
		teShuSkillDataList,
		showServerId,
	)
	if !flag {
		log.WithFields(
			log.Fields{}).Warn("login:玩家加载数据,失败")
		crossplayerlogic.ExitCross(pl)
		return
	}

	//加载后
	loginHandler := login.GetLoginHandler(crossType)
	if loginHandler == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"crossType": crossType.String(),
			}).Warn("login:登陆处理器没有")
		crossplayerlogic.ExitCross(pl)
		return
	}

	if pl.IsGuaJiPlayer() {
		pl.StartGuaJi(scenetypes.GuaJiTypeCross)
	}
	if len(crossArgs) < 1 {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"crossType": crossType.String(),
			}).Warn("login:参数太少")
		crossplayerlogic.ExitCross(pl)
		return
	}

	if len(crossArgs) > 1 {
		crossBehaviorInt, err := strconv.ParseInt(crossArgs[0], 10, 64)
		if err != nil {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"crossType": crossType.String(),
					"err":       err,
					"behavior":  crossArgs[0],
				}).Warn("login:解析错误")
			crossplayerlogic.ExitCross(pl)
			return nil
		}
		crossBehavior := crosstypes.CrossBehaviorType(crossBehaviorInt)
		if crossBehavior == crosstypes.CrossBehaviorTypeTrack {
			if len(crossArgs) <= 1 {
				log.WithFields(
					log.Fields{
						"playerId":  pl.GetId(),
						"crossType": crossType.String(),
						"behavior":  crossArgs[0],
					}).Warn("login:参数不足,没有可以跟踪的玩家")
				crossplayerlogic.ExitCross(pl)
				return nil
			}
			trackPlayerId, err := strconv.ParseInt(crossArgs[1], 10, 64)
			if err != nil {
				log.WithFields(
					log.Fields{
						"playerId":      pl.GetId(),
						"crossType":     crossType.String(),
						"err":           err,
						"behavior":      crossArgs[0],
						"trackPlayerId": crossArgs[1],
					}).Warn("login:解析错误")
				crossplayerlogic.ExitCross(pl)
				return nil
			}
			trackPlayer := player.GetOnlinePlayerManager().GetPlayerById(trackPlayerId)
			if trackPlayer == nil {
				log.WithFields(
					log.Fields{
						"playerId":      pl.GetId(),
						"crossType":     crossType.String(),
						"err":           err,
						"behavior":      crossArgs[0],
						"trackPlayerId": crossArgs[1],
					}).Warn("login:玩家不在跨服")
				crossplayerlogic.ExitCross(pl)
				return nil
			}
			s := trackPlayer.GetScene()
			if s == nil {
				log.WithFields(
					log.Fields{
						"playerId":      pl.GetId(),
						"crossType":     crossType.String(),
						"err":           err,
						"behavior":      crossArgs[0],
						"trackPlayerId": crossArgs[1],
					}).Warn("login:玩家不在场景中")
				crossplayerlogic.ExitCross(pl)
				return nil
			}
			pos := trackPlayer.GetPos()
			scenelogic.PlayerEnterScene(pl, s, pos)
			return nil
		}
	}
	crossArgs = crossArgs[1:]
	flag = loginHandler.Login(pl, crossType, crossArgs...)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"crossType": crossType.String(),
			}).Warn("login:登陆失败")
		crossplayerlogic.ExitCross(pl)
		return
	}

	return nil
}
