package types

type BattleEventType string

const (
	EventTypePlayerEnterBattle  BattleEventType = "PlayerEnterBattle"
	EventTypePlayerExitBattle                   = "PlayerExitBattle"
	EventTypePlayerEnterPVP                     = "PlayerEnterPVP"
	EventTypePlayerExitPVP                      = "PlayerExitPVP"
	EventTypePlayerSyncNeighbor                 = "PlayerSyncNeighbor"
	EventTypePlayerCheckHP                      = "PlayerCheckHP"
)

const (
	EventTypeBattlePlayerEnterScene    = "BattlePlayerEnterScene"
	EventTypeBattlePlayerExitScene     = "BattlePlayerExitScene"
	EventTypeBattlePlayerMove          = "BattlePlayerMove"
	EventTypeBattlePlayerBackLastScene = "BattlePlayerBackLastScene"
	EventTypeBattlePlayerMountCheck    = "BattlePlayerMountCheck"
	EventTypeBattlePlayerPkCheck       = "BattlePlayerPkCheck"
)

const (
	EventTypeBattlePlayerPkStateSwitch  = "BattlePlayerStateSwitch"
	EventTypeBattlePlayerPkValueChanged = "BattlePlayerPkValueChanged"
)

const (
	EventTypeBattlePlayerSpeedChanged            = "BattleObjectSpeedChanged"
	EventTypeBattlePlayerMaxHPChanged            = "BattleObjectMaxHPChanged"
	EventTypeBattlePlayerHPChanged               = "BattleObjectHPChanged"
	EventTypeBattlePlayerForceChanged            = "BattleObjectForceChanged"
	EventTypeBattlePlayerPropertyChanged         = "BattlePlayerPropertyChanged"
	EventTypeBattlePlayerAutoReborn              = "BattlePlayerAutoReborn"
	EventTypeBattlePlayerMaxTPChanged            = "BattlePlayerMaxTPChanged"
	EventTypeBattlePlayerTPChanged               = "BattlePlayerTPChanged"
	EventTypeBattlePlayerXueChiRecover           = "BattlePlayerXueChiRecover"
	EventTypeBattlePlayerXueChiBloodAdd          = "BattlePlayerXueChiBloodAdd"
	EventTypeBattlePlayerXueChiBloodLineChanged  = "BattlePlayerXueChiBloodLineChanged"
	EventTypeBattlePlayerXueChiBloodSync         = "BattlePlayerXueChiBloodSync"
	EventTypeBattlePlayerReliveRefresh           = "BattlePlayerReliveRefresh"
	EventTypeBattlePlayerRelive                  = "BattlePlayerRelive"
	EventTypeBattlePlayerReliveSync              = "BattlePlayerReliveSync"
	EventTypeBattlePlayerLevelChanged            = "BattlePlayerLevelChanged"
	EventTypeBattlePlayerZhuanShengChanged       = "BattlePlayerZhuanShengChanged"
	EventTypeBattlePlayerSoulAwakenChanged       = "BattlePlayerSoulAwakenChanged"
	EventTypeBattlePlayerVipChanged              = "BattlePlayerVipChanged"
	EventTypeBattlePlayerDenseWatNumChanged      = "BattlePlayerDenseWatNumChanged"
	EventTypeBattlePlayerDenseWatEndTimeSet      = "BattlePlayerDenseWatEndTimeSet"
	EventTypeBattlePlayerDenseWatSync            = "BattlePlayerDenseWatSync"
	EventTypeBattlePlayerShenMoGongXunNumChanged = "BattlePlayerShenMoGongXunNumChanged"
	EventTypeBattlePlayerShenMoKillNumChanged    = "BattlePlayerShenMoKillNumChanged"
	EventTypeBattlePlayerShenMoEndTimeSet        = "BattlePlayerShenMoEndTimeSet"
	EventTypeBattlePlayerShenMoSync              = "BattlePlayerShenMoSync"
)

