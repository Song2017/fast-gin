/*
 * Order Service
 *
 * Services that provides the capabilities for orders and Odoo
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package order_api

import (
	"net/http"

	service "server/server/service/odoo"

	"github.com/gin-gonic/gin"
)

// CreateShipUsingPOST - createShip
func CreateShipUsingPOST(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// OdooUsingGET - odoo
func OdooUsingGET(c *gin.Context) {
	o := new(service.OdooGet)
	if !o.Invoke(c,
		o.Odoo.Input,
		o.Input,
		o.Process,
	) {
		return
	}

	c.JSON(http.StatusOK, o.Response)
}