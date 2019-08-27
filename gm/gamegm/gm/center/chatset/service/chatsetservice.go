package service

import (
	"context"
	constant "fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	gmError "fgame/fgame/gm/gamegm/error"
	chatSetmodel "fgame/fgame/gm/gamegm/gm/center/chatset/model"
	"fgame/fgame/gm/gamegm/pkg/timeutils"
	"fmt"
	"net/http"
	"time"

	"fgame/fgame/gm/gamegm/common"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

type IChatSetService interface {
	AddChatSet(p_platformid int, p_serverId int, p_worldVip int, p_worldPlayerLevel int, p_pChatVip int, p_pChatPlayerLevel int, p_guildVip int, p_guildPlayerLevel int, p_sdkType int, p_teamVip int, p_teamPlayerLevel int) error
	UpdateChatSet(p_id int, p_platformid int, p_serverId int, p_worldVip int, p_worldPlayerLevel int, p_pChatVip int, p_pChatPlayerLevel int, p_guildVip int, p_guildPlayerLevel int, p_sdkType int, p_teamVip int, p_teamPlayerLevel int) error
	DeleteChatSet(p_chatSetid int64) error
	GetChatSet(p_platformid int, p_serverId int) (*chatSetmodel.ChatSetInfo, error)

	GetChatSetList(p_platformid int, p_serverId int, p_index int, p_platformList []int64) ([]*chatSetmodel.ChatSetInfo, error)
	GetAllChatSetList() ([]*chatSetmodel.ChatSetInfo, error)
	GetChatSetCount(p_platformid int, p_serverId int, p_platformList []int64) (int, error)

	GetChatSetInfo(p_id int64) (*chatSetmodel.ChatSetInfo, error)

	AddChatSetPlatform(p_platformid int, p_worldVip int, p_worldPlayerLevel int, p_pChatVip int, p_pChatPlayerLevel int, p_guildVip int, p_guildPlayerLevel int, p_teamVip int, p_teamPlayerLevel int) error
	UpdateChatSetPlatform(p_id int, p_platformid int, p_worldVip int, p_worldPlayerLevel int, p_pChatVip int, p_pChatPlayerLevel int, p_guildVip int, p_guildPlayerLevel int, p_teamVip int, p_teamPlayerLevel int) error
	DeleteChatSetPlatform(p_chatSetid int64) error
	GetChatSetPlatform(p_platformid int) (*chatSetmodel.PlatformChatSetInfo, error)
	GetChatSetPlatformById(p_id int) (*chatSetmodel.PlatformChatSetInfo, error)
	GetChatSetListPlatform(p_platformid int, p_index int, p_platformList []int64) ([]*chatSetmodel.PlatformChatSetInfo, error)
	GetChatSetCountPlatform(p_platformid int, p_platformList []int64) (int, error)
}

type chatSetService struct {
	db gmdb.DBService
}

func (m *chatSetService) AddChatSet(p_platformid int, p_serverId int, p_worldVip int, p_worldPlayerLevel int, p_pChatVip int, p_pChatPlayerLevel int, p_guildVip int, p_guildPlayerLevel int, p_sdkType int, p_teamVip int, p_teamPlayerLevel int) error {
	if p_platformid < 0 || p_serverId < 0 {
		return gmError.GetError(gmError.ErrorCodeChatSetEmpty)
	}

	exflag, err := m.existsChatSet(p_platformid, p_serverId)
	if err != nil {
		return err
	}
	if exflag {
		return gmError.GetError(gmError.ErrorCodeChatSetExist)
	}

	modelInfo := &chatSetmodel.ChatSetInfo{
		PlatformId:       p_platformid,
		ServerId:         p_serverId,
		WorldVip:         p_worldVip,
		WorldPlayerLevel: p_worldPlayerLevel,
		PChatVip:         p_pChatVip,
		PChatPlayerLevel: p_pChatPlayerLevel,
		GuildVip:         p_guildVip,
		GuildPlayerLevel: p_guildPlayerLevel,
		SdkType:          p_sdkType,
		TeamVip:          p_teamVip,
		TeamPlayerLevel:  p_teamPlayerLevel,
	}
	err = m.saveChatSet(modelInfo)
	if err != nil {
		return err
	}
	return nil
}

func (m *chatSetService) UpdateChatSet(p_id int, p_platformid int, p_serverId int, p_worldVip int, p_worldPlayerLevel int, p_pChatVip int, p_pChatPlayerLevel int, p_guildVip int, p_guildPlayerLevel int, p_sdkType int, p_teamVip int, p_teamPlayerLevel int) error {
	if p_platformid < 0 || p_serverId < 0 {
		return gmError.GetError(gmError.ErrorCodeChatSetEmpty)
	}

	exflag, err := m.existsChatSetWithId(p_id, p_platformid, p_serverId)
	if err != nil {
		return err
	}
	if exflag {
		return gmError.GetError(gmError.ErrorCodeChatSetExist)
	}

	modelInfo := &chatSetmodel.ChatSetInfo{
		Id:               p_id,
		PlatformId:       p_platformid,
		ServerId:         p_serverId,
		WorldVip:         p_worldVip,
		WorldPlayerLevel: p_worldPlayerLevel,
		PChatVip:         p_pChatVip,
		PChatPlayerLevel: p_pChatPlayerLevel,
		GuildVip:         p_guildVip,
		GuildPlayerLevel: p_guildPlayerLevel,
		SdkType:          p_sdkType,
		TeamVip:          p_teamVip,
		TeamPlayerLevel:  p_teamPlayerLevel,
	}
	err = m.saveChatSet(modelInfo)
	if err != nil {
		return err
	}
	return nil
}

func (m *chatSetService) DeleteChatSet(p_chatSetid int64) error {
	now := timeutils.TimeToMillisecond(time.Now())
	errdb := m.db.DB().Table("t_chat_set").Where("id = ?", p_chatSetid).Update("deleteTime", now)
	if errdb.Error != nil {
		log.WithFields(log.Fields{
			"chatSetid": p_chatSetid,
			"error":     errdb.Error,
		}).Error("删除聊天配置失败")
		return errdb.Error
	}
	return nil
}

func (m *chatSetService) GetChatSet(p_platformid int, p_serverId int) (*chatSetmodel.ChatSetInfo, error) {
	info := &chatSetmodel.ChatSetInfo{}
	exdb := m.db.DB().Where("platformId = ? and serverId = ? and deleteTime = 0", p_platformid, p_serverId).First(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"platformId": p_platformid,
			"serverId":   p_serverId,
			"error":      exdb.Error,
		}).Error("查询聊天配置信息失败existsChatSet")
		return nil, exdb.Error
	}
	return info, nil
}

