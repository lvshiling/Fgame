package service

import (
	"context"
	"fgame/fgame/gm/gamegm/common"
	constant "fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	gmError "fgame/fgame/gm/gamegm/error"
	centerPlatformmodel "fgame/fgame/gm/gamegm/gm/center/model"
	"fgame/fgame/gm/gamegm/pkg/timeutils"
	"fmt"
	"net/http"
	"time"

	platformmodel "fgame/fgame/gm/gamegm/gm/center/platform/model"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

type ICenterPlatformService interface {
	AddCenterPlatform(p_name string) error
	UpdateCenterPlatform(p_centerPlatformid int64, p_name string) error
	DeleteCenterPlatform(p_centerPlatformid int64) error

	GetCenterPlatformList(p_name string, p_index int) ([]*centerPlatformmodel.CenterPlatformInfo, error)
	GetAllCenterPlatformList() ([]*centerPlatformmodel.CenterPlatformInfo, error)
	GetCenterPlatformCount(p_name string) (int, error)

	GetCenterPlatformInfo(p_id int64) (*centerPlatformmodel.CenterPlatformInfo, error)

	GetPlatformIdBySdkType(p_sdkType []int) ([]int64, error)

	SavePlatformMarrySet(p_info *platformmodel.CenterPlatformMarryPriceSetInfo) error
	GetPlatformMarrySet(p_centerPlatformid int64) (*platformmodel.CenterPlatformMarryPriceSetInfo, error)

	SavePlatformSet(p_info *platformmodel.CenterPlatformSetInfo) error
	GetPlatformSet(p_centerPlatformId int64) (*platformmodel.CenterPlatformSetInfo, error)
	GetCenterPlatformSettingList(p_name string, p_index int) ([]*platformmodel.CenterPlatformSetQueryInfo, error)
	GetCenterPlatformSettingCount(p_name string) (int, error)

	AddPlatformMarryServerLog(p_info *platformmodel.PlatformMarrySendLog) error
	GetPlatformMarryServerLogList(p_centerPlatformid int64, p_flag int32, p_index int) ([]*platformmodel.PlatformMarrySendLog, error)
	GetPlatformMarryServerLogCount(p_centerPlatformid int64, p_flag int32) (int, error)
	UpdatePlatformMarryServerLogState(p_id int32, p_state int32) error
}

type centerPlatformService struct {
	db gmdb.DBService

	gmDb gmdb.DBService
}

func (m *centerPlatformService) AddCenterPlatform(p_name string) error {
	if len(p_name) == 0 {
		return gmError.GetError(gmError.ErrorCodeCenterPlatformEmpty)
	}

	exflag, err := m.existsCenterPlatform(p_name)
	if err != nil {
		return err
	}
	if exflag {
		return gmError.GetError(gmError.ErrorCodeCenterPlatformExist)
	}

	// exflag, err = m.existsCenterPlatformSdk(p_skdType)
	// if err != nil {
	// 	return err
	// }
	// if exflag {
	// 	return gmError.GetError(gmError.ErrorCodeCenterPlatformSdkExist)
	// }

	modelInfo := &centerPlatformmodel.CenterPlatformInfo{
		Name: p_name,
		// SkdType: p_skdType,
	}
	err = m.saveCenterPlatform(modelInfo)
	if err != nil {
		return err
	}
	return nil
}

func (m *centerPlatformService) UpdateCenterPlatform(p_centerPlatformid int64, p_name string) error {
	if len(p_name) == 0 {
		return gmError.GetError(gmError.ErrorCodeCenterPlatformEmpty)
	}

	exflag, err := m.existsCenterPlatformWithId(p_centerPlatformid, p_name)
	if err != nil {
		return err
	}
	if exflag {
		return gmError.GetError(gmError.ErrorCodeCenterPlatformExist)
	}

	// exflag, err = m.existsCenterPlatformSkdTypeWithId(p_centerPlatformid)
	// if err != nil {
	// 	return err
	// }
	// if exflag {
	// 	return gmError.GetError(gmError.ErrorCodeCenterPlatformSdkExist)
	// }

	modelInfo := &centerPlatformmodel.CenterPlatformInfo{
		PlatformId: p_centerPlatformid,
		Name:       p_name,
	}
	err = m.saveCenterPlatform(modelInfo)
	if err != nil {
		return err
	}
	return nil
}

func (m *centerPlatformService) DeleteCenterPlatform(p_centerPlatformid int64) error {
	now := timeutils.TimeToMillisecond(time.Now())
	errdb := m.db.DB().Table("t_platform").Where("id = ?", p_centerPlatformid).Update("deleteTime", now)
	if errdb.Error != nil {
		log.WithFields(log.Fields{
			"centerPlatformid": p_centerPlatformid,
			"error":            errdb.Error,
		}).Error("删除中心平台失败")
		return errdb.Error
	}
	return nil
}

func (m *centerPlatformService) GetCenterPlatformList(p_name string, p_index int) ([]*centerPlatformmodel.CenterPlatformInfo, error) {
	return m.getCenterPlatformList(p_name, p_index)
}

func (m *centerPlatformService) GetAllCenterPlatformList() ([]*centerPlatformmodel.CenterPlatformInfo, error) {
	return m.getAllCenterPlatformList()
}

func (m *centerPlatformService) GetCenterPlatformCount(p_name string) (int, error) {
	return m.getCenterPlatformCount(p_name)
}

func (m *centerPlatformService) GetCenterPlatformInfo(p_id int64) (*centerPlatformmodel.CenterPlatformInfo, error) {
	info := &centerPlatformmodel.CenterPlatformInfo{}
	exdb := m.db.DB().Where("id = ?", p_id).First(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return info, nil
}

func (m *centerPlatformService) GetPlatformIdBySdkType(p_sdkType []int) ([]int64, error) {
	if len(p_sdkType) == 0 {
		return nil, nil
	}
	rst := make([]*centerPlatformmodel.CenterPlatformInfo, 0)
	where := "deleteTime=0 "
	if len(p_sdkType) > 0 {
		where += fmt.Sprintf(" and sdkType IN (%s)", common.CombinIntArray(p_sdkType))
	}
	exerr := m.db.DB().Where(where).Find(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"error": exerr.Error,
		}).Error("获取skdtype中心平台列表失败")
		return nil, exerr.Error
	}
	result := make([]int64, 0)
	if len(rst) > 0 {
		for _, value := range rst {
			result = append(result, value.PlatformId)
		}
	}
	return result, nil
}

