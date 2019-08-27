package template

import (
	"fgame/fgame/core/template"
	friendtypes "fgame/fgame/game/friend/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sync"
)

//好友推送配置
type FriendNoticeTemplateService interface {
	GetFriendNoticeTemplate(typ friendtypes.FriendNoticeType) []*gametemplate.FriendNoticeTemplate
	GetFriendNoticeTemplateByCondition(typ friendtypes.FriendNoticeType, condition int32) *gametemplate.FriendNoticeTemplate
	GetFriendNoticeConstanTemplate() *gametemplate.NoticeConstantTemplate
	GetFriendAddTemplate(friendNum int32) *gametemplate.FriendAddTemplate
	// 好友添加奖励最大数量
	GetFriendRewMaxAddNum() int32
}

type friendNoticeTemplateService struct {
	//好友赞赏配置
	friendNoticeListMap map[friendtypes.FriendNoticeType][]*gametemplate.FriendNoticeTemplate
	//好友/仇人推送常量配置
	noticeConstant *gametemplate.NoticeConstantTemplate
	//添加好友奖励配置
	friendAddMap map[int32]*gametemplate.FriendAddTemplate
}

//初始化
func (s *friendNoticeTemplateService) init() error {
	s.friendNoticeListMap = make(map[friendtypes.FriendNoticeType][]*gametemplate.FriendNoticeTemplate)

	//好友推送
	friendNoticeTempMap := template.GetTemplateService().GetAll((*gametemplate.FriendNoticeTemplate)(nil))
	for _, temp := range friendNoticeTempMap {
		friendNoticeTemplate, _ := temp.(*gametemplate.FriendNoticeTemplate)
		s.friendNoticeListMap[friendNoticeTemplate.GetFriendNoticeType()] = append(s.friendNoticeListMap[friendNoticeTemplate.GetFriendNoticeType()], friendNoticeTemplate)
	}

	// 推送常量配置
	noticeConstantTempMap := template.GetTemplateService().GetAll((*gametemplate.NoticeConstantTemplate)(nil))
	if len(noticeConstantTempMap) != 1 {
		return fmt.Errorf("friend:推送常量配置应该有且只有一条[%d]", len(noticeConstantTempMap))
	}
	for _, temp := range noticeConstantTempMap {
		noticeConstantTemplate, _ := temp.(*gametemplate.NoticeConstantTemplate)
		s.noticeConstant = noticeConstantTemplate
	}

	//添加好友奖励配置
	s.friendAddMap = make(map[int32]*gametemplate.FriendAddTemplate)
	friendAddTempMap := template.GetTemplateService().GetAll((*gametemplate.FriendAddTemplate)(nil))
	for _, temp := range friendAddTempMap {
		friendAddTemplate, _ := temp.(*gametemplate.FriendAddTemplate)
		s.friendAddMap[friendAddTemplate.Num] = friendAddTemplate
	}

	return nil
}

func (s *friendNoticeTemplateService) GetFriendNoticeTemplate(typ friendtypes.FriendNoticeType) []*gametemplate.FriendNoticeTemplate {
	tempList, ok := s.friendNoticeListMap[typ]
	if !ok {
		return nil
	}

	return tempList
}

func (s *friendNoticeTemplateService) GetFriendNoticeTemplateByCondition(typ friendtypes.FriendNoticeType, condition int32) *gametemplate.FriendNoticeTemplate {
	tempList, ok := s.friendNoticeListMap[typ]
	if !ok {
		return nil
	}

	for _, temp := range tempList {
		if temp.TiaoJian == condition {
			return temp
		}
	}

	return nil
}

func (s *friendNoticeTemplateService) GetFriendNoticeConstanTemplate() *gametemplate.NoticeConstantTemplate {
	return s.noticeConstant
}

func (s *friendNoticeTemplateService) GetFriendAddTemplate(friendNum int32) *gametemplate.FriendAddTemplate {
	temp, ok := s.friendAddMap[friendNum]
	if !ok {
		return nil
	}

	return temp
}

func (s *friendNoticeTemplateService) GetFriendRewMaxAddNum() int32 {
	maxAddNum := int32(0)
	for _, temp := range s.friendAddMap {
		if maxAddNum < temp.Num {
			maxAddNum = temp.Num
		}
	}
	return maxAddNum
}

var (
	once sync.Once
	cs   *friendNoticeTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &friendNoticeTemplateService{}
		err = cs.init()
	})
	return err
}

func GetFriendNoticeTemplateService() FriendNoticeTemplateService {
	return cs
}
