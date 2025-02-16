.PHONY: build
build: build-server build-storage

.PHONY: build-server
build-server:
	CGO_ENABLED=0 GOOS=linux go build -o bin/server cmd/server/main.go

.PHONY: build-storage
build-storage:
	CGO_ENABLED=0 GOOS=linux go build -o bin/storage cmd/storage-server/main.go

.PHONY: generate
generate: vendor-proto .generate

.PHONY: .generate
.generate:
		mkdir -p pkg/helper
		protoc -I vendor.protogen \
				--go_out=pkg/helper --go_opt=paths=import \
				--go-grpc_out=pkg/helper --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/helper \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/helper \
				api/helper/helper.proto
		mv pkg/helper/github.com/pashest/object-storage-service/pkg/helper/* pkg/helper/
		rm -rf pkg/helper/github.com
		mkdir -p pkg/storage
		protoc -I vendor.protogen \
				--go_out=pkg/storage --go_opt=paths=import \
				--go-grpc_out=pkg/storage --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/storage \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/storage \
				api/storage/storage.proto
		mv pkg/storage/github.com/pashest/object-storage-service/pkg/storage/* pkg/storage/
		rm -rf pkg/storage/github.com

.PHONY: vendor-proto
vendor-proto: .vendor-proto

.PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/helper
		mkdir -p vendor.protogen/api/storage
		cp api/helper/helper.proto vendor.protogen/api/helper
		cp api/storage/storage.proto vendor.protogen/api/storage
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/github.com/envoyproxy &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
		fi

.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
		go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go get -u google.golang.org/grpc
		go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go get -u github.com/envoyproxy/protoc-gen-validate
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
		go install github.com/envoyproxy/protoc-gen-validate


