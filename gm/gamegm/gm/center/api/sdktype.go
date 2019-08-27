package api

import (
	sdktype "fgame/fgame/account/login/types"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"
	"sort"

	"github.com/xozrc/pkg/httputils"
)

type skdTypeRespon struct {
	ItemArray []*skdTypeResponItem `json:"itemArray"`
}

type skdTypeResponItem struct {
	Key  int    `json:"key"`
	Name string `json:"name"`
}

func handleSdkType(rw http.ResponseWriter, req *http.Request) {
	sdkMap := sdktype.SdkMap
	respon := &skdTypeRespon{}
	respon.ItemArray = make([]*skdTypeResponItem, 0)
	for key, value := range sdkMap {
		item := &skdTypeResponItem{}
		item.Key = int(key)
		item.Name = value
		respon.ItemArray = append(respon.ItemArray, item)
	}
	sort.Sort(skdTypeResponItemSlice(respon.ItemArray))

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}

type skdTypeResponItemSlice []*skdTypeResponItem

func (a skdTypeResponItemSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a skdTypeResponItemSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a skdTypeResponItemSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].Key > a[i].Key
}
