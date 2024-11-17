OS=linux
TESTS=go test $$(go list ./... | grep -v /vendor/) -cover

build:
	go build -o ${BINARY}

unittest:
	go test -short $$(go list ./... | grep -v /vendor/)

test:
	go test ./... -coverprofile=coverage.out -json > test.out
	gocover-cobertura < coverage.out > coverage.xml

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

run:
	go run main.go

swag:
	swag init --outputTypes yaml --output ./

lint:
	golangci-lint run

static:
	CGO_ENABLED=0 GOOS=${OS} go build -o ${BINARY} -ldflags="-w -s" -a -installsuffix cgo

docker:
	docker build --no-cache -t ${GCP_CONTAINER_REGISTRY}/${BINARY}:${VERSION} -t ${GCP_CONTAINER_REGISTRY}/${BINARY}:latest .
	docker push ${GCP_CONTAINER_REGISTRY}/${BINARY}:${VERSION}
	docker rmi ${GCP_CONTAINER_REGISTRY}/${BINARY}:${VERSION}

migrate:
	goose -dir migration mysql "${DB_CREDENTIALS_USR}:${DB_CREDENTIALS_PSW}@tcp(${DB_HOST}:3306)/${BINARY}?parseTime=true" up

rollback:
	goose -dir migration mysql "${DB_CREDENTIALS_USR}:${DB_CREDENTIALS_PSW}@tcp(${DB_HOST}:3306)/${BINARY}?parseTime=true" down-to ${VERSION}

.PHONY: clean unittest test

swagger-gen:
	swag init --outputTypes yaml --output ./

before:
	@go mod tidy
	@go test ./...
	@golangci-lint run