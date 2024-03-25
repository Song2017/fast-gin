package order_dao

import (
	db "server/_db"
	"time"
)

var allACColumn = ` o.id, o.order_id, o.outer_order_id, o.parent_order_id, o.store_id,
o.delivery_address, o.delivery_express_company, o.delivery_waybill_id,
o.status, o.create_timestamp, o.update_timestamp `

func (f *OrderAddressCombination) TableName() string {
	return "order_address_combination"
}

// OrderAddressCombination is an object representing the database table.
type OrderAddressCombination struct {
	ID                     int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	ParentOrderID          string    `boil:"parent_order_id" json:"parent_order_id,omitempty" toml:"parent_order_id" yaml:"parent_order_id,omitempty"`
	OrderID                string    `boil:"order_id" json:"order_id,omitempty" toml:"order_id" yaml:"order_id,omitempty"`
	OuterOrderID           string    `boil:"outer_order_id" json:"outer_order_id,omitempty" toml:"outer_order_id" yaml:"outer_order_id,omitempty"`
	DeliveryAddress        string    `boil:"delivery_address" json:"delivery_address,omitempty" toml:"delivery_address" yaml:"delivery_address,omitempty"`
	DeliveryWaybillID      string    `boil:"delivery_waybill_id" json:"delivery_waybill_id,omitempty" toml:"delivery_waybill_id" yaml:"delivery_waybill_id,omitempty"`
	DeliveryExpressCompany string    `boil:"delivery_express_company" json:"delivery_express_company,omitempty" toml:"delivery_express_company" yaml:"delivery_express_company,omitempty"`
	DeliveryURL            string    `boil:"delivery_url" json:"delivery_url,omitempty" toml:"delivery_url" yaml:"delivery_url,omitempty"`
	StoreID                string    `boil:"store_id" json:"store_id,omitempty" toml:"store_id" yaml:"store_id,omitempty"`
	Status                 int       `boil:"status" json:"status" toml:"status" yaml:"status"`
	CreateUser             string    `boil:"create_user" json:"create_user,omitempty" toml:"create_user" yaml:"create_user,omitempty"`
	UpdateUser             string    `boil:"update_user" json:"update_user,omitempty" toml:"update_user" yaml:"update_user,omitempty"`
	CreateTimestamp        time.Time `boil:"create_timestamp" json:"create_timestamp" toml:"create_timestamp" yaml:"create_timestamp"`
	UpdateTimestamp        time.Time `boil:"update_timestamp" json:"update_timestamp" toml:"update_timestamp" yaml:"update_timestamp"`
}

func QueryAllByOrderId(orderId string) *[]OrderAddressCombination {
	var orders []OrderAddressCombination
	sqlText := "select" + allACColumn
	sqlText += `from "order_address_combination" o
	where o.parent_order_id in (
	select distinct parent_order_id from "order_address_combination"
	where parent_order_id = ? or order_id = ? )
	order by create_timestamp desc`
	db.Raw(&orders, sqlText, orderId, orderId)
	return &orders
}
