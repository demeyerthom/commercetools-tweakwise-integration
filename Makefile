generate: generate/server

openapi-generator:
	npm install -g @openapitools/openapi-generator-cli

generate/server:
	# Generate code
	 openapi-generator-cli generate \
 		--generator-name go-server \
 		--input-spec openapi.json \
 		--output internal \
 		--ignore-file-override .openapi-generator-ignore \
 		--additional-properties packageName=server \
 		--additional-properties outputAsLibrary=true \
 		--additional-properties sourceFolder=server
	# Cleanup
	rm -rf -d internal/.openapi-generator internal/api
