# -*- mode: makefile -*-
GOPROXY=""

go-deps:i
	@echo "Verifying and cleaning up dependencies for ${GOSERVICE}..."
	GOPROXY=$(GOPROXY) go mod tidy && go mod verify

go-build:
	@echo "Building '${GOSERVICE}'..."
	GOPROXY=$(GOPROXY) CGO_ENABLED=0 go build -o ${GOSERVICE} -v .

go-test:
	@echo "Running all ${GOSERVICE} tests.."
	cd ./sorting-service
	SERVICE_ENV=test GOPROXY=${GOPROXY} go test -cover -v -race ./...

go-run:
	SERVICE_ENV=development GOPROXY=${GOPROXY} SERVICE_LOG=debug go run ./sorting-service/main.go ./sorting-service/service.go