func (m *centerPlatformService) existsCenterPlatform(p_name string) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_platform").Where("name = ? and deleteTime = 0", p_name).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"centerPlatformName": p_name,
			"error":              exdb.Error,
		}).Error("查询中心平台信息失败existsCenterPlatform")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *centerPlatformService) existsCenterPlatformWithId(p_centerPlatformid int64, p_name string) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_platform").Where("name = ? and id != ? and deleteTime=0", p_name, p_centerPlatformid).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"centerPlatformName": p_name,
			"centerPlatformid":   p_centerPlatformid,
			"error":              exdb.Error,
		}).Error("查询中心平台信息失败existsCenterPlatformWithId")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *centerPlatformService) existsCenterPlatformSdk(p_skdType int) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_platform").Where("sdkType = ? and deleteTime = 0", p_skdType).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"skdType": p_skdType,
			"error":   exdb.Error,
		}).Error("查询中心平台信息失败existsCenterPlatform")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *centerPlatformService) existsCenterPlatformSkdTypeWithId(p_centerPlatformid int64) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_platform").Where(" id != ? and deleteTime=0", p_centerPlatformid).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{

			"error": exdb.Error,
		}).Error("查询中心平台信息失败existsCenterPlatformWithId")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *centerPlatformService) saveCenterPlatform(p_info *centerPlatformmodel.CenterPlatformInfo) error {
	exdb := m.db.DB().Save(p_info)
	now := timeutils.TimeToMillisecond(time.Now())
	if p_info.PlatformId > 0 {
		p_info.CreateTime = now
	} else {
		p_info.UpdateTime = now
	}
	if exdb.Error != nil {
		log.WithFields(log.Fields{
			"centerPlatformName": p_info.Name,
			"centerPlatformid":   p_info.PlatformId,
			"error":              exdb.Error,
		}).Error("保存中心平台失败")
		return exdb.Error
	}
	return nil
}

