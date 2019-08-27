package types

type LingTongEventType string

const (
	EventTypeBattleLingTongShowFashionChanged LingTongEventType = "BattleLingTongShowFashionChanged"
	EventTypeBattleLingTongShowWeaponChanged                    = "BattleLingTongShowWeaponChanged"
	EventTypeBattleLingTongShowTitleChanged                     = "BattleLingTongShowTitleChanged"
	EventTypeBattleLingTongShowWingChanged                      = "BattleLingTongShowWingChanged"
	EventTypeBattleLingTongShowMountChanged                     = "BattleLingTongShowMountChanged"
	EventTypeBattleLingTongShowMountHidden                      = "BattleLingTongShowMountHidden"
	EventTypeBattleLingTongShowShenFaChanged                    = "BattleLingTongShowShenFaChanged"
	EventTypeBattleLingTongShowLingYuChanged                    = "BattleLingTongShowLingYuChanged"
	EventTypeBattleLingTongShowFaBaoChanged                     = "BattleLingTongShowFaBaoChanged"
	EventTypeBattleLingTongShowXianTiChanged                    = "BattleLingTongShowXianTiChanged"
	EventTypeBattleLingTongChanged                              = "BattleLingTongChanged"
	EventTypeBattleLingTongExitScene                            = "BattleLingTongExitScene"
	EventTypeBattleLingTongEnterScene                           = "BattleLingTongEnterScene"
	EventTypeBattleLingTongRename                               = "BattleLingTongRename"
)

const (
	EventTypeLingTongActivate              LingTongEventType = "LingTongActivate"
	EventTypeLingTongLevelChanged                            = "LingTongLevelChanged"
	EventTypeLingTongRename                                  = "LingTongRename"
	EventTypeLingTongChuZhanChanged                          = "LingTongChuZhanChanged"
	EventTypeLingTongSystemPropertyChanged                   = "LingTongSystemPropertyChanged"
)
