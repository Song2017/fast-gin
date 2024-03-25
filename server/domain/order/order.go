package order_domain

import (
	"context"
	"fmt"
	initial "server/_init"
	te "server/_pkg"
	swag "server/swagger-server/order"
	"strings"
)

var (
	SEQUENCE_PATTERN  = "Order:Id:%s:%s"
	ORDER_ID_TEMPLATE = "DJ%s#%s"

	COMBINE_ORDER_ID_TEMPLATE = "DJCOMBINE%s%s%s"
)

type Order struct {
	swag.Order
}

func (o Order) generateId(storeId string, orderID string, isCombine bool) string {
	storeId = strings.ToUpper(strings.ReplaceAll(storeId, ".", ""))
	if !isCombine {
		return fmt.Sprintf(ORDER_ID_TEMPLATE, storeId, orderID)
	}
	redisKey := fmt.Sprintf(SEQUENCE_PATTERN, storeId, te.GetDay(nil))
	c := context.Background()
	initial.O_REDIS.SetNX(c, redisKey, 0, te.OneHour)
	sequence := initial.O_REDIS.IncrBy(c, redisKey, 1)
	return fmt.Sprintf(COMBINE_ORDER_ID_TEMPLATE,
		storeId, te.GetMin(nil), fmt.Sprintf("%05s", sequence.String()))
}