const (
	EventTypeBattlePlayerShowFashionChanged       = "BattlePlayerShowFashionChanged"
	EventTypeBattlePlayerShowWeaponChanged        = "BattlePlayerShowWeaponChanged"
	EventTypeBattlePlayerShowTitleChanged         = "BattlePlayerShowTitleChanged"
	EventTypeBattlePlayerShowWingChanged          = "BattlePlayerShowWingChanged"
	EventTypeBattlePlayerShowMountChanged         = "BattlePlayerShowMountChanged"
	EventTypeBattlePlayerShowMountHidden          = "BattlePlayerShowMountHidden"
	EventTypeBattlePlayerShowMountSync            = "BattlePlayerShowMountSync"
	EventTypeBattlePlayerShowShenFaChanged        = "BattlePlayerShowShenFaChanged"
	EventTypeBattlePlayerShowLingYuChanged        = "BattlePlayerShowLingYuChanged"
	EventTypeBattlePlayerShowFourGodKeyChanged    = "BattlePlayerShowFourGodKeyChanged"
	EventTypeBattlePlayerShowRealmChanged         = "BattlePlayerShowRealmChanged"
	EventTypeBattlePlayerShowSpouseChanged        = "BattlePlayerShowSpouseChanged"
	EventTypeBattlePlayerShowWeddingStatusChanged = "BattlePlayerShowWeddingStatusChanged"
	EventTypeBattlePlayerShowModelChanged         = "BattlePlayerShowModelChanged"
	EventTypeBattlePlayerShowRingTypeChanged      = "BattlePlayerShowRingTypeChanged"
	EventTypeBattlePlayerShowFaBaoChanged         = "BattlePlayerShowFaBaoChanged"
	EventTypeBattlePlayerShowPetChanged           = "BattlePlayerShowPetChanged"
	EventTypeBattlePlayerShowXianTiChanged        = "BattlePlayerShowXianTiChanged"
	EventTypeBattlePlayerShowBaGuaChanged         = "BattlePlayerShowBaGuaChanged"
	EventTypeBattlePlayerShowFlyPetChanged        = "BattlePlayerShowFlyPetChanged"
	EventTypeBattlePlayerShowShenYuKeyChanged     = "BattlePlayerShowShenYuKeyChanged"
)

const (
	EventTypeBattlePlayerArenaReliveTimeChanged = "BattlePlayerArenaReliveTimeChanged"
	EventTypeBattlePlayerArenaWinTimeChanged    = "BattlePlayerArenaWinTimeChanged"
)

const (
	EventTypeBattlePlayerArenapvpReliveTimesChanged = "BattlePlayerArenapvpReliveTimesChanged"
)

const (
	EventTypeBattlePlayerTeamChanged     = "BattlePlayerTeamChanged"
	EventTypeBattlePlayerAllianceChanged = "BattlePlayerAllianceChanged"
)

const (
	EventTypeBattlePlayerGuaJiGetItem    = "BattlePlayerGuaJiGetItem"
	EventTypeBattlePlayerLingTongChanged = "BattlePlayerLingTongChanged"
	EventTypeBattlePlayerLingTongShow    = "BattlePlayerLingTongShow"
)

const (
	EventTypeBattlePlayerTeamCopyReliveTimeChanged = "BattlePlayerTeamCopyReliveTimeChanged"
)

const (
	EventTypeBattlePlayerActivityExit = "BattlePlayerActivityExit"
)

const (
	EventTypeBattlePlayerActivityPkDataChanged = "BattlePlayerActivityPkDataChanged"
	EventTypeBattlePlayerActivityPkDataSync    = "BattlePlayerActivityPkDataSync"

	EventTypeBattlePlayerActivityRankDataRefresh = "BattlePlayerActivityRankDataRefresh"
	EventTypeBattlePlayerActivityRankDataChanged = "BattlePlayerActivityRankDataChanged"

	EventTypeBattlePlayerActivityCollectDataRefresh = "BattlePlayerActivityCollectDataRefresh"
	EventTypeBattlePlayerActivityCollectDataChanged = "BattlePlayerActivityCollectDataChanged"

	EventTypeBattlePlayerActivityTickRewDataChanged = "BattlePlayerActivityTickRewDataChanged"
)

const (
	EventTypeBattlePlayerJieYiChanged   = "BattlePlayerJieYiChanged"
	EventTypeBattlePlayerCampChanged    = "BattlePlayerCampChanged"
	EventTypeBattlePlayerGuanZhiChanged = "BattlePlayerGuanZhiChanged"
	EventTypeBattlePlayerCollect        = "BattlePlayerCollect"
)

const (
	EventTypeBattlePlayerBossReliveDataSync = "PlayerBossReliveDataSync"
)

const (
	EventTypeBattlePlayerChuangShiReliveTimesChanged = "BattlePlayerChuangShiReliveTimesChanged"
)
