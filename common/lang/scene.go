package lang

const (
	SceneMapNoExist = SceneBase + iota
	SceneRepeatEnter
	SceneExitSceneFailed
	SceneClosed
	SceneNotWorldScene
	SceneNotFuBenScene
	SceneNPCNoExist
	SceneMapNoTransfer
	SceneMoveOutside
	ScenePlayerBattleStatus
	SceneEnterTypeError
)

var (
	sceneLangMap = map[LangCode]string{
		SceneMapNoExist:         "场景地图不存在",
		SceneRepeatEnter:        "当前已在该场景",
		SceneExitSceneFailed:    "退出场景失败",
		SceneClosed:             "场景暂未开放",
		SceneNotWorldScene:      "场景不是世界场景",
		SceneNotFuBenScene:      "场景不是副本场景",
		SceneNPCNoExist:         "npc不存在",
		SceneMapNoTransfer:      "场景地图不支持传送",
		SceneMoveOutside:        "您已超出当前副本区域！",
		ScenePlayerBattleStatus: "PK状态无法退出场景，请稍后再试",
		SceneEnterTypeError:     "场景不支持的进入方式",
	}
)

func init() {
	mergeLang(sceneLangMap)
}
