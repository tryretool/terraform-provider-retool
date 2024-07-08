#-----------------------------------------------------------------------------------------------------------------------
# Testing
#-----------------------------------------------------------------------------------------------------------------------
.PHONY: test-unit test-acc

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
