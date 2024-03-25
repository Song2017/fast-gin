package order_domain

import (
	"context"
	"encoding/json"
	"log"
	initial "server/_init"
	client "server/server/remote"
)

type StoreMeta struct {
	Status int `json:"status,omitempty"`

	StoreID                     string          `json:"storeId,omitempty"`
	StoreName                   string          `json:"storeName,omitempty"`
	SendOrderNotification       bool            `json:"sendOrderNotification,omitempty"`
	Carrier                     string          `json:"carrier,omitempty"`
	ExpressCompany              string          `json:"expressCompany,omitempty"`
	ExpressType                 string          `json:"expressType,omitempty"`
	PaymentMerchant             string          `json:"paymentMerchant,omitempty"`
	Customs                     string          `json:"customs,omitempty"`
	IncludeTax                  bool            `json:"includeTax,omitempty"`
	Platform                    string          `json:"platform,omitempty"`
	UseCartTax                  bool            `json:"useCartTax,omitempty"`
	CartProductPriceIncludesTax bool            `json:"cartProductPriceIncludesTax,omitempty"`
	ShipFromAddress             ShipFromAddress `json:"shipFromAddress,omitempty"`
	ShippingRate                ShippingRate    `json:"shippingRate,omitempty"`
	BaseCurrency                string          `json:"baseCurrency,omitempty"`
	CheckStock                  bool            `json:"checkStock,omitempty"`
	ReduceStock                 bool            `json:"reduceStock,omitempty"`
	ShopifyServiceVersion       int             `json:"shopifyServiceVersion,omitempty"`
	Region                      string          `json:"region,omitempty"`
	ManualWaybill               bool            `json:"manualWaybill,omitempty"`
	EmailEnabled                bool            `json:"emailEnabled,omitempty"`
	CombineOrders               bool            `json:"combineOrders,omitempty"`
}

var STORE_CACHE_KEY = "Store:Meta:"

func GetStore(storeId string) *StoreMeta {
	// Todo: add local cache
	count := ExistStore(storeId)
	store := &StoreMeta{StoreID: storeId}
	if count == 0 {
		return RefreshStore(storeId)
	} else {
		var err error
		if storeBytes, err := initial.O_REDIS.Get(
			context.Background(), STORE_CACHE_KEY+storeId).Bytes(); err == nil {
			if err = json.Unmarshal(storeBytes, &store); err == nil {
				return store
			}
		}
		log.Println(err)
	}
	return nil
}

func ExistStore(storeId string) int {
	count, err := initial.O_REDIS.Exists(
		context.Background(), STORE_CACHE_KEY+storeId).Result()
	if err != nil {
		log.Println(err)
		return 0
	}
	return int(count)
}

func RefreshStore(storeId string) *StoreMeta {
	resp := &StoreResponse{StoreID: storeId}
	bodyBytes := client.GetStore(storeId)
	var err error
	err = json.Unmarshal(bodyBytes, &resp)
	if err != nil || resp.Status == 404 {
		log.Println(resp.Status, err)
		return nil
	}

	store := &StoreMeta{StoreID: storeId}
	json.Unmarshal(bodyBytes, &store)
	store.ManualWaybill = resp.Carriers[0].ManualWaybill
	store.Carrier = resp.Carriers[0].CarrierID
	store.ExpressCompany = resp.Carriers[0].Type
	store.Customs = resp.Carriers[0].CustomsPort
	store.EmailEnabled = resp.Features.EmailEnabled
	store.CombineOrders = resp.Features.CombineOrders

	// cache store meta
	storeStr, err := json.Marshal(store)
	if err == nil {
		err = initial.O_REDIS.Set(context.Background(),
			STORE_CACHE_KEY+storeId, string(storeStr), 0).Err()
	}
	if err != nil {
		log.Println("RedisStoreSetError!", err)
		return nil
	}

	return store
}
