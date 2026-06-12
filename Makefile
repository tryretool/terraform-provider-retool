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
.PHONY: test-unit test-acc test-acc-record test-acc-record-each test-acc-replay test-acc-sweep

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
		-p 1 \
		-parallel 1 \
		-run "$(FILTER)" \
		-timeout 30m \
		./internal/provider/...

## Run acceptance tests and record previously unseen HTTP requests. In most cases, you need to run a specific test - pass the FILTER var. Usage `make test-acc-record FILTER="TestAccResourceServer"`
## Don't forget to set RETOOL_HOST, RETOOL_SCHEME and RETOOL_ACCESS_TOKEN env vars. 
## Note: You'll need to delete the recordings manually if you want to re-record them.
## Tests run fully serially (-p 1 across packages, -parallel 1 within a package) to stay
## under the instance's API rate limit when hitting the live API.
test-acc-record: 
	${call print, "Running acceptance tests"}
	@RETOOL_HTTP_RECORDINGS=on \
	    TF_ACC=1 \
		go test \
		-v \
		-p 1 \
		-parallel 1 \
		-run "$(FILTER)" \
		-timeout 30m \
		./internal/provider/...

## Record acceptance tests one at a time, pausing between them to respect the instance API rate limit.
## Resumable: already-recorded cassettes are skipped, so re-run after a 429 to finish the rest.
## Don't forget to set RETOOL_HOST, RETOOL_SCHEME and RETOOL_ACCESS_TOKEN env vars.
## Optionally pass a name regex (FILTER) and/or SLEEP_SECONDS. Usage `make test-acc-record-each FILTER=TestAccSSO SLEEP_SECONDS=60`
test-acc-record-each:
	${call print, "Recording acceptance tests one at a time"}
	@SLEEP_SECONDS=$(SLEEP_SECONDS) ./scripts/record_acc_tests.sh "$(FILTER)"

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
