package types

type PlayerEventType string

const (
	EventTypePlayerLoadFinish            = "PlayerLoadFinish"
	EventTypePlayerAfterLoadFinish       = "PlayerAfterLoadFinish"
	EventTypePlayerLogoutBeforeLoaded    = "PlayerLogoutBeforeLoaded"
	EventTypePlayerBeforeLogout          = "PlayerBeforeLogout"
	EventTypePlayerExitSceneBeforeLogout = "PlayerExitSceneBeforeLogout"
	EventTypePlayerExitCrossBeforeLogout = "PlayerExitCrossBeforeLogout"
	EventTypePlayerLogout                = "PlayerLogout"
	EventTypePlayerLogoutCrossInGame     = "PlayerLogoutCrossInGame"
	EventTypePlayerLogoutCrossInCross    = "PlayerLogoutCrossInCross"
	EventTypePlayerLogoutCrossInGlobal   = "PlayerLogoutCrossInGlobal"

	EventTypePlayerOnlineTimeChanged = "PlayerOnlineTimeChangend"
	EventTypePlayerNameChanged       = "PlayerNameChanged"
	EventTypePlayerSexChanged        = "PlayerSexChanged"
	EventTypePlayerStats             = "PlayerStats"
)

type PlayerSystemEventType string

const (
	EventTypeOnlineNumSync PlayerSystemEventType = "OnlineNumSync"
)
