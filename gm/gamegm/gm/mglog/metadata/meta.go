package metadata

type MsgType int

const (
	MsgTypeServerLog = iota
	MsgTypePlayerLog
	MsgTypeAllianceLog
	MsgTypeJiaoYiLog
)

type MsgItemInfo struct {
	Label      string `json:"lab"`
	DataColumn string `json:"key"`
	ShowType   string `json:"showType"`
}

type MsgInfo struct {
	TableName string `json:"key"`
	ShowName  string `json:"value"`
}

var msgServerLogMap map[string][]*MsgItemInfo
var msgPlayerLogMap map[string][]*MsgItemInfo
var msgAllianceLogMap map[string][]*MsgItemInfo
var msgJiaoYiLogMap map[string][]*MsgItemInfo

var msgServerMsgArray []*MsgInfo
var msgPlayerMsgArray []*MsgInfo
var msgAllianceMsgArray []*MsgInfo
var msgJiaoYiLogArray []*MsgInfo

func RegisterMsgItemInfo(p_name string, p_showName string, p_type MsgType, p_infoArray []*MsgItemInfo) {
	if p_type == MsgTypeServerLog {
		if _, ok := msgServerLogMap[p_name]; ok {
			return
		}
		msgServerLogMap[p_name] = p_infoArray
		info := &MsgInfo{
			TableName: p_name,
			ShowName:  p_showName,
		}
		msgServerMsgArray = append(msgServerMsgArray, info)
		return
	}
	if p_type == MsgTypePlayerLog {
		if _, ok := msgPlayerLogMap[p_name]; ok {
			return
		}
		msgPlayerLogMap[p_name] = p_infoArray

		info := &MsgInfo{
			TableName: p_name,
			ShowName:  p_showName,
		}
		msgPlayerMsgArray = append(msgPlayerMsgArray, info)
		return
	}
	if p_type == MsgTypeAllianceLog {
		if _, ok := msgAllianceLogMap[p_name]; ok {
			return
		}
		msgAllianceLogMap[p_name] = p_infoArray

		info := &MsgInfo{
			TableName: p_name,
			ShowName:  p_showName,
		}
		msgAllianceMsgArray = append(msgAllianceMsgArray, info)
		return
	}
	if p_type == MsgTypeJiaoYiLog {
		if _, ok := msgJiaoYiLogMap[p_name]; ok {
			return
		}
		msgJiaoYiLogMap[p_name] = p_infoArray

		info := &MsgInfo{
			TableName: p_name,
			ShowName:  p_showName,
		}
		msgJiaoYiLogArray = append(msgJiaoYiLogArray, info)
		return
	}
}

func GetMsgItemInfo(p_msgType MsgType, p_name string) []*MsgItemInfo {
	if p_msgType == MsgTypePlayerLog {
		return GetPlayerMsgItemInfo(p_name)
	}
	if p_msgType == MsgTypeAllianceLog {
		return GetAllianceMsgItemInfo(p_name)
	}
	if p_msgType == MsgTypeJiaoYiLog {
		return GetJiaoYiMsgItemInfo(p_name)
	}

	return GetServerMsgItemInfo(p_name)
}

func GetServerMsgItemInfo(p_name string) []*MsgItemInfo {
	if value, ok := msgServerLogMap[p_name]; ok {
		return value
	}
	return nil
}

func GetPlayerMsgItemInfo(p_name string) []*MsgItemInfo {
	if value, ok := msgPlayerLogMap[p_name]; ok {
		return value
	}
	return nil
}

func GetAllianceMsgItemInfo(p_name string) []*MsgItemInfo {
	if value, ok := msgAllianceLogMap[p_name]; ok {
		return value
	}
	return nil
}

func GetJiaoYiMsgItemInfo(p_name string) []*MsgItemInfo {
	if value, ok := msgJiaoYiLogMap[p_name]; ok {
		return value
	}
	return nil
}

func GetMstList(p_type MsgType) []*MsgInfo {
	if p_type == MsgTypePlayerLog {
		return msgPlayerMsgArray
	}
	if p_type == MsgTypeAllianceLog {
		return msgAllianceMsgArray
	}
	if p_type == MsgTypeJiaoYiLog {
		return msgJiaoYiLogArray
	}
	return msgServerMsgArray
}

func init() {
	msgServerLogMap = make(map[string][]*MsgItemInfo)
	msgPlayerLogMap = make(map[string][]*MsgItemInfo)
	msgAllianceLogMap = make(map[string][]*MsgItemInfo)
	msgServerMsgArray = make([]*MsgInfo, 0)
	msgPlayerMsgArray = make([]*MsgInfo, 0)
	msgAllianceMsgArray = make([]*MsgInfo, 0)
	msgJiaoYiLogMap = make(map[string][]*MsgItemInfo)
	msgJiaoYiLogArray = make([]*MsgInfo, 0)
}
