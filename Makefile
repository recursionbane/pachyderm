#### VARIABLES
# RUNARGS: arguments for run
# DOCKER_OPTS: docker-compose options for run, test, launch-*
# TESTPKGS: packages for test, default ./src/...
# TESTFLAGS: flags for test
# VENDOR_ALL: do not ignore some vendors when updating vendor directory
# VENDOR_IGNORE_DIRS: ignore vendor dirs
# KUBECTLFLAGS: flags for kubectl
####

ifndef TESTPKGS
	TESTPKGS = ./src/...
endif
ifndef VENDOR_IGNORE_DIRS
	VENDOR_IGNORE_DIRS = go.pedge.io
endif
ifdef VENDOR_ALL
	VENDOR_IGNORE_DIRS =
endif

COMPILE_RUN_ARGS = -v /var/run/docker.sock:/var/run/docker.sock --privileged=true
CLUSTER_NAME = pachyderm

all: build

version-prefix: 
	@echo 'package main; import "fmt"; import "github.com/pachyderm/pachyderm/src/client"; func main() { fmt.Printf("%v.%v", client.Version.Major, client.Version.Minor) }' > /tmp/pachyderm_version.go
	@go run /tmp/pachyderm_version.go

version:
	@echo 'package main; import "fmt"; import "github.com/pachyderm/pachyderm/src/client"; func main() { fmt.Println(client.Version.VersionString()) }' > /tmp/pachyderm_version.go
	@go run /tmp/pachyderm_version.go

deps:
	GO15VENDOREXPERIMENT=0 go get -d -v ./src/... ./.

deps-client: 
	GO15VENDOREXPERIMENT=0 go get -d -v ./src/client/...

update-deps:
	GO15VENDOREXPERIMENT=0 go get -d -v -u -f ./src/... ./.

test-deps:
	GO15VENDOREXPERIMENT=0 go get -d -v -t ./src/... ./.

update-test-deps:
	GO15VENDOREXPERIMENT=0 go get -d -v -t -u -f ./src/... ./.

build-clean-vendored-client:
	rm -rf src/server/vendor/github.com/pachyderm/pachyderm/src/client

build: 
	GO15VENDOREXPERIMENT=1 go build $$(go list ./src/client/... | grep -v '/src/client$$')
	GO15VENDOREXPERIMENT=1 go build $$(go list ./src/server/... | grep -v '/src/server/vendor/' | grep -v '/src/server$$')

install:
	# GOPATH/bin must be on your PATH to access these binaries:
	GO15VENDOREXPERIMENT=1 go install ./src/server/cmd/pachctl ./src/server/cmd/pach-deploy ./src/server/cmd/pachctl-doc

release: deps-client
	./etc/build/tag_release
docker-build-compile:
	docker build -t pachyderm_compile .

docker-build-job-shim: docker-build-compile
	docker run $(COMPILE_RUN_ARGS) pachyderm_compile sh etc/compile/compile.sh job-shim

docker-build-pachd: docker-build-compile
	docker run $(COMPILE_RUN_ARGS) pachyderm_compile sh etc/compile/compile.sh pachd

docker-build: docker-build-job-shim docker-build-pachd

docker-push-job-shim: docker-build-job-shim
	docker push pachyderm/job-shim

docker-push-pachd: docker-build-pachd
	docker push pachyderm/pachd

docker-push: docker-push-job-shim docker-push-pachd

launch-kube:
	etc/kube/start-kube-docker.sh

clean-launch-kube:
	docker kill $$(docker ps -q)

kube-cluster-assets: install
	pach-deploy -s 32 >etc/kube/pachyderm.json

launch: install
	kubectl $(KUBECTLFLAGS) create -f etc/kube/pachyderm.json
	# wait for the pachyderm to come up
	# if we can create a repo, that means that the cluster is ready to serve
	until timeout 5s $(GOPATH)/bin/pachctl list-repo 2>/dev/null >/dev/null; do sleep 5; done

launch-dev: launch-kube launch

clean-launch:
	kubectl $(KUBECTLFLAGS) delete --ignore-not-found job -l suite=pachyderm
	kubectl $(KUBECTLFLAGS) delete --ignore-not-found all -l suite=pachyderm
	kubectl $(KUBECTLFLAGS) delete --ignore-not-found serviceaccount -l suite=pachyderm
	kubectl $(KUBECTLFLAGS) delete --ignore-not-found secret -l suite=pachyderm

