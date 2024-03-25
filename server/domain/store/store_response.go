package order_domain

import "time"

type StoreResponse struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`

	StoreID         string        `json:"storeId"`
	ChildStoreIds   []interface{} `json:"childStoreIds"`
	ShipFromAddress struct {
		Zip         string `json:"zip"`
		City        string `json:"city"`
		County      string `json:"county"`
		Address1    string `json:"address1"`
		CountryCode string `json:"countryCode"`
		Country     string `json:"country"`
	} `json:"shipFromAddress"`
	DefaultUserID          string        `json:"defaultUserId"`
	DefaultAccountID       string        `json:"defaultAccountId"`
	DefaultCarrierID       string        `json:"defaultCarrierId"`
	Name                   string        `json:"name"`
	Region                 string        `json:"region"`
	ExpressType            string        `json:"expressType"`
	Domain                 string        `json:"domain"`
	Platform               string        `json:"platform"`
	MerchantStoreURL       string        `json:"merchantStoreUrl"`
	AppInstallLink         interface{}   `json:"appInstallLink"`
	RequireInstallHelp     bool          `json:"requireInstallHelp"`
	IsSmkStore             bool          `json:"isSmkStore"`
	IsTestStore            bool          `json:"isTestStore"`
	IsStoreActive          bool          `json:"isStoreActive"`
	ProductAPI             string        `json:"productApi"`
	OrderAPI               string        `json:"orderApi"`
	MiniProgramAppID       string        `json:"miniProgramAppId"`
	MiniProgramAppPath     string        `json:"miniProgramAppPath"`
	CampaignTrackingParams []interface{} `json:"campaignTrackingParams"`
	BaseCurrency           string        `json:"baseCurrency"`
	Meta                   struct {
		Lang        string        `json:"lang"`
		Title       string        `json:"title"`
		Description string        `json:"description"`
		HeaderTitle string        `json:"headerTitle"`
		ImageLogo   interface{}   `json:"imageLogo"`
		ImageIcon   interface{}   `json:"imageIcon"`
		Theme       string        `json:"theme"`
		ThemeImages []interface{} `json:"themeImages"`
		Banners     []interface{} `json:"banners"`
		BrandLogos  []interface{} `json:"brandLogos"`
	} `json:"meta"`
	Features struct {
		UseSMSVerification          bool `json:"useSMSVerification"`
		UseSMEProducts              bool `json:"useSMEProducts"`
		UseStoreShippingRate        bool `json:"useStoreShippingRate"`
		HasCategoryBar              bool `json:"hasCategoryBar"`
		HasHomeBanner               bool `json:"hasHomeBanner"`
		ShowHint                    bool `json:"showHint"`
		UseMerchantLayout           bool `json:"useMerchantLayout"`
		RequiresLogin               bool `json:"requiresLogin"`
		RemoveLocalSalesTax         bool `json:"removeLocalSalesTax"`
		AddCbecTax                  bool `json:"addCbecTax"`
		MiniProgramEnabled          bool `json:"miniProgramEnabled"`
		WithCheckoutRoutes          bool `json:"withCheckoutRoutes"`
		SendOrderSMS                bool `json:"sendOrderSMS"`
		UseCartTax                  bool `json:"useCartTax"`
		CartProductPriceIncludesTax bool `json:"cartProductPriceIncludesTax"`
		CalculateStock              bool `json:"calculateStock"`
		ShopifyServiceVersion       int  `json:"shopifyServiceVersion"`
		EmailEnabled                bool `json:"emailEnabled"`
		CombineOrders               bool `json:"combineOrders"`
	} `json:"features"`
	CreatedAt time.Time `json:"createdAt"`
	Carriers  []struct {
		CarrierID     string `json:"carrierId"`
		ManualWaybill bool   `json:"manualWaybill"`
		ExpressType   string `json:"expressType"`
		CustomsPort   string `json:"customsPort"`
		Type          string `json:"type"`
		ShippingRates []struct {
			RateID       string        `json:"rateId"`
			Type         string        `json:"type"`
			Currency     string        `json:"currency"`
			Value        int           `json:"value"`
			WeightLimits []interface{} `json:"weightLimits"`
			MinimumSpend int           `json:"minimumSpend"`
		} `json:"shippingRates"`
	} `json:"carriers"`
}
