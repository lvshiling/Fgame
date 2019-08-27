package api

import (
	"bytes"
	"fgame/fgame/gm/gamegm/gm/center/redeem/pbmodel"
	"fgame/fgame/gm/gamegm/gm/center/redeem/service"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/tealeg/xlsx"
	"github.com/xozrc/pkg/httputils"
)

type redeemCodeListExportRequest struct {
	Id int `json:"id"`
}

type redeemCodeListExportRespon struct {
	ItemArray  []*pbmodel.RedeemCodeInfo `json:"itemArray"`
	TotalCount int                       `json:"total"`
}

func handleRedeemCodeListExport(rw http.ResponseWriter, req *http.Request) {
	form := &redeemCodeListExportRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("兑换码列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rds := service.RedeemServiceInContext(req.Context())
	if rds == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("兑换码列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := rds.GetRedeemCodeList(form.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("兑换码列表，添加异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	heads := []string{"序号", "兑换码", "已使用次数"}
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("sheet")
	row := sheet.AddRow()
	for _, value := range heads {
		row.AddCell().Value = value
	}
	for index, rowValue := range rst {
		newRow := sheet.AddRow()
		newRow.AddCell().Value = strconv.Itoa(index)
		newRow.AddCell().Value = rowValue.RedeemCode
		newRow.AddCell().Value = strconv.Itoa(rowValue.UseNum)
	}
	var buffer bytes.Buffer
	if err := file.Write(&buffer); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("兑换码列表，写入缓存失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	// r := bytes.NewReader(buffer.Bytes())

	rw.Header().Add("Content-Disposition", "attachment")
	rw.Header().Add("Content-Type", "application/vnd.ms-excel")
	rw.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	rw.Write(buffer.Bytes())
	// http.ServeContent(rw, req, "兑换码.xlsx", time.Now(), r)
}
