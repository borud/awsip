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

release: all
	@cd cmd/awsip && GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o ../../bin/awsip
	@cd bin && zip awsip.amd64-linux.zip awsip && rm awsip

	@cd cmd/awsip && GOOS=darwin GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o ../../bin/awsip
	@cd bin && zip awsip.amd64-macos.zip awsip && rm awsip
	
	@cd cmd/awsip && GOOS=darwin GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o ../../bin/awsip
	@cd bin && zip awsip.arm64-macos.zip awsip && rm awsip

	@cd cmd/awsip && GOOS=windows GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o ../../bin/awsip.exe
	@cd bin && zip awsip.amd64-win.zip awsip.exe && rm awsip.exe

	@cd cmd/awsip && GOOS=linux GOARCH=arm GOARM=5 go build -ldflags=$(LDFLAGS) -o ../../bin/awsip
	@cd bin && zip awsip.arm5-rpi-linux.zip awsip && rm awsip