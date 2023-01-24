# Makefile to build mendix-privatecloud-go-sdk library
.PHONY: test
default:

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w .

goimports: gci

gci:
	@echo "==> Fixing imports code with gci..."
	gci write -s standard -s default -s "prefix(github.com\yogendra0sharma\mendix-privatecloud-go-sdk)" -s blank -s dot .

###############################################################################
###                                Linting                                  ###
###############################################################################
lint:
	@echo "==> Checking source code against linters..."
	golangci-lint run --timeout 2m0s ./...

.PHONY: lint lint-fix

format:
	golangci-lint run --fix

.PHONY: format

download:
	@echo Download go.mod dependencies
	@go mod download

install-tools: download
	@echo Installing tools from tools.go
	@go install $$(go list -f '{{range .Imports}}{{.}} {{end}}' tools.go)

###############################################################################
###                              Documentation                              ###
###############################################################################
docs:
	godoc -http=:6060

###############################################################################
###                           Tests & Simulation                            ###
###############################################################################
test:
	go test -v -cpu 4 -covermode=count -coverpkg github.com/yogendra0sharma/mendix-privatecloud-go-sdk/... -coverprofile=coverage.out ./...

cover: test
	@echo "==> generate test coverage..."
	go tool cover -html=coverage.out

upgrade:
	@echo "==> Upgrading Go"
	@GO111MODULE=on go get -u all && go mod tidy
	@echo "==> Upgrading pre-commit"
	@pre-commit autoupdate
	@echo "Please, upgrade workflows manually"

bumpversion-minor:
	@echo "==> Updating sdk version minor"
	bumpversion minor

bumpversion-major:
	@echo "==> Updating sdk version major"
	bumpversion major

bumpversion-patch:
	@echo "==> Updating sdk version patch"
	bumpversion patch