ifndef VENDOR_IGNORE_DIRS
	VENDOR_IGNORE_DIRS = go.pedge.io
endif

TMP_VENDORIZE_CLIENT_DIR = $(GOPATH)/src/github.com/pachyderm/tmp_vendorize_client

vendor-client:
	rm -rf $(TMP_VENDORIZE_CLIENT_DIR)
	mkdir -p $(TMP_VENDORIZE_CLIENT_DIR)
	cp tmp_vendorize_client_go $(TMP_VENDORIZE_CLIENT_DIR)/tmp_vendorize_client.go
	cd $(TMP_VENDORIZE_CLIENT_DIR) && govendor init && govendor add +external
	mkdir -p vendor/github.com/pachyderm/pachyderm/src/
	cp -R $(TMP_VENDORIZE_CLIENT_DIR)/vendor/github.com/pachyderm/pachyderm/src/client vendor/github.com/pachyderm/pachyderm/src

vendor-update:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GO15VENDOREXPERIMENT=0 go get -d -v -t -u -f ./src/... ./.

vendor-without-update:
	go get -v github.com/kardianos/govendor
	rm -rf vendor || echo "no existing vendor directory"
	govendor init
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 govendor add +external
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 govendor update +vendor
	$(foreach vendor_dir, $(VENDOR_IGNORE_DIRS), rm -rf vendor/$(vendor_dir) || exit; git checkout vendor/$(vendor_dir) || exit;)

vendor: vendor-update vendor-without-update
