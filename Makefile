build_main:
	@echo "Building Main"
	@go build -o ../ecom/ecom main.go

build_api:
	@echo "Building API"
	@go build -buildmode=plugin -o ../ecom/system/api.so ../go-ecom-plugin-api/

build: build_main build_api
