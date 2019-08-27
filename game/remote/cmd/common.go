package cmd

const (
	ErrorCodeCommonCmdNoFound ErrorCode = ErrorCodeCommon + 1 + iota
	ErrorCodeCommonCmdHandlerNoFound
	ErrorCodeCommonPlayerNoExist
	ErrorCodeCommonArgumentInvalid
	ErrorCodeCommonPlayerExist
	ErrorCodeCommonPlayerNoSuitableName
	ErrorCodeCommonArgumentDirty
	ErrorCodeCommonPlatformWrong
	ErrorCodeCommonServerWrong
	ErrorCodeCommonFirstChargeFailed
	ErrorCodeCommonAllianceActivity
)

var (
	errorCodeCommonMap = map[ErrorCode]string{
		ErrorCodeCommonCmdNoFound:           "指令没找到",
		ErrorCodeCommonCmdHandlerNoFound:    "指令处理器没找到",
		ErrorCodeCommonPlayerNoExist:        "角色不存在",
		ErrorCodeCommonArgumentInvalid:      "参数错误",
		ErrorCodeCommonPlayerExist:          "重复创建角色",
		ErrorCodeCommonPlayerNoSuitableName: "没有合适的名字",
		ErrorCodeCommonArgumentDirty:        "包含敏感非法字符",
		ErrorCodeCommonPlatformWrong:        "平台错误",
		ErrorCodeCommonServerWrong:          "服务器错误",
		ErrorCodeCommonFirstChargeFailed:    "首充重置失败,还在活动期间",
		ErrorCodeCommonAllianceActivity:     "仙盟活动中,不能解散",
	}
)

func init() {
	MergeErrorCodeMap(errorCodeCommonMap)
}

const (
	ErrorCodeMailFormatWrong ErrorCode = ErrorCodeMail + 1 + iota
	ErrorCodeMailAttachmentFormatWrong
	ErrorCodeMailPlayerFormatWrong
)

var (
	errorCodeMailMap = map[ErrorCode]string{
		ErrorCodeMailFormatWrong:           "邮件格式错误",
		ErrorCodeMailPlayerFormatWrong:     "邮件玩家格式错误",
		ErrorCodeMailAttachmentFormatWrong: "邮件附件格式错误",
	}
)

func init() {
	MergeErrorCodeMap(errorCodeMailMap)
}

const (
	ErrorCodePrivilegeChargeGoldWrong ErrorCode = ErrorCodePrivilegeCharge + 1 + iota
	ErrorCodePrivilegeChargeSDKWrong
)

var (
	errorCodePrivilegeChargeMap = map[ErrorCode]string{
		ErrorCodePrivilegeChargeGoldWrong: "扶持金额不对",
		ErrorCodePrivilegeChargeSDKWrong:  "玩家sdk无效",
	}
)

func init() {
	MergeErrorCodeMap(errorCodePrivilegeChargeMap)
}
