openapi-generator-cli generate \
-i ./api/openapi.yaml \
-g go \
-c ./openapi/config.yml \
-t ./openapi/templates \
--skip-validate-spec