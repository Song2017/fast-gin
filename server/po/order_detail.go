package order_dao

import (
	"time"

	"github.com/volatiletech/null/v8"
)

func (f *OrderDetail) TableName() string {
	return "order_detail"
}

// OrderDetail is an object representing the database table.
type OrderDetail struct {
	ID              int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	OrderID         string      `boil:"order_id" json:"order_id" toml:"order_id" yaml:"order_id"`
	SkuID           string      `boil:"sku_id" json:"sku_id" toml:"sku_id" yaml:"sku_id"`
	ProductName     string      `boil:"product_name" json:"product_name" toml:"product_name" yaml:"product_name"`
	ImageURL        null.String `boil:"image_url" json:"image_url,omitempty" toml:"image_url" yaml:"image_url,omitempty"`
	Spec            null.String `boil:"spec" json:"spec,omitempty" toml:"spec" yaml:"spec,omitempty"`
	RRP             null.Int    `boil:"rrp" json:"rrp,omitempty" toml:"rrp" yaml:"rrp,omitempty"`
	Price           int         `boil:"price" json:"price" toml:"price" yaml:"price"`
	Currency        string      `boil:"currency" json:"currency" toml:"currency" yaml:"currency"`
	Quantity        int         `boil:"quantity" json:"quantity" toml:"quantity" yaml:"quantity"`
	ProductID       null.String `boil:"product_id" json:"product_id,omitempty" toml:"product_id" yaml:"product_id,omitempty"`
	Tax             null.Int    `boil:"tax" json:"tax,omitempty" toml:"tax" yaml:"tax,omitempty"`
	Discount        null.Int    `boil:"discount" json:"discount,omitempty" toml:"discount" yaml:"discount,omitempty"`
	TaxRate         float64     `boil:"tax_rate" json:"tax_rate,omitempty" toml:"tax_rate" yaml:"tax_rate,omitempty"`
	UpdateTimestamp time.Time   `boil:"update_timestamp" json:"update_timestamp" toml:"update_timestamp" yaml:"update_timestamp"`
	CreateTimestamp time.Time   `boil:"create_timestamp" json:"create_timestamp" toml:"create_timestamp" yaml:"create_timestamp"`
	BlockedItem     int         `boil:"blocked_item" json:"blocked_item" toml:"blocked_item" yaml:"blocked_item"`
	OuterProductID  null.String `boil:"outer_product_id" json:"outer_product_id,omitempty" toml:"outer_product_id" yaml:"outer_product_id,omitempty"`
	OuterSkuID      null.String `boil:"outer_sku_id" json:"outer_sku_id,omitempty" toml:"outer_sku_id" yaml:"outer_sku_id,omitempty"`
	Barcode         null.String `boil:"barcode" json:"barcode,omitempty" toml:"barcode" yaml:"barcode,omitempty"`
}
