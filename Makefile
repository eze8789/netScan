help:  ## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build: ## Build go binary
	CGO_ENABLED=0 GOOS=linux go build -a -o netscan -ldflags="-s -w" .

run:   ## Run go program, you can pass arguments using make run ARGS="--help"
	go run ./... $(ARGS)