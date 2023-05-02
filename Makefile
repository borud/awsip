all: lint vet test build

build: awsip

awsip:
	@cd cmd/$@ && go build -o ../../bin/$@

test:
	@go test ./...

vet:
	@go vet ./...

lint:
	@revive ./...

clean:
	@rm -rf bin

dep-install:
	@go install github.com/mgechev/revive@latest