func (m *chatSetService) GetChatSetList(p_platformid int, p_serverId int, p_index int, p_platformList []int64) ([]*chatSetmodel.ChatSetInfo, error) {
	return m.getChatSetList(p_platformid, p_serverId, p_index, p_platformList)
}

func (m *chatSetService) GetAllChatSetList() ([]*chatSetmodel.ChatSetInfo, error) {
	return m.getAllChatSetList()
}

func (m *chatSetService) GetChatSetCount(p_platformid int, p_serverId int, p_platformList []int64) (int, error) {
	return m.getChatSetCount(p_platformid, p_serverId, p_platformList)
}

func (m *chatSetService) GetChatSetInfo(p_id int64) (*chatSetmodel.ChatSetInfo, error) {
	info := &chatSetmodel.ChatSetInfo{}
	exdb := m.db.DB().Where("id = ?", p_id).First(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return info, nil
}

func (m *chatSetService) existsChatSet(p_platformid int, p_serverId int) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_chat_set").Where("platformId = ? and serverId = ? and deleteTime = 0", p_platformid, p_serverId).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"platformId": p_platformid,
			"serverId":   p_serverId,
			"error":      exdb.Error,
		}).Error("查询聊天配置信息失败existsChatSet")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *chatSetService) existsChatSetWithId(p_chatSetid int, p_platformid int, p_serverId int) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_chat_set").Where("id != ? and platformId = ? and serverId = ? and deleteTime = 0", p_chatSetid, p_platformid, p_serverId).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"platformId": p_platformid,
			"serverId":   p_serverId,
			"ID":         p_chatSetid,
			"error":      exdb.Error,
		}).Error("查询聊天配置信息失败existsChatSet")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *chatSetService) saveChatSet(p_info *chatSetmodel.ChatSetInfo) error {
	exdb := m.db.DB().Save(p_info)
	now := timeutils.TimeToMillisecond(time.Now())
	if p_info.Id == 0 {
		p_info.CreateTime = now
	} else {
		p_info.UpdateTime = now
	}
	if exdb.Error != nil {
		log.WithFields(log.Fields{
			"chatSetid": p_info.PlatformId,
			"error":     exdb.Error,
		}).Error("保存聊天配置失败")
		return exdb.Error
	}
	return nil
}

