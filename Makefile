.PHONY: dep-install lint vet test coverage build clean deploy

# GENERAL VARS
LLM-TEST-LAMBDA_FILE=llm-test-v0.0.1.zip

# Dep install
dep-install:
	npm install serverless
	serverless plugin install -n serverless-go-plugin

# LINT / VET / TEST
lint:
	@echo "Linting ..."
	@golangci-lint run --timeout=10m

vet:
	@echo "Vetting ..."
	@go vet ./...

test:
	@echo "Running unit tests..."
	@go clean -testcache go test ./... -cover -v -coverprofile=cover.out

coverage: test
	@go tool cover -html=cover.out -o cover.html

#########################################
# LAMBDA ACTIONS #
#########################################

clean:
	rm -rf ./bin ./vendor Gopkg.lock

remove:
	sls remove

deploy: clean
	sls deploy --verbose