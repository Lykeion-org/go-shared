PROTO_DIR=../proto
OUT_DIR=pkg/grpc/generated

PROTOC_GEN_GO := $(shell which protoc-gen-go)
PROTOC_GEN_GO_GRPC := $(shell which protoc-gen-go-grpc)

gen-proto: 
	@if [ -z "$(PROTOC_GEN_GO)" ] || [ -z "$(PROTOC_GEN_GO_GRPC)" ]; then \
		echo "protoc-gen-go or protoc-gen-go-grpc not found in PATH"; \
		echo "Install with: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"; \
		echo "and: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"; \
		exit 1; \
	fi
	protoc -I=$(PROTO_DIR) \
		--go_out=paths=source_relative:$(OUT_DIR) \
		--go-grpc_out=paths=source_relative:$(OUT_DIR) \
		$(PROTO_DIR)/aegis/*.proto \
		$(PROTO_DIR)/demes/*.proto \
		$(PROTO_DIR)/kosmos/*.proto \
		$(PROTO_DIR)/krites/*.proto \
		$(PROTO_DIR)/lexora/*.proto  \
		$(PROTO_DIR)/shared/*.proto  \
		$(PROTO_DIR)/paideia/*.proto
 