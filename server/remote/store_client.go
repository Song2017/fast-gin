package order_remote

import (
	"log"
	pkg "server/_pkg"
	conf "server/server/conf"
)

func GetStore(StoreId string) []byte {
	ffResp, ffBodyStr, errs := pkg.NewRequest().
		Get(conf.O_PROJECT_CONFIG.StoreService.BaseURL + "/stores/" + StoreId).
		Timeout(pkg.GetTimeOutSeconds()).
		EndBytes()
	log.Println(ffResp, ffBodyStr, errs)
	return ffBodyStr
}
