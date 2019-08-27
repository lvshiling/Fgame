package middleware

import (
	"crypto/md5"
	"encoding/hex"
	plservice "fgame/fgame/gm/gamegm/gm/platform/service"
	"net/http"
	"sort"
	"strconv"
	"strings"

	paramapi "fgame/fgame/gm/gamegm/gm/openapi/param"

	gmerr "fgame/fgame/gm/gamegm/error"
	errhttp "fgame/fgame/gm/gamegm/error/utils"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/xozrc/pkg/httputils"
)

const (
	signKeyName     = "sign"
	platformKeyName = "gameId"
)

func OpenApiHandlerMiddleware() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		if !strings.HasPrefix(req.URL.Path, "/api/openapi") {
			hf.ServeHTTP(rw, req)
			return
		}
		paramMap := make(map[string]string)

		log.WithFields(log.Fields{
			"paramMap": paramMap,
		}).Infof("OpenApiHandlerMiddleware:参数%#v", paramMap)

		err := httputils.Bind(req, &paramMap)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("OpenApiHandlerMiddleware:解析错误")
			err = gmerr.GetError(gmerr.ErrorCodeOpenApiParamBindError)
			errhttp.ResponseWithError(rw, err)
			return
		}
		platformIdString, exists := paramMap[platformKeyName]
		if !exists {
			log.WithFields(log.Fields{
				"platformKey": platformKeyName,
			}).Warn("OpenApiHandlerMiddleware:platformKey错误")

			errhttp.ResponseWithError(rw, gmerr.GetError(gmerr.ErrorCodeOpenApiUserNotExists))
			return
		}
		requestSign, exists := paramMap[signKeyName]
		if !exists {
			log.WithFields(log.Fields{
				"signKeyName": signKeyName,
			}).Warn("OpenApiHandlerMiddleware:signKey错误")

			errhttp.ResponseWithError(rw, gmerr.GetError(gmerr.ErrorCodeOpenApiSign))
			return
		}
		platformID, err := strconv.Atoi(platformIdString)
		if err != nil || platformID < 1 {
			log.WithFields(log.Fields{
				"platformIdString": platformIdString,
			}).Warn("OpenApiHandlerMiddleware:platformIdString错误")

			errhttp.ResponseWithError(rw, gmerr.GetError(gmerr.ErrorCodeOpenApiUserNotExists))
			return
		}
		plService := plservice.PlatformServiceInContext(req.Context())
		if plService == nil {
			log.WithFields(log.Fields{}).Error("OpenApiHandlerMiddleware:platformService is nil")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		plInfo, err := plService.GetPlatformInfo(int64(platformID))
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("OpenApiHandlerMiddleware:获取平台信息异常")
			errhttp.ResponseWithError(rw, gmerr.GetError(gmerr.ErrorCodeOpenApiUserNotExists))
			return
		}
		if plInfo == nil || plInfo.PlatformID < 1 || len(plInfo.SignKey) == 0 {
			log.WithFields(log.Fields{
				"error": err,
			}).Warn("OpenApiHandlerMiddleware:平台为空或者签名key为空")
			errhttp.ResponseWithError(rw, gmerr.GetError(gmerr.ErrorCodeOpenApiUserNotExists))
			return
		}
		platformSignKey := plInfo.SignKey
		requestParam := ""

		keyArray := make([]string, 0)
		for value, _ := range paramMap {
			if value == signKeyName {
				continue
			}
			keyArray = append(keyArray, value)
		}
		sort.Sort(sort.StringSlice(keyArray))
		for index, value := range keyArray {
			if index > 0 {
				requestParam += "&"
			}
			requestParam += value + "=" + paramMap[value]
		}
		if len(requestParam) > 0 {
			requestParam += "&"
		}
		requestParam += "key=" + platformSignKey

		curSign := Md5Encode(requestParam)

		log.WithFields(log.Fields{
			"requestParam":      requestParam,
			"requestSign":       requestSign,
			"requestParamAfter": curSign,
		}).Debug("OpenApiHandlerMiddleware:加密前字符串")

		if curSign != requestSign {
			log.WithFields(log.Fields{
				"requestParam":      requestParam,
				"requestSign":       requestSign,
				"requestParamAfter": curSign,
			}).Warn("OpenApiHandlerMiddleware:签名异常")
			errhttp.ResponseWithError(rw, gmerr.GetError(gmerr.ErrorCodeOpenApiSign))
			return
		}
		ctx := plservice.WithApiPlatformId(req.Context(), int64(platformID))
		ctx = plservice.WithApiPlatformInfo(req.Context(), plInfo)
		ctx = paramapi.WithApiData(ctx, paramMap)
		nreq := req.WithContext(ctx)
		hf.ServeHTTP(rw, nreq)
		return
	})
}

func Md5Encode(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
