#-----------------------------------------------------------------------------------------------------------------------
# Dependencies
#-----------------------------------------------------------------------------------------------------------------------
$(GO_BIN)/golangci-lint:
	${call print, "Installing golangci-lint"}
	@go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@latest

#-----------------------------------------------------------------------------------------------------------------------
# Docs
#-----------------------------------------------------------------------------------------------------------------------
.PHONY: docs check-docs

docs: ## Generate docs
	${call print, "Generating docs"}
	@go generate

check-docs: ## Check if docs are up-to-date
	${call print, "Checking that documentation was generated correctly"}
	@go generate
	@if [ -n "$$(git status --porcelain)" ]; \
	then \
		echo "Go generate resulted in changed files:"; \
		echo "$$(git diff)"; \
		echo "Please run \`make docs\` to regenerate docs."; \
		exit 1; \
	fi
	@echo "Documentation is generated correctly."

#-----------------------------------------------------------------------------------------------------------------------
# Static analysis and linting
#-----------------------------------------------------------------------------------------------------------------------
.PHONY: lint

lint: $(GO_BIN)/golangci-lint ## Run go linter checks
	${call print, "Running golangci-lint over project"}
	@golangci-lint run -v --fix -c .golangci.yml ./...


#-----------------------------------------------------------------------------------------------------------------------
# Testing
#-----------------------------------------------------------------------------------------------------------------------
.PHONY: test-unit test-acc test-acc-record test-acc-replay test-acc-sweep

test-unit: ## Run unit tests. To run a specific test, pass the FILTER var. Usage `make test-unit FILTER="TestAccResourceServer"`
	${call print, "Running unit tests"}
	@TF_ACC= \
		go test \
		-v \
		-run "$(FILTER)" \
		-timeout 30s \
		./...

test-acc: ## Run acceptance tests. Don't forget to set RETOOL_HOST, RETOOL_SCHEME and RETOOL_ACCESS_TOKEN env vars. To run a specific test, pass the FILTER var. Usage `make test-acc FILTER="TestAccResourceServer"`
	${call print, "Running acceptance tests"}
	@TF_ACC=1 \
		go test \
		-v \
		-run "$(FILTER)" \
		-timeout 30m \
		./internal/provider/...

## Run acceptance tests and record previously unseen HTTP requests. In most cases, you need to run a specific test - pass the FILTER var. Usage `make test-acc-record FILTER="TestAccResourceServer"`
## Don't forget to set RETOOL_HOST, RETOOL_SCHEME and RETOOL_ACCESS_TOKEN env vars. 
## Note: You'll need to delete the recordings manually if you want to re-record them.
test-acc-record: 
	${call print, "Running acceptance tests"}
	@RETOOL_HTTP_RECORDINGS=on \
	    TF_ACC=1 \
		go test \
		-v \
		-run "$(FILTER)" \
		-timeout 30m \
		./internal/provider/...

## Run acceptance tests using replayed HTTP requests. Usage `make test-acc-replay FILTER="TestAccResourceServer"`
## You don't need to set RETOOL_HOST, RETOOL_SCHEME and RETOOL_ACCESS_TOKEN env vars for this command. 
test-acc-replay: 
	${call print, "Running acceptance tests"}
	@RETOOL_HTTP_RECORDINGS=on \
		RETOOL_HOST=recorded.retool.dev \
		RETOOL_SCHEME=https \
		RETOOL_ACCESS_TOKEN=some-token \
	    TF_ACC=1 \
		go test \
		-v \
		-run "$(FILTER)" \
		-timeout 30m \
		./internal/provider/...

test-acc-sweep: ## Remove all resources created by acceptance tests. Don't forget to set RETOOL_HOST, RETOOL_SCHEME and RETOOL_ACCESS_TOKEN env vars.
	${call print, "Removing all resources created by acceptance tests"}
	@TF_ACC=1  \
		go test \
		./internal/provider/... \
		-sweep=default \
		-v
