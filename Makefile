lint:
	golangci-lint run ./...

format:
	gofmt -w -s .
	gci -local github.com/kyoukaya/subsonar -w .