func (m *centerPlatformService) getCenterPlatformList(p_name string, p_index int) ([]*centerPlatformmodel.CenterPlatformInfo, error) {
	offset := (p_index - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize
	rst := make([]*centerPlatformmodel.CenterPlatformInfo, 0)

	exerr := m.db.DB().Where("deleteTime =0 and name like ?", "%"+p_name+"%").Offset(offset).Limit(limit).Find(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"centerPlatformName": p_name,
			"error":              exerr.Error,
		}).Error("获取中心平台列表失败")
		return nil, exerr.Error
	}
	return rst, nil
}

func (m *centerPlatformService) getAllCenterPlatformList() ([]*centerPlatformmodel.CenterPlatformInfo, error) {

	rst := make([]*centerPlatformmodel.CenterPlatformInfo, 0)

	exerr := m.db.DB().Where("deleteTime=0").Find(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"error": exerr.Error,
		}).Error("获取ALL中心平台列表失败")
		return nil, exerr.Error
	}
	return rst, nil
}

func (m *centerPlatformService) getCenterPlatformCount(p_name string) (int, error) {
	rst := 0

	exerr := m.db.DB().Table("t_platform").Where("deleteTime =0 and name like ?", "%"+p_name+"%").Count(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"centerPlatformName": p_name,
			"error":              exerr.Error,
		}).Error("获取中心平台列表失败")
		return 0, exerr.Error
	}
	return rst, nil
}

func (m *centerPlatformService) SavePlatformMarrySet(p_info *platformmodel.CenterPlatformMarryPriceSetInfo) error {
	exdb := m.db.DB().Save(p_info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	return nil
}

func (m *centerPlatformService) GetPlatformMarrySet(p_centerPlatformid int64) (*platformmodel.CenterPlatformMarryPriceSetInfo, error) {
	rst := &platformmodel.CenterPlatformMarryPriceSetInfo{}

	exdb := m.db.DB().Where("platformId = ?", p_centerPlatformid).First(rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}

	return rst, nil
}

func (m *centerPlatformService) SavePlatformSet(p_info *platformmodel.CenterPlatformSetInfo) error {
	if p_info.Id == 0 {
		dbinfo := &platformmodel.CenterPlatformSetInfo{}
		exdb := m.db.DB().Where("platformId=? and deleteTime=0", p_info.PlatformId).First(dbinfo)
		if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
			return exdb.Error
		}
		if dbinfo != nil && dbinfo.Id > 0 {
			p_info.Id = dbinfo.Id
			p_info.CreateTime = dbinfo.CreateTime
		}
	}
	exdb := m.db.DB().Save(p_info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	return nil
}

func (m *centerPlatformService) GetPlatformSet(p_centerPlatformId int64) (*platformmodel.CenterPlatformSetInfo, error) {
	rst := &platformmodel.CenterPlatformSetInfo{}

	exdb := m.db.DB().Where("platformId = ?", p_centerPlatformId).First(rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}

	return rst, nil
}

var (
	centerPlatformSettingListSql = `SELECT
	A.id
	,A.createTime
	,A.updateTime
	,A.deleteTime
	,A.settingContent
	,B.id AS platformId
	,B.name AS platformName
FROM
	t_platform B
	LEFT JOIN t_platform_setting A
	ON A.platformId = B.id AND A.deleteTime = 0
WHERE
	B.deleteTime = 0 AND B.name like ?`
)

func (m *centerPlatformService) GetCenterPlatformSettingList(p_name string, p_index int) ([]*platformmodel.CenterPlatformSetQueryInfo, error) {
	offset := (p_index - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize
	rst := make([]*platformmodel.CenterPlatformSetQueryInfo, 0)
	exerr := m.db.DB().Raw(centerPlatformSettingListSql, "%"+p_name+"%").Offset(offset).Limit(limit).Scan(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"centerPlatformName": p_name,
			"error":              exerr.Error,
		}).Error("获取中心平台列表失败")
		return nil, exerr.Error
	}
	return rst, nil
}

