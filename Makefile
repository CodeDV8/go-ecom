build.main:
	@echo "Building Main"
	@go build -o ../ecom/ecom main.go

build.api:
	@echo "Building API"
	@go build -buildmode=plugin -o ../ecom/system/api.so ../go-ecom-plugin-api/

build.cms:
	@echo "Building CMS"
	@go build -buildmode=plugin -o ../ecom/system/cms.so ../go-ecom-plugin-cms/

build: build.main build.api build.cms
