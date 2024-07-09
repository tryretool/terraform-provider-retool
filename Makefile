#-----------------------------------------------------------------------------------------------------------------------
# Testing
#-----------------------------------------------------------------------------------------------------------------------
.PHONY: test-unit test-acc test-acc-record

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