integration-tests:
	CGOENABLED=0 go test ./src/server -timeout 120s

docker-proto-run:
	cd /go/src/github.com/pachyderm/pachyderm && \
	go install ./src/server/cmd/protofix && \
	rm -rf src/server/vendor && \
	protoeasy --grpc --go --go-import-path github.com/pachyderm/pachyderm/src src && \
	protofix fix src && \
	git checkout src/server/vendor

docker-build-proto:
	docker build -t pachyderm/protofix -f Dockerfile.proto .

docker-proto:
	docker run -v $(PWD):/go/src/github.com/pachyderm/pachyderm pachyderm/protofix make -f /go/src/github.com/pachyderm/pachyderm/Makefile docker-proto-run
	sudo chown -R `whoami` src/

proto:
	go get -v go.pedge.io/protoeasy/cmd/protoeasy
	rm -rf src/server/vendor
	sudo env PATH=$(PATH) GOPATH=$(GOPATH) protoeasy --grpc --go --go-import-path github.com/pachyderm/pachyderm/src src

protofix:
	go install github.com/pachyderm/pachyderm/src/server/cmd/protofix
	protofix fix src
	git checkout src/server/vendor
	sudo chown -R `whoami` src/

pretest:
	go get -v github.com/kisielk/errcheck
	go get -v github.com/golang/lint/golint
	rm -rf src/server/vendor
	for file in $$(find "./src" -name '*.go' | grep -v '\.pb\.go' | grep -v '\.pb\.gw\.go'); do \
		golint $$file | grep -v unexported; \
		if [ -n "$$(golint $$file | grep -v unexported)" ]; then \
		exit 1; \
		fi; \
		done;
	go vet -n ./src/... | while read line; do \
		modified=$$(echo $$line | sed "s/ [a-z0-9_/]*\.pb\.gw\.go//g"); \
		$$modified; \
		if [ -n "$$($$modified)" ]; then \
		exit 1; \
		fi; \
		done
	git checkout src/server/vendor
	#errcheck $$(go list ./src/... | grep -v src/cmd/ppsd | grep -v src/pfs$$ | grep -v src/pps$$)

test: pretest localtest docker-build clean-launch launch integration-tests

localtest: deps-client
	CGOENABLED=0 GO15VENDOREXPERIMENT=1 go test -cover -v -short $$(go list ./src/client/...)
	CGOENABLED=0 GO15VENDOREXPERIMENT=1 go test -cover -v -short $$(go list ./src/server/... | grep -v '/src/server/vendor/')

clean: clean-launch clean-launch-kube

doc: install
	# we rename to pachctl because the program name is used in generating docs
	cp $(GOPATH)/bin/pachctl-doc ./pachctl
	rm -rf doc/pachctl && mkdir doc/pachctl
	./pachctl
	rm ./pachctl

grep-data:
	go run examples/grep/generate.go >examples/grep/set1.txt
	go run examples/grep/generate.go >examples/grep/set2.txt

grep-example:
	sh examples/grep/run.sh

logs:
	kubectl get pod -l app=pachd | sed '1d' | cut -f1 -d ' ' | xargs -n 1 -I pod sh -c 'echo pod && kubectl logs pod'

kubectl:
	gcloud config set container/cluster $(CLUSTER_NAME)
	gcloud container clusters get-credentials $(CLUSTER_NAME)

cluster:
	gcloud container clusters create $(CLUSTER_NAME) --scopes storage-rw
	gcloud config set container/cluster $(CLUSTER_NAME)
	gcloud container clusters get-credentials $(CLUSTER_NAME)
	gcloud components update kubectl
	gcloud compute firewall-rules create pachd --allow=tcp:30650

clean-cluster:
	gcloud container clusters delete $(CLUSTER_NAME)


.PHONY: \
	doc \
	all \
	version \
	deps \
	update-deps \
	test-deps \
	update-test-deps \
	vendor-update \
	vendor-without-update \
	vendor \
	build \
	install \
	docker-build-test \
	docker-build-compile \
	docker-build \
	docker-build-pachd \
	docker-push \
	docker-push-pachd \
	run \
	launch \
	proto \
	pretest \
	docker-clean-test \
	go-test \
	go-test-long \
	test \
	test-long \
	clean
