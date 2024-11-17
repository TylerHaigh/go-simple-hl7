PACKAGE_ROOT=github.com/TylerHaigh/go-simple-hl7
PACKAGE_NAME=github.com/TylerHaigh/go-simple-hl7/pkg/hl7

test:
	go test -timeout 30s -v $(PACKAGE_NAME)

cover:
	go test -timeout 30s -v -coverprofile=coverage.out $(PACKAGE_NAME)

coverage_report: cover
	go tool cover -html=coverage.out

protoc:
	protoc -I proto \
	--go_out . \
	--go_opt module=$(PACKAGE_ROOT) \
	--go-grpc_out . \
	--go-grpc_opt module=$(PACKAGE_ROOT) \
	proto/hl7-server.proto
