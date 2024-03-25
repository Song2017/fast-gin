package order_odoo

import (
	core "server/_sys/core"
	domain_store "server/server/domain/store"
	swag "server/swagger-server/order"

	"github.com/gin-gonic/gin"
)

type Odoo struct {
	core.ServiceGin
	Store *domain_store.StoreMeta
}

func (o *Odoo) Input(c *gin.Context) {
	o.ServiceGin.Input()

	storeId := c.Param("storeId")
	o.Store = domain_store.GetStore(storeId)
	if o.Store == nil || o.Store.Status == 404 {
		o.Response["error"] = &swag.BaseResponseOfOrder{
			Code: 400,
			Msg:  "Please check store ID",
		}
	}

}
