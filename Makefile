PACKAGE_NAME=com.tylerhaigh/go-simple-hl7/pkg/hl7

test:
	go test -timeout 30s -v $(PACKAGE_NAME)

cover:
	go test -timeout 30s -v -coverprofile=coverage.out $(PACKAGE_NAME)

coverage_report: cover
	go tool cover -html=coverage.out
