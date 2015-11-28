PKGS := $(shell go list ./... | grep -v 'go.pedge.io/protoeasy/vendor')

export GO15VENDOREXPERIMENT=1

all: test install

deps:
	GO15VENDOREXPERIMENT=0 go get -d -v $(PKGS)

updatedeps:
	GO15VENDOREXPERIMENT=0 go get -d -v -u -f $(PKGS)

testdeps:
	GO15VENDOREXPERIMENT=0 go get -d -v -t $(PKGS)

updatetestdeps:
	GO15VENDOREXPERIMENT=0 go get -d -v -t -u -f $(PKGS)

vendor:
	go get -v github.com/tools/godep
	rm -rf Godeps
	rm -rf vendor
	GOOS=linux GOARCH=AMD64 godep save \
			 $(PKGS) \
			 github.com/gengo/grpc-gateway/protoc-gen-grpc-gateway/... \
			 github.com/gengo/grpc-gateway/third_party/googleapis/... \
			 github.com/golang/protobuf/proto/... \
			 github.com/golang/protobuf/protoc-gen-go/... \
			 go.pedge.io/google-protobuf/... \
			 go.pedge.io/googleapis/... \
			 google.golang.org/grpc
	rm -rf Godeps

build:
	go build $(PKGS)

install:
	go install \
		$(PKGS) \
		go.pedge.io/protoeasy/vendor/github.com/gengo/grpc-gateway/protoc-gen-grpc-gateway \
		go.pedge.io/protoeasy/vendor/github.com/golang/protobuf/protoc-gen-go

proto: install
	go get -v go.pedge.io/pkg/cmd/strip-package-comments
	protoeasy --go --grpc --go_import_path go.pedge.io/protoeasy --exclude vendor --exclude example .
	find . -name *\.pb\*\.go | grep -v vendor | xargs strip-package-comments

example: install
	rm -rf _example-out
	protoeasy \
		--out=_example-out \
		--cpp --cpp_rel_out=cpp \
		--csharp --csharp_rel_out=csharp \
		--objectivec --objectivec_rel_out=objectivec \
		--python --python_rel_out=python \
		--ruby --ruby_rel_out=ruby \
		--go --go_rel_out=go --go_import_path=go.pedge.io/protoeasy/example/example-out/go \
		--grpc \
		--grpc-gateway \
		example

lint:
	go get -v github.com/golang/lint/golint
	for file in $$(find . -name '*.go' | grep -v '\.pb\.go' | grep -v '\.pb\.gw\.go' | grep -v '\.pb\.log\.go'); do \
		golint $$file; \
		if [ -n "$$(golint $$file)" ]; then \
			exit 1; \
		fi; \
	done

vet:
	go vet $(PKGS)

errcheck:
	go get -v github.com/kisielk/errcheck
	errcheck $(PKGS)

#pretest: lint vet errcheck
pretest: vet errcheck

test: pretest
	go test $(PKGS)

clean:
	go clean $(PKGS)
	rm -rf _example-out

docker-build:
	docker build -t quay.io/pedge/protoeasy .

docker-push: docker-build
	docker push quay.io/pedge/protoeasy

docker-pull:
	docker pull quay.io/pedge/protoeasy

docker-launch:
	docker rm -f protoeasy || true
	docker run -d -p 6789:6789 --name=protoeasy quay.io/pedge/protoeasy

docker-test: docker-build
	docker run quay.io/pedge/protoeasy make test

.PHONY: \
	all \
	deps \
	updatedeps \
	testdeps \
	updatetestdeps \
	vendor \
	build \
	install \
	proto \
	example \
	lint \
	vet \
	errcheck \
	pretest \
	test \
	clean \
	docker-build \
	docker-push \
	docker-pull \
	docker-launch \
	docker-test
