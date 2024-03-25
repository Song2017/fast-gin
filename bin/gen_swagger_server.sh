#!/usr/bin/env sh

# configuration
input="order.yml"
output="swagger-server"
type="go-gin-server"
pkg_name="order"
pkg_version="0.1.0"

source _bin/generate_swagger_lib.sh
generate_swagger_lib $input $output $type $pkg_name $pkg_version

sed -e 's/sw ".\/order"/sw "server\/swagger-server\/order"/' $output/main.go >$output/main.go.tmp
mv -f $output/main.go.tmp $output/main.go

# swagger file
rm -f ./resources/openapi.yaml
cp -f $output/api/openapi.yaml ./resources
