package chat

import (
	"fgame/fgame/game/center/center"
	"fgame/fgame/game/chat/dao"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/global"
	"fgame/fgame/pkg/idutil"
	"sync"
)

type ChatService interface {
	GetWorldChatList() []*ChatData                                                     //获取世界纪录
	AddWorldChat(sendId int64, content []byte, args string, msgType chattypes.MsgType) //添加世界纪录
	GetSystemChatList() []*ChatData                                                    //获取系统纪录
	AddSystemChat(content []byte, msgType chattypes.MsgType)                           //添加系统纪录
	ChatSet(worldVipLevel, worldLevel, allianceVipLevel, allianceLevel, privateVipLevel, privateLevel, teamVipLevel, teamLevel int32)
	GetWorldLevel() int32
	GetWorldVipLevel() int32
	GetPrivateLevel() int32
	GetPrivateVipLevel() int32
	GetAllianceLevel() int32
	GetAllianceVipLevel() int32
	GetTeamLevel() int32
	GetTeamVipLevel() int32
}

type ChatData struct {
	sendId   int64
	content  []byte
	msgType  chattypes.MsgType
	sendTime int64
	args     string
}

func (d *ChatData) GetSendId() int64 {
	return d.sendId
}

func (d *ChatData) GetSendTime() int64 {
	return d.sendTime
}

func (d *ChatData) GetContent() []byte {
	return d.content
}

func (d *ChatData) GetMsgType() chattypes.MsgType {
	return d.msgType
}

func (d *ChatData) GetArgs() string {
	return d.args
}

func CreateChatData(sendId int64, content []byte, args string, msgType chattypes.MsgType) *ChatData {
	now := global.GetGame().GetTimeService().Now()
	d := &ChatData{
		sendId:   sendId,
		content:  content,
		msgType:  msgType,
		args:     args,
		sendTime: now,
	}
	return d
}

const (
	maxChatLen = 20
)

type chatService struct {
	rwm            sync.RWMutex
	worldChatList  []*ChatData
	systemChatList []*ChatData
	//聊天设置
	chatSettingObject *ChatSettingObject
}

func (s *chatService) GetWorldChatList() []*ChatData {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.worldChatList
}

func (s *chatService) AddWorldChat(sendId int64, content []byte, args string, msgType chattypes.MsgType) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	d := CreateChatData(sendId, content, args, msgType)
	s.worldChatList = append(s.worldChatList, d)

	curLen := len(s.worldChatList)
	if curLen > maxChatLen {
		startIndex := curLen - maxChatLen
		s.worldChatList = s.worldChatList[startIndex:]
	}
}

func (s *chatService) GetSystemChatList() []*ChatData {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.systemChatList
}

func (s *chatService) AddSystemChat(content []byte, msgType chattypes.MsgType) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	d := CreateChatData(0, content, "", msgType)
	s.systemChatList = append(s.systemChatList, d)

	curLen := len(s.systemChatList)
	if curLen > maxChatLen {
		startIndex := curLen - maxChatLen
		s.systemChatList = s.systemChatList[startIndex:]
	}
}

func (s *chatService) ChatSet(worldVipLevel, worldLevel, allianceVipLevel, allianceLevel, privateVipLevel, privateLevel, teamVipLevel, teamLevel int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	serverId := global.GetGame().GetServerIndex()
	now := global.GetGame().GetTimeService().Now()
	if s.chatSettingObject == nil {
		obj := createChatSettingObject()
		obj.id, _ = idutil.GetId()
		obj.serverId = serverId
		obj.worldVipLevel = worldVipLevel
		obj.worldLevel = worldLevel
		obj.allianceVipLevel = allianceVipLevel
		obj.allianceLevel = allianceLevel
		obj.privateLevel = privateLevel
		obj.privateVipLevel = privateVipLevel
		obj.teamLevel = teamLevel
		obj.teamVipLevel = teamVipLevel
		obj.createTime = now
		obj.SetModified()
		s.chatSettingObject = obj
		return
	}
	s.chatSettingObject.worldVipLevel = worldVipLevel
	s.chatSettingObject.worldLevel = worldLevel
	s.chatSettingObject.allianceVipLevel = allianceVipLevel
	s.chatSettingObject.allianceLevel = allianceLevel
	s.chatSettingObject.privateLevel = privateLevel
	s.chatSettingObject.privateVipLevel = privateVipLevel
	s.chatSettingObject.teamLevel = teamLevel
	s.chatSettingObject.teamVipLevel = teamVipLevel
	s.chatSettingObject.updateTime = now
	s.chatSettingObject.SetModified()

}

