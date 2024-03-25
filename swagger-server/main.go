/*
 * Order Service
 *
 * Services that provides the capabilities for orders and Odoo
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//sw "github.com/GIT_USER_ID/GIT_REPO_ID/order"
	//
	sw "server/swagger-server/order"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(router.Run(":8080"))
}
