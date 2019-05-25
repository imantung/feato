-include .env

mock:
	@echo "  >  Generate mock class..."
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@$(GOPATH)/bin/mockgen -source=toggle_router.go -destination=mock_toggle_router.go -package=feato;

example-simple:
	@go run ./example/simple/main.go
	
.PHONY: mock