func (m *centerPlatformService) AddPlatformMarryServerLog(p_info *platformmodel.PlatformMarrySendLog) error {
	exdb := m.gmDb.DB().Save(p_info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}

	return nil
}

func (m *centerPlatformService) GetPlatformMarryServerLogList(p_centerPlatformid int64, p_flag int32, p_index int) ([]*platformmodel.PlatformMarrySendLog, error) {
	rst := make([]*platformmodel.PlatformMarrySendLog, 0)
	offset := (p_index - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize
	where := " 1=1"
	if p_centerPlatformid > 0 {
		where += fmt.Sprintf(" and platformId = %d", p_centerPlatformid)
	}
	if p_flag > -1 {
		where += fmt.Sprintf(" and successFlag = %d", p_flag)
	}
	exdb := m.gmDb.DB().Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {

		return nil, exdb.Error
	}
	return rst, nil
}

func (m *centerPlatformService) GetPlatformMarryServerLogCount(p_centerPlatformid int64, p_flag int32) (int, error) {
	rst := 0
	where := " 1=1"
	if p_centerPlatformid > 0 {
		where += fmt.Sprintf(" and platformId = %d", p_centerPlatformid)
	}
	if p_flag > -1 {
		where += fmt.Sprintf(" and successFlag = %d", p_flag)
	}
	exdb := m.gmDb.DB().Table("t_marry_set_log").Where(where).Count(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {

		return 0, exdb.Error
	}
	return rst, nil
}

func (m *centerPlatformService) UpdatePlatformMarryServerLogState(p_id int32, p_state int32) error {
	exdb := m.gmDb.DB().Table("t_marry_set_log").Where("id = ?", p_id).Update("successFlag", p_state)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	return nil
}

func NewCenterPlatformService(p_db gmdb.DBService, p_gmdb gmdb.DBService) ICenterPlatformService {
	rst := &centerPlatformService{
		db:   p_db,
		gmDb: p_gmdb,
	}
	return rst
}

var (
	centerPlatformSettingCountSql = `SELECT
	COUNT(1) AS myRow
FROM
	t_platform B
WHERE
	 B.deleteTime = 0 AND B.name like ?`
)

func (m *centerPlatformService) GetCenterPlatformSettingCount(p_name string) (int, error) {
	rst := 0
	exerr := m.db.DB().Raw(centerPlatformSettingCountSql, "%"+p_name+"%").Count(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"centerPlatformName": p_name,
			"error":              exerr.Error,
		}).Error("获取中心平台列表失败")
		return 0, exerr.Error
	}
	return rst, nil
}

type contextKey string

const (
	centerPlatformServiceKey = contextKey("CenterPlatformService")
)

func WithCenterPlatformService(ctx context.Context, ls ICenterPlatformService) context.Context {
	return context.WithValue(ctx, centerPlatformServiceKey, ls)
}

func CenterPlatformServiceInContext(ctx context.Context) ICenterPlatformService {
	us, ok := ctx.Value(centerPlatformServiceKey).(ICenterPlatformService)
	if !ok {
		return nil
	}
	return us
}

func SetupCenterPlatformServiceHandler(ls ICenterPlatformService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithCenterPlatformService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
