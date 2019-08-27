package service

import (
	"context"
	constant "fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	gmError "fgame/fgame/gm/gamegm/error"
	channelmodel "fgame/fgame/gm/gamegm/gm/channel/model"
	"fgame/fgame/gm/gamegm/pkg/timeutils"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

type IChannelService interface {
	AddChannel(p_name string) error
	UpdateChannel(p_channelid int64, p_name string) error
	DeleteChannel(p_channelid int64) error

	GetChannelList(p_name string, p_index int) ([]*channelmodel.ChannelInfo, error)
	GetAllChannelList() ([]*channelmodel.ChannelInfo, error)
	GetChannelCount(p_name string) (int, error)

	GetChannelInfo(p_id int64) (*channelmodel.ChannelInfo, error)
}

type channelService struct {
	db gmdb.DBService
}

func (m *channelService) AddChannel(p_name string) error {
	if len(p_name) == 0 {
		return gmError.GetError(gmError.ErrorCodeChannelEmpty)
	}

	exflag, err := m.existsChannel(p_name)
	if err != nil {
		return err
	}
	if exflag {
		return gmError.GetError(gmError.ErrorCodeChannelExist)
	}

	modelInfo := &channelmodel.ChannelInfo{
		ChannelName: p_name,
	}
	err = m.saveChannel(modelInfo)
	if err != nil {
		return err
	}
	return nil
}

func (m *channelService) UpdateChannel(p_channelid int64, p_name string) error {
	if len(p_name) == 0 {
		return gmError.GetError(gmError.ErrorCodeChannelEmpty)
	}

	exflag, err := m.existsChannelWithId(p_channelid, p_name)
	if err != nil {
		return err
	}
	if exflag {
		return gmError.GetError(gmError.ErrorCodeChannelExist)
	}

	modelInfo := &channelmodel.ChannelInfo{
		ChannelID:   p_channelid,
		ChannelName: p_name,
	}
	err = m.saveChannel(modelInfo)
	if err != nil {
		return err
	}
	return nil
}

func (m *channelService) DeleteChannel(p_channelid int64) error {
	now := timeutils.TimeToMillisecond(time.Now())
	errdb := m.db.DB().Table("t_channel").Where("channelId = ?", p_channelid).Update("deleteTime", now)
	if errdb.Error != nil {
		log.WithFields(log.Fields{
			"channelid": p_channelid,
			"error":     errdb.Error,
		}).Error("删除渠道失败")
		return errdb.Error
	}
	return nil
}

func (m *channelService) GetChannelList(p_name string, p_index int) ([]*channelmodel.ChannelInfo, error) {
	return m.getChannelList(p_name, p_index)
}

func (m *channelService) GetAllChannelList() ([]*channelmodel.ChannelInfo, error) {
	return m.getAllChannelList()
}

func (m *channelService) GetChannelCount(p_name string) (int, error) {
	return m.getChannelCount(p_name)
}

func (m *channelService) GetChannelInfo(p_id int64) (*channelmodel.ChannelInfo, error) {
	info := &channelmodel.ChannelInfo{}
	exdb := m.db.DB().Where("channelId = ?", p_id).First(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return info, nil
}

func (m *channelService) existsChannel(p_name string) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_channel").Where("channelName = ? and deleteTime = 0", p_name).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"channelName": p_name,
			"error":       exdb.Error,
		}).Error("查询渠道信息失败existsChannel")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *channelService) existsChannelWithId(p_channelid int64, p_name string) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_channel").Where("channelName = ? and channelId != ? and deleteTime=0", p_name, p_channelid).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"channelName": p_name,
			"channelid":   p_channelid,
			"error":       exdb.Error,
		}).Error("查询渠道信息失败existsChannelWithId")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *channelService) saveChannel(p_info *channelmodel.ChannelInfo) error {
	exdb := m.db.DB().Save(p_info)
	now := timeutils.TimeToMillisecond(time.Now())
	if p_info.ChannelID > 0 {
		p_info.CreateTime = now
	} else {
		p_info.UpdateTime = now
	}
	if exdb.Error != nil {
		log.WithFields(log.Fields{
			"channelName": p_info.ChannelID,
			"channelid":   p_info.ChannelName,
			"error":       exdb.Error,
		}).Error("保存渠道失败")
		return exdb.Error
	}
	return nil
}

func (m *channelService) getChannelList(p_name string, p_index int) ([]*channelmodel.ChannelInfo, error) {
	offset := (p_index - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize
	rst := make([]*channelmodel.ChannelInfo, 0)

	exerr := m.db.DB().Where("deleteTime =0 and channelName like ?", "%"+p_name+"%").Offset(offset).Limit(limit).Find(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"channelName": p_name,
			"error":       exerr.Error,
		}).Error("获取渠道列表失败")
		return nil, exerr.Error
	}
	return rst, nil
}

func (m *channelService) getAllChannelList() ([]*channelmodel.ChannelInfo, error) {

	rst := make([]*channelmodel.ChannelInfo, 0)

	exerr := m.db.DB().Where("deleteTime=0").Find(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"error": exerr.Error,
		}).Error("获取ALL渠道列表失败")
		return nil, exerr.Error
	}
	return rst, nil
}

func (m *channelService) getChannelCount(p_name string) (int, error) {
	rst := 0

	exerr := m.db.DB().Table("t_channel").Where("deleteTime =0 and channelName like ?", "%"+p_name+"%").Count(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"channelName": p_name,
			"error":       exerr.Error,
		}).Error("获取渠道列表失败")
		return 0, exerr.Error
	}
	return rst, nil
}

func NewChannelService(p_db gmdb.DBService) IChannelService {
	rst := &channelService{
		db: p_db,
	}
	return rst
}

type contextKey string

const (
	channelServiceKey = contextKey("ChannelService")
)

func WithChannelService(ctx context.Context, ls IChannelService) context.Context {
	return context.WithValue(ctx, channelServiceKey, ls)
}

func ChannelServiceInContext(ctx context.Context) IChannelService {
	us, ok := ctx.Value(channelServiceKey).(IChannelService)
	if !ok {
		return nil
	}
	return us
}

func SetupChannelServiceHandler(ls IChannelService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithChannelService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
