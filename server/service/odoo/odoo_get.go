package order_odoo

import (
	"log"
	db "server/_db"
	pkg "server/_pkg"
	po "server/server/po"
	swag "server/swagger-server/order"
	"strings"

	"github.com/gin-gonic/gin"
)

var strOdooParas = []string{"created_range", "updated_range", "order_id", "client_order_ref"}

type OdooGet struct {
	Odoo
}

func (o *OdooGet) Input(c *gin.Context) {
	notEmpty := ""
	for _, str := range strOdooParas {
		o.Request[str] = c.Query(str)
		notEmpty += c.Query(str)
	}
	resp := o.Response
	odooResp := &swag.OdooResponse{
		Code:    400,
		Message: "Failed",
	}
	if notEmpty != "" {
		odooResp.Message = "Please check time range"
		resp["error"] = odooResp
		return
	}
	o.Request["page_no"] = pkg.ToInt(c.Query("page_no"), 1)
	o.Request["page_size"] = pkg.ToInt(c.Query("page_size"), 50)
	o.Request["combine_orders"] = pkg.ToBool(
		c.Query("combine_orders"), false)

	if !o.Request["combine_orders"].(bool) {
		o.Request["combine_orders"] = o.Store.CombineOrders
	}

	// 2024-03-22T07:40:52.997973Z/2024-03-22T09:41:18.668858Z
	if c.Query("created_range") != "" {
		date_range := strings.Split(c.Query("created_range"), "/")
		for i, v := range date_range {
			date_range[i] = pkg.TruncateStr(v, ".")
		}
		o.Request["date_range"] = []string{}
	}
	log.Println("OdooUsingGET", o.Request)
}

func (o *OdooGet) Process(c *gin.Context) {
	// orders := po.QueryAllByOrderId(o.Request["order_id"].(string))
	if o.Request["client_order_ref"] != nil {
		var orders []po.OrderPO
		db.SelectAllByField(&orders, o.Request["client_order_ref"].(string), "outer_order_id")
		o.Response["result"] = orders
		return
	}
	if o.Request["order_id"] != nil {
		var order po.OrderPO
		if !o.Request["combine_orders"].(bool) {
			db.SelectById(o.Request["order_id"].(string), &order)
			o.Response["result"] = order
		} else {

		}
		return

	}

}
