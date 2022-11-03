.PHONY: install
install:
	go install github.com/matryer/moq@latest
	go install github.com/go-delve/delve/cmd/dlv@latest
	go install github.com/cosmtrek/air@latest
	go install github.com/kisielk/errcheck@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	brew install golangci-lint
	brew upgrade golangci-lint
	go install golang.org/x/tools/cmd/goimports@latest
.PHONY: moqgen
moqgen: 
	go generate ./repository/...
	go generate ./usecase/...
.PHONY: ci_setup_server
ci_setup_server:
	FIRESTORE_EMULATOR_HOST=localhost:3600 APP_ENV=local go run main.go &
.PHONY: precommit
precommit: goimports fmt vet errcheck staticcheck
.PHONY: goimports
goimports:
	find . -print | grep --regex '.*\.go' | grep -v "./.go-version"  | grep -v "./.golangci.yml" | grep -v "moq" | xargs goimports -w -local "github.com/wheatandcat/memoir-backend" 
fmt:
	go fmt ./...
vet:
	go vet ./...
errcheck:
	errcheck
staticcheck:
	staticcheck ./...
golangci:
	golangci-lint run ./...