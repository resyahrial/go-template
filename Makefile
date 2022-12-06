test-cover: 
	go test -failfast -coverprofile cover.out ./...

display-test-cover:
	go tool cover -html cover.out

# Make task for CI golangci-lint report fo CI sonarscanner
lint: 
	go mod verify\
	&& go mod download\
	&& wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.47.1\
	&& ./bin/golangci-lint run --out-format checkstyle --issues-exit-code 0 > golangci-lint.out
	
wire-inject:
	cd internal/factory; \
	wire

generate-rest:
	cd internal/api/rest/v1; \
	oapi-codegen -config types.cfg.yml docs.yml

	cd internal/api/rest/v1; \
	oapi-codegen -config server.cfg.yml docs.yml