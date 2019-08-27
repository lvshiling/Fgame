package lang

const (
	CommonArgumentInvalid = CommonBase + iota
	CommonNameDirty
	CommonFuncNoOpen
	CommonOperFrequent
	CommonCollectNoDistance
	CommonCollectOtherExist
	CommonCollectIsDead
	CommonCollectNoNum
	CommonTemplateNotExist
	CommonHandlerNotExist
)

var (
	commonLangMap = map[LangCode]string{
		CommonArgumentInvalid:   "参数无效",
		CommonNameDirty:         "名称含有脏字",
		CommonFuncNoOpen:        "功能暂未开放",
		CommonOperFrequent:      "操作过于频繁,请等待冷却时间后再试",
		CommonCollectNoDistance: "您当前不在采集范围内",
		CommonCollectOtherExist: "其它玩家正在采集",
		CommonCollectIsDead:     "该采集物已被采集过,请等待重生",
		CommonCollectNoNum:      "您采集数量已达上限,无法继续采集",
		CommonTemplateNotExist:  "模板不存在",
		CommonHandlerNotExist:   "处理器不存在",
	}
)

func init() {
	mergeLang(commonLangMap)
}
