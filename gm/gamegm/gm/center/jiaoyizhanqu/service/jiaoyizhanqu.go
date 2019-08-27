package service

import (
	"context"
	"fgame/fgame/core/db"
	constant "fgame/fgame/gm/gamegm/constant"
	jiaoyimodel "fgame/fgame/gm/gamegm/gm/center/jiaoyizhanqu/model"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

type IJiaoYiZhanQuService interface {
	AddJiaoYiZhanQu(info *jiaoyimodel.JiaoYiZhanQuEntity) error
	UpdateJiaoYiZhanQu(info *jiaoyimodel.JiaoYiZhanQuEntity) error
	DeleteZhanQu(id int32) error
	GetJiaoYiZhanQu(platformId int32, serverId int32) (*jiaoyimodel.JiaoYiZhanQuEntity, error)
	GetJiaoYiZhanQuList(platformId int32, pageIndex int32) ([]*jiaoyimodel.JiaoYiZhanQuEntity, error)
	GetJiaoYiZhanQuCount(platformId int32) (int32, error)
	GetAllJiaoYiZhanQuList(platformId int32) ([]*jiaoyimodel.JiaoYiZhanQuEntity, error)
}

type jiaoYiZhanQuService struct {
	ds db.DBService
}

func (m *jiaoYiZhanQuService) AddJiaoYiZhanQu(info *jiaoyimodel.JiaoYiZhanQuEntity) error {
	exdb := m.ds.DB().Save(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	return nil
}

func (m *jiaoYiZhanQuService) UpdateJiaoYiZhanQu(info *jiaoyimodel.JiaoYiZhanQuEntity) error {
	if info.Id < 1 {
		return fmt.Errorf("id 不能为0")
	}
	exdb := m.ds.DB().Save(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	return nil
}

func (m *jiaoYiZhanQuService) DeleteZhanQu(id int32) error {
	if id < 1 {
		return nil
	}
	now := timeutils.TimeToMillisecond(time.Now())
	exdb := m.ds.DB().Table("t_jiaoyi_zhanqu").Where("id=?", id).Updates(map[string]interface{}{"deleteTime": now})
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil
	}
	return exdb.Error
}

func (m *jiaoYiZhanQuService) GetJiaoYiZhanQu(platformId int32, serverId int32) (*jiaoyimodel.JiaoYiZhanQuEntity, error) {
	info := &jiaoyimodel.JiaoYiZhanQuEntity{}
	exdb := m.ds.DB().Where("platformId=? and serverId=? and deleteTime=0", platformId, serverId).First(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return info, nil
}

func (m *jiaoYiZhanQuService) GetJiaoYiZhanQuList(platformId int32, pageIndex int32) ([]*jiaoyimodel.JiaoYiZhanQuEntity, error) {
	rst := make([]*jiaoyimodel.JiaoYiZhanQuEntity, 0)
	offect := (pageIndex - 1) * constant.DefaultPageSize
	if offect < 0 {
		offect = 0
	}
	exdb := m.ds.DB().Where("platformId=? and deleteTime=0", platformId).Offset(offect).Limit(constant.DefaultPageSize).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *jiaoYiZhanQuService) GetJiaoYiZhanQuCount(platformId int32) (int32, error) {
	totalCount := int32(0)

	exdb := m.ds.DB().Table("t_jiaoyi_zhanqu").Where("platformId=? and deleteTime=0", platformId).Count(&totalCount)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return totalCount, exdb.Error
	}

	return totalCount, nil
}

func (m *jiaoYiZhanQuService) GetAllJiaoYiZhanQuList(platformId int32) ([]*jiaoyimodel.JiaoYiZhanQuEntity, error) {
	rst := make([]*jiaoyimodel.JiaoYiZhanQuEntity, 0)
	exdb := m.ds.DB().Where("platformId=? and deleteTime=0", platformId).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func NewJiaoYiZhanQuService(p_db db.DBService) IJiaoYiZhanQuService {
	rst := &jiaoYiZhanQuService{}
	rst.ds = p_db
	return rst
}

type contextKey string

const (
	jiaoYiZhanQuServiceKey = contextKey("jiaoYiZhanQuService")
)

func WithJiaoYiZhanQuService(ctx context.Context, ls IJiaoYiZhanQuService) context.Context {
	return context.WithValue(ctx, jiaoYiZhanQuServiceKey, ls)
}

func JiaoYiZhanQuServiceInContext(ctx context.Context) IJiaoYiZhanQuService {
	us, ok := ctx.Value(jiaoYiZhanQuServiceKey).(IJiaoYiZhanQuService)
	if !ok {
		return nil
	}
	return us
}

func SetupJiaoYiZhanQuServiceHandler(ls IJiaoYiZhanQuService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithJiaoYiZhanQuService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
