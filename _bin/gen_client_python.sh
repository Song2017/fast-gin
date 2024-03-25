#!/usr/bin/env sh

# configuration
input="z.yml"
pkg_path="com.gitlab"
output="tests/python_client"
type="python"
pkg_name="z_cli"
pkg_version="0.1.0"

source bin/generate_swagger_lib.sh
generate_swagger_lib $input $output $type $pkg_name $pkg_version