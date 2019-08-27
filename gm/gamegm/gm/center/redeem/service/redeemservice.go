package service

import (
	"context"
	"fgame/fgame/gm/gamegm/common"
	"fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	"fgame/fgame/gm/gamegm/gm/center/redeem/model"
	"fgame/fgame/gm/gamegm/gm/center/redeem/pbmodel"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/satori/go.uuid"

	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

type IRedeemService interface {
	AddRedeem(p_info *pbmodel.RedeemInfo, p_centerPlatList []int64) error
	UpdateRedeem(p_info *pbmodel.RedeemInfo, p_centerPlatList []int64) error
	GetRedeemList(p_name string, p_sdkType int, p_pageIndex int, p_centerPlatList []int64) ([]*pbmodel.RedeemInfo, error)
	GetRedeemCount(p_name string, p_sdkType int, p_centerPlatList []int64) (int, error)
	DeleteRedeem(p_id int) error
	NewRedeemCode(p_id int) error

	GetRedeemCodeList(p_id int) ([]*pbmodel.RedeemCodeInfo, error)
	GetRedeemCodeCount(p_id int) (int, error)
}

type redeemService struct {
	db gmdb.DBService
}

func (m *redeemService) AddRedeem(p_info *pbmodel.RedeemInfo, p_centerPlatList []int64) error {
	now := timeutils.TimeToMillisecond(time.Now())
	info := changePuToDbRedeem(p_info)
	info.CreateTime = now
	exdb := m.db.DB().Begin()

	mydb := exdb.Save(info)
	if mydb.Error != nil {
		exdb.Rollback()
		return mydb.Error
	}
	if len(p_centerPlatList) > 0 {
		for _, value := range p_centerPlatList {
			plModel := &model.RedeemPlatform{
				PlatformId: value,
				RedeemId:   info.Id,
				CreateTime: now,
			}
			mydb = exdb.Save(plModel)
			if mydb.Error != nil {
				exdb.Rollback()
				return mydb.Error
			}
		}
	}
	exdb.Commit()
	return nil
}

func (m *redeemService) UpdateRedeem(p_info *pbmodel.RedeemInfo, p_centerPlatList []int64) error {
	now := timeutils.TimeToMillisecond(time.Now())
	info := changePuToDbRedeem(p_info)
	exdb := m.db.DB().Begin()
	mydb := exdb.Save(info)
	if mydb.Error != nil {
		exdb.Rollback()
		return mydb.Error
	}
	mydb = exdb.Table("t_redeem_platform").Where("redeemId=?", info.Id).Update("deleteTime", now)
	if mydb.Error != nil {
		exdb.Rollback()
		return mydb.Error
	}
	if len(p_centerPlatList) > 0 {
		for _, value := range p_centerPlatList {
			plModel := &model.RedeemPlatform{
				PlatformId: value,
				RedeemId:   info.Id,
				CreateTime: now,
			}
			mydb = exdb.Save(plModel)
			if mydb.Error != nil {
				exdb.Rollback()
				return mydb.Error
			}
		}
	}
	exdb.Commit()
	return nil
}

func (m *redeemService) GetRedeemList(p_name string, p_sdkType int, p_pageindex int, p_centerPlatList []int64) ([]*pbmodel.RedeemInfo, error) {
	rst := make([]*model.Redeem, 0)
	offset := (p_pageindex - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize

	where := "deleteTime = 0"
	if len(p_name) > 0 {
		where += fmt.Sprintf(" and giftBagName LIKE '%s'", "%"+p_name+"%")
	}
	if p_sdkType > 0 {
		where += fmt.Sprintf(" and CONCAT(',',sdkTypes,',') LIKE '%s'", "%"+strconv.Itoa(p_sdkType)+"%")
	}
	if len(p_centerPlatList) > 0 {
		where += fmt.Sprintf(" and EXISTS(SELECT 1 FROM t_redeem_platform WHERE t_redeem_platform.redeemId=t_redeem.id and platformId IN (%s)) ", common.CombinInt64Array(p_centerPlatList))
	}

	exdb := m.db.DB().Where(where).Offset(offset).Limit(limit).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}

	result := make([]*pbmodel.RedeemInfo, 0)
	for _, value := range rst {
		item := changeDbToPuRedeem(value)
		result = append(result, item)
	}

	return result, nil
}

func (m *redeemService) GetRedeemCount(p_name string, p_sdkType int, p_centerPlatList []int64) (int, error) {
	rst := 0

	where := "deleteTime = 0"
	if len(p_name) > 0 {
		where += fmt.Sprintf(" and giftBagName LIKE '%s'", "%"+p_name+"%")
	}
	if p_sdkType > 0 {
		where += fmt.Sprintf(" and CONCAT(',',sdkTypes,',') LIKE '%s'", "%"+strconv.Itoa(p_sdkType)+"%")
	}
	if len(p_centerPlatList) > 0 {
		where += fmt.Sprintf(" and EXISTS(SELECT 1 FROM t_redeem_platform WHERE t_redeem_platform.redeemId=t_redeem.id and platformId IN (%s)) ", common.CombinInt64Array(p_centerPlatList))
	}

	exdb := m.db.DB().Table("t_redeem").Where(where).Count(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return 0, exdb.Error
	}
	return rst, nil
}

func (m *redeemService) DeleteRedeem(p_id int) error {
	now := timeutils.TimeToMillisecond(time.Now())
	exdb := m.db.DB().Begin()
	erdb := exdb.Table("t_redeem").Where("id = ?", p_id).Update("deleteTime", now)
	if erdb.Error != nil {
		exdb.Rollback()
		return erdb.Error
	}
	erdb = exdb.Table("t_redeem_code").Where("redeemId = ?", p_id).Update("deleteTime", now)
	if erdb.Error != nil {
		exdb.Rollback()
		return erdb.Error
	}
	exdb.Commit()
	return nil
}

func (m *redeemService) NewRedeemCode(p_id int) error {
	now := timeutils.TimeToMillisecond(time.Now())
	info := &model.Redeem{}
	erdb := m.db.DB().Where("id = ?", p_id).First(info)
	if erdb.Error != nil {
		return erdb.Error
	}
	if info.CreateFlag > 0 {
		return fmt.Errorf("该礼包已发放")
	}
	exdb := m.db.DB().Begin()
	erdb = exdb.Table("t_redeem").Where("id = ?", p_id).Update("createFlag", 1)
	if erdb.Error != nil {
		exdb.Rollback()
		return erdb.Error
	}
	codeList := m.newRedeemCode(info.RedeemNum)
	if len(codeList) > 0 {
		for _, value := range codeList {
			item := &model.RedeemCode{
				RedeemCode: value,
				RedeemId:   p_id,
				CreateTime: now,
			}
			erdb = exdb.Save(item)
			if erdb.Error != nil {
				exdb.Rollback()
				return erdb.Error
			}
		}
	}
	exdb.Commit()
	return nil
}

func (m *redeemService) GetRedeemCodeList(p_id int) ([]*pbmodel.RedeemCodeInfo, error) {
	rst := make([]*model.RedeemCode, 0)
	exdb := m.db.DB().Where("redeemId = ? and deleteTime = 0", p_id).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	result := make([]*pbmodel.RedeemCodeInfo, 0)
	if len(rst) > 0 {
		for _, value := range rst {
			item := &pbmodel.RedeemCodeInfo{
				Id:         value.Id,
				RedeemCode: value.RedeemCode,
				RedeemId:   value.RedeemId,
				UseNum:     value.UseNum,
			}
			result = append(result, item)
		}
	}
	return result, nil
}
func (m *redeemService) GetRedeemCodeCount(p_id int) (int, error) {
	rst := 0
	exdb := m.db.DB().Table("t_redeem_code").Where("redeemId = ? and deleteTime = 0", p_id).Count(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return 0, exdb.Error
	}
	return rst, nil
}

func (m *redeemService) newRedeemCode(p_number int) []string {
	rst := make([]string, 0)
	for i := 0; i < p_number; i++ {
		item := strings.Replace(uuid.NewV4().String(), "-", "", -1)
		rst = append(rst, item)
	}
	return rst
}

func NewRedeemService(p_db gmdb.DBService) IRedeemService {
	rst := &redeemService{
		db: p_db,
	}
	return rst
}

type contextKey string

const (
	redeemServiceKey = contextKey("RedeemService")
)

func WithRedeemService(ctx context.Context, ls IRedeemService) context.Context {
	return context.WithValue(ctx, redeemServiceKey, ls)
}

func RedeemServiceInContext(ctx context.Context) IRedeemService {
	us, ok := ctx.Value(redeemServiceKey).(IRedeemService)
	if !ok {
		return nil
	}
	return us
}

func SetupRedeemServiceHandler(ls IRedeemService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithRedeemService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}

func changePuToDbRedeem(p_info *pbmodel.RedeemInfo) *model.Redeem {
	sdkTypes := ""
	if len(p_info.SdkTypes) > 0 {
		for i, value := range p_info.SdkTypes {
			if i > 0 {
				sdkTypes += ","
			}
			sdkTypes += strconv.Itoa(value)
		}
	}
	rst := &model.Redeem{
		Id:                 p_info.Id,
		GiftBagName:        p_info.GiftBagName,
		GiftBagDesc:        p_info.GiftBagDesc,
		GiftBagContent:     p_info.GiftBagContent,
		RedeemNum:          p_info.RedeemNum,
		RedeemUseNum:       p_info.RedeemUseNum,
		RedeemPlayerUseNum: p_info.RedeemPlayerUseNum,
		RedeemServerUseNum: p_info.RedeemServerUseNum,
		SdkTypes:           sdkTypes,
		SendType:           p_info.SendType,
		StartTime:          p_info.StartTime,
		EndTime:            p_info.EndTime,
		MinPlayerLevel:     p_info.MinPlayerLevel,
		MinVipLevel:        p_info.MinVipLevel,
		CreateFlag:         p_info.CreateFlag,
	}
	return rst
}

func changeDbToPuRedeem(p_info *model.Redeem) *pbmodel.RedeemInfo {
	sdkArray := make([]int, 0)
	if len(p_info.SdkTypes) > 0 {
		strArray := strings.Split(p_info.SdkTypes, ",")
		if len(strArray) > 0 {
			for _, value := range strArray {
				item, err := strconv.Atoi(value)
				if err != nil {
					continue
				}
				sdkArray = append(sdkArray, item)
			}
		}
	}
	rst := &pbmodel.RedeemInfo{
		Id:                 p_info.Id,
		GiftBagName:        p_info.GiftBagName,
		GiftBagDesc:        p_info.GiftBagDesc,
		GiftBagContent:     p_info.GiftBagContent,
		RedeemNum:          p_info.RedeemNum,
		RedeemUseNum:       p_info.RedeemUseNum,
		RedeemPlayerUseNum: p_info.RedeemPlayerUseNum,
		RedeemServerUseNum: p_info.RedeemServerUseNum,
		SdkTypes:           sdkArray,
		SendType:           p_info.SendType,
		StartTime:          p_info.StartTime,
		EndTime:            p_info.EndTime,
		MinPlayerLevel:     p_info.MinPlayerLevel,
		MinVipLevel:        p_info.MinVipLevel,
		CreateFlag:         p_info.CreateFlag,
	}
	return rst
}