func (s *chatService) GetWorldLevel() int32 {
	if s.chatSettingObject != nil {
		return s.chatSettingObject.worldLevel
	}
	chatSet := center.GetCenterService().GetChatSet()
	if chatSet != nil {
		return chatSet.GetWorldPlayerLevel()
	}
	return constant.GetConstantService().GetConstant(constanttypes.ConstantTypeChatWorldLevel)
}

func (s *chatService) GetWorldVipLevel() int32 {
	if s.chatSettingObject != nil {
		return s.chatSettingObject.worldVipLevel
	}
	chatSet := center.GetCenterService().GetChatSet()
	if chatSet != nil {
		return chatSet.GetWorldVip()
	}
	return constant.GetConstantService().GetConstant(constanttypes.ConstantTypeChatWorldVipLevel)
}

func (s *chatService) GetPrivateLevel() int32 {
	if s.chatSettingObject != nil {
		return s.chatSettingObject.privateLevel
	}
	chatSet := center.GetCenterService().GetChatSet()
	if chatSet != nil {
		return chatSet.GetPrivatePlayerLevel()
	}
	return constant.GetConstantService().GetConstant(constanttypes.ConstantTypeChatPrivateLevel)
}

func (s *chatService) GetPrivateVipLevel() int32 {
	if s.chatSettingObject != nil {
		return s.chatSettingObject.privateVipLevel
	}
	chatSet := center.GetCenterService().GetChatSet()
	if chatSet != nil {
		return chatSet.GetPrivateVip()
	}
	return constant.GetConstantService().GetConstant(constanttypes.ConstantTypeChatPrivateVipLevel)
}

func (s *chatService) GetAllianceLevel() int32 {
	if s.chatSettingObject != nil {
		return s.chatSettingObject.allianceLevel
	}
	chatSet := center.GetCenterService().GetChatSet()
	if chatSet != nil {
		return chatSet.GetAlliancePlayerLevel()
	}
	return constant.GetConstantService().GetConstant(constanttypes.ConstantTypeChatAllianceLevel)
}

func (s *chatService) GetAllianceVipLevel() int32 {
	if s.chatSettingObject != nil {
		return s.chatSettingObject.allianceVipLevel
	}
	chatSet := center.GetCenterService().GetChatSet()
	if chatSet != nil {
		return chatSet.GetAllianceVip()
	}
	return constant.GetConstantService().GetConstant(constanttypes.ConstantTypeChatAllianceVipLevel)
}

func (s *chatService) GetTeamLevel() int32 {
	if s.chatSettingObject != nil {
		return s.chatSettingObject.teamLevel
	}
	chatSet := center.GetCenterService().GetChatSet()
	if chatSet != nil {
		return chatSet.GetTeamPlayerLevel()
	}
	return 0
}

func (s *chatService) GetTeamVipLevel() int32 {
	if s.chatSettingObject != nil {
		return s.chatSettingObject.teamVipLevel
	}
	chatSet := center.GetCenterService().GetChatSet()
	if chatSet != nil {
		return chatSet.GetTeamVip()
	}
	return 0
}

func (s *chatService) init() (err error) {
	serverId := global.GetGame().GetServerIndex()
	chatSetttingEntity, err := dao.GetChatDao().GetChatSetting(serverId)
	if err != nil {
		return
	}
	if chatSetttingEntity != nil {
		s.chatSettingObject = createChatSettingObject()
		err = s.chatSettingObject.FromEntity(chatSetttingEntity)
		if err != nil {
			return
		}
	}
	return nil
}

var (
	once sync.Once
	cs   *chatService
)

func Init() (err error) {
	once.Do(func() {
		cs = &chatService{}
		err = cs.init()
	})
	return err
}

func GetChatService() ChatService {
	return cs
}