func (m *chatSetService) getChatSetList(p_platformid int, p_serverId int, p_index int, p_platformList []int64) ([]*chatSetmodel.ChatSetInfo, error) {
	offset := (p_index - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize
	rst := make([]*chatSetmodel.ChatSetInfo, 0)
	where := ""
	if p_platformid > 0 {
		where += fmt.Sprintf(" and platformId=%d", p_platformid)
	}
	if p_serverId > 0 {
		where += fmt.Sprintf(" and serverId=%d", p_serverId)
	}

	if len(p_platformList) > 0 {
		where += fmt.Sprintf(" and platformId IN (%s)", common.CombinInt64Array(p_platformList))
	}

	exerr := m.db.DB().Where("deleteTime =0" + where).Order("platformId ASC,serverId ASC").Offset(offset).Limit(limit).Find(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"error": exerr.Error,
		}).Error("获取聊天配置列表失败")
		return nil, exerr.Error
	}
	return rst, nil
}

func (m *chatSetService) getAllChatSetList() ([]*chatSetmodel.ChatSetInfo, error) {

	rst := make([]*chatSetmodel.ChatSetInfo, 0)

	exerr := m.db.DB().Where("deleteTime=0").Find(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"error": exerr.Error,
		}).Error("获取ALL聊天配置列表失败")
		return nil, exerr.Error
	}
	return rst, nil
}

func (m *chatSetService) getChatSetCount(p_platformid int, p_serverId int, p_platformList []int64) (int, error) {
	rst := 0
	where := ""
	if p_platformid > 0 {
		where += fmt.Sprintf(" and platformId=%d", p_platformid)
	}
	if p_serverId > 0 {
		where += fmt.Sprintf(" and serverId=%d", p_serverId)
	}
	if len(p_platformList) > 0 {
		where += fmt.Sprintf(" and platformId IN (%s)", common.CombinInt64Array(p_platformList))
	}

	exerr := m.db.DB().Table("t_chat_set").Where("deleteTime =0" + where).Count(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"error": exerr.Error,
		}).Error("获取聊天配置列表失败")
		return 0, exerr.Error
	}
	return rst, nil
}

func NewChatSetService(p_db gmdb.DBService) IChatSetService {
	rst := &chatSetService{
		db: p_db,
	}
	return rst
}

type contextKey string

const (
	chatSetServiceKey = contextKey("ChatSetService")
)

func WithChatSetService(ctx context.Context, ls IChatSetService) context.Context {
	return context.WithValue(ctx, chatSetServiceKey, ls)
}

func ChatSetServiceInContext(ctx context.Context) IChatSetService {
	us, ok := ctx.Value(chatSetServiceKey).(IChatSetService)
	if !ok {
		return nil
	}
	return us
}

func SetupChatSetServiceHandler(ls IChatSetService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithChatSetService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
