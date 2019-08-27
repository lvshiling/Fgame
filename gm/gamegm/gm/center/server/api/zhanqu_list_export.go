package api

import (
	"bytes"
	gmcenterServer "fgame/fgame/gm/gamegm/gm/center/server/service"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/tealeg/xlsx"
	"github.com/xozrc/pkg/httputils"
)

func handleCenterServerZhanQuListExport(rw http.ResponseWriter, req *http.Request) {

	form := &centerServerZhanQuRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("handleCenterServerZhanQuList，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmcenterServer.CenterServerServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("handleCenterServerZhanQuList，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	zhanQuList, err := service.GetCenterServerZhanQu(form.PlatformId)
	if err != nil {
		log.WithFields(log.Fields{
			"platformId": form.PlatformId,
			"error":      err,
		}).Error("handleCenterServerZhanQuList，获取战区服务器异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	zhanQuServerId := make([]int, 0)
	for _, value := range zhanQuList {
		zhanQuServerId = append(zhanQuServerId, value.ServerId)
	}

	childServer, err := service.GetZhanQuServer(form.PlatformId, zhanQuServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"platformId": form.PlatformId,
			"zhanQuList": zhanQuServerId,
			"error":      err,
		}).Error("handleCenterServerZhanQuList，子服务器异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := &centerServerZhanQuRespon{}
	respon.ItemArray = make([]*centerServerZhanQuResponItem, 0)
	for _, value := range zhanQuList {
		item := &centerServerZhanQuResponItem{
			ZhanQuServerId:   int32(value.ServerId),
			ZhanQuServerName: value.ServerName,
		}
		for _, childValue := range childServer {
			if childValue.ParentServerId != value.ServerId {
				continue
			}
			if len(item.ZhanQuChild) > 0 {
				item.ZhanQuChild += ","
			}
			item.ZhanQuChild += strconv.Itoa(childValue.ServerId)
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}

	heads := []string{"战区服务器序号", "战区服务器名", "战区下服务器序号"}
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("sheet")
	row := sheet.AddRow()
	for _, value := range heads {
		row.AddCell().Value = value
	}
	for _, rowValue := range respon.ItemArray {
		newRow := sheet.AddRow()
		newRow.AddCell().Value = strconv.Itoa(int(rowValue.ZhanQuServerId))
		newRow.AddCell().Value = rowValue.ZhanQuServerName
		newRow.AddCell().Value = rowValue.ZhanQuChild
	}
	var buffer bytes.Buffer
	if err := file.Write(&buffer); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("handleCenterServerZhanQuList，写入缓存失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	// r := bytes.NewReader(buffer.Bytes())

	rw.Header().Add("Content-Disposition", "attachment")
	rw.Header().Add("Content-Type", "application/vnd.ms-excel")
	rw.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	rw.Write(buffer.Bytes())
}
