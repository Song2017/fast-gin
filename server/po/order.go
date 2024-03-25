package order_dao

import (
	"time"
)

func (f *OrderPO) TableName() string {
	return "order"
}

// Order is an object representing the database table.
type OrderPO struct {
	ID                       string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	CartID                   string    `boil:"cart_id" json:"cart_id,omitempty" toml:"cart_id" yaml:"cart_id,omitempty"`
	OuterOrderID             string    `boil:"outer_order_id" json:"outer_order_id,omitempty" toml:"outer_order_id" yaml:"outer_order_id,omitempty"`
	UserID                   string    `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`
	OuterUserID              string    `boil:"outer_user_id" json:"outer_user_id,omitempty" toml:"outer_user_id" yaml:"outer_user_id,omitempty"`
	StoreID                  string    `boil:"store_id" json:"store_id,omitempty" toml:"store_id" yaml:"store_id,omitempty"`
	Currency                 string    `boil:"currency" json:"currency,omitempty" toml:"currency" yaml:"currency,omitempty"`
	Total                    int       `boil:"total" json:"total" toml:"total" yaml:"total"`
	Discount                 int       `boil:"discount" json:"discount" toml:"discount" yaml:"discount"`
	Tax                      int       `boil:"tax" json:"tax,omitempty" toml:"tax" yaml:"tax,omitempty"`
	ShippingFee              int       `boil:"shipping_fee" json:"shipping_fee,omitempty" toml:"shipping_fee" yaml:"shipping_fee,omitempty"`
	DeliveryUserName         string    `boil:"delivery_user_name" json:"delivery_user_name,omitempty" toml:"delivery_user_name" yaml:"delivery_user_name,omitempty"`
	DeliveryAddress          string    `boil:"delivery_address" json:"delivery_address,omitempty" toml:"delivery_address" yaml:"delivery_address,omitempty"`
	DeliveryPhone            string    `boil:"delivery_phone" json:"delivery_phone,omitempty" toml:"delivery_phone" yaml:"delivery_phone,omitempty"`
	Status                   int       `boil:"status" json:"status" toml:"status" yaml:"status"`
	CreateTimestamp          time.Time `boil:"create_timestamp" json:"create_timestamp" toml:"create_timestamp" yaml:"create_timestamp"`
	UpdateTimestamp          time.Time `boil:"update_timestamp" json:"update_timestamp" toml:"update_timestamp" yaml:"update_timestamp"`
	DeliveryPhoneCountryCode string    `boil:"delivery_phone_country_code" json:"delivery_phone_country_code,omitempty" toml:"delivery_phone_country_code" yaml:"delivery_phone_country_code,omitempty"`
	AddressID                int       `boil:"address_id" json:"address_id,omitempty" toml:"address_id" yaml:"address_id,omitempty"`
	AreaID                   int       `boil:"area_id" json:"area_id,omitempty" toml:"area_id" yaml:"area_id,omitempty"`
	DeliveryUserIDCard       string    `boil:"delivery_user_id_card" json:"delivery_user_id_card,omitempty" toml:"delivery_user_id_card" yaml:"delivery_user_id_card,omitempty"`
	CancelURL                string    `boil:"cancel_url" json:"cancel_url,omitempty" toml:"cancel_url" yaml:"cancel_url,omitempty"`
	NotifyURL                string    `boil:"notify_url" json:"notify_url,omitempty" toml:"notify_url" yaml:"notify_url,omitempty"`
	DeliveryWaybillID        string    `boil:"delivery_waybill_id" json:"delivery_waybill_id,omitempty" toml:"delivery_waybill_id" yaml:"delivery_waybill_id,omitempty"`
	DeliveryExpressCompany   string    `boil:"delivery_express_company" json:"delivery_express_company,omitempty" toml:"delivery_express_company" yaml:"delivery_express_company,omitempty"`
	Comment                  string    `boil:"comment" json:"comment,omitempty" toml:"comment" yaml:"comment,omitempty"`
	Channel                  string    `boil:"channel" json:"channel,omitempty" toml:"channel" yaml:"channel,omitempty"`
	Source                   string    `boil:"source" json:"source,omitempty" toml:"source" yaml:"source,omitempty"`
	TotalProductAmount       int       `boil:"total_product_amount" json:"total_product_amount" toml:"total_product_amount" yaml:"total_product_amount"`
	SuccessRedirectURL       string    `boil:"success_redirect_url" json:"success_redirect_url,omitempty" toml:"success_redirect_url" yaml:"success_redirect_url,omitempty"`
	ExchangeRate             float64   `boil:"exchange_rate" json:"exchange_rate,omitempty" toml:"exchange_rate" yaml:"exchange_rate,omitempty"`
	DeliveryCarrierID        string    `boil:"delivery_carrier_id" json:"delivery_carrier_id,omitempty" toml:"delivery_carrier_id" yaml:"delivery_carrier_id,omitempty"`
	ProductTax               int       `boil:"product_tax" json:"product_tax,omitempty" toml:"product_tax" yaml:"product_tax,omitempty"`
	ShippingTax              int       `boil:"shipping_tax" json:"shipping_tax,omitempty" toml:"shipping_tax" yaml:"shipping_tax,omitempty"`
	DeliveryCustoms          string    `boil:"delivery_customs" json:"delivery_customs,omitempty" toml:"delivery_customs" yaml:"delivery_customs,omitempty"`
	DeclareUserName          string    `boil:"declare_user_name" json:"declare_user_name,omitempty" toml:"declare_user_name" yaml:"declare_user_name,omitempty"`
	Platform                 string    `boil:"platform" json:"platform,omitempty" toml:"platform" yaml:"platform,omitempty"`
	ParentOrderID            string    `boil:"parent_order_id" json:"parent_order_id,omitempty" toml:"parent_order_id" yaml:"parent_order_id,omitempty"`
	DiscountCode             string    `boil:"discount_code" json:"discount_code,omitempty" toml:"discount_code" yaml:"discount_code,omitempty"`
	CampaignID               int       `boil:"campaign_id" json:"campaign_id,omitempty" toml:"campaign_id" yaml:"campaign_id,omitempty"`
	CouponInstanceID         int64     `boil:"coupon_instance_id" json:"coupon_instance_id,omitempty" toml:"coupon_instance_id" yaml:"coupon_instance_id,omitempty"`
	CampaignTrackingData     string    `boil:"campaign_tracking_data" json:"campaign_tracking_data,omitempty" toml:"campaign_tracking_data" yaml:"campaign_tracking_data,omitempty"`
	DeliveryEmail            string    `boil:"delivery_email" json:"delivery_email,omitempty" toml:"delivery_email" yaml:"delivery_email,omitempty"`
	CustomsClearanceType     string    `boil:"customs_clearance_type" json:"customs_clearance_type,omitempty" toml:"customs_clearance_type" yaml:"customs_clearance_type,omitempty"`
}
