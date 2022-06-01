# System Setup
SHELL = bash

GRAPHQL_SCHEMA_VERSION ?= 0.0.5

# Go Stuff
CGO_ENABLED ?= 0

# General Vars
COVERAGE_DIR=coverage

.PHONY: all
all: test build

.PHONY: travis_setup
travis_setup: ## Setup the travis environmnet
	@if [[ -n "$$BUILD_ENV" ]] && [[ "$$BUILD_ENV" == "testing" ]]; then echo -e "$(INFO_COLOR)THIS IS EXECUTING AGAINST THE TESTING ENVIRONMEMNT$(NO_COLOR)"; fi
	@echo "Downloading latest Ionize"
	@wget --quiet https://s3.amazonaws.com/public.ionchannel.io/files/ionize/linux/bin/ionize
	@chmod +x ionize && mkdir -p $$HOME/.local/bin && mv ionize $$HOME/.local/bin
	@echo "Installing Go Linter"
	@go get -u golang.org/x/lint/golint

.PHONY: analyze
analyze:  ## Perform an analysis of the project
	@if [[ -n "$$BUILD_ENV" ]] && [[ "$$BUILD_ENV" == "testing" ]]; then \
		IONCHANNEL_SECRET_KEY=$$TESTING_APIKEY IONCHANNEL_ENDPOINT_URL=$$TESTING_ENDPOINT_URL ionize --config .ionize.test.yaml analyze; \
	else \
		ionize analyze; \
	fi

.PHONY: clean
clean: ## Cleans out all generated items
	-@go clean
	-@rm -rf coverage

.PHONY: coverage
coverage:  ## Generates the code coverage from all the tests
	@echo "Total Coverage: $$(make --no-print-directory coverage_compfriendly | tee coverage.txt)%"

.PHONY: coverage_compfriendly
coverage_compfriendly:  ## Generates the code coverage in a computer friendly manner
	-@rm -rf coverage
	-@mkdir -p $(COVERAGE_DIR)/tmp
	@for j in $$(go list ./... | grep -v '/vendor/' | grep -v '/ext/'); do go test -covermode=count -coverprofile=$(COVERAGE_DIR)/$$(basename $$j).out $$j > /dev/null 2>&1; done
	@echo 'mode: count' > $(COVERAGE_DIR)/tmp/full.out
	@tail -q -n +2 $(COVERAGE_DIR)/*.out >> $(COVERAGE_DIR)/tmp/full.out
	@go tool cover -func=$(COVERAGE_DIR)/tmp/full.out | tail -n 1 | sed -e 's/^.*statements)[[:space:]]*//' -e 's/%//'

.PHONY: help
help:  ## Show This Help
	@for line in $$(cat Makefile | grep "##" | grep -v "grep" | sed  "s/:.*##/:/g" | sed "s/\ /!/g"); do verb=$$(echo $$line | cut -d ":" -f 1); desc=$$(echo $$line | cut -d ":" -f 2 | sed "s/!/\ /g"); printf "%-30s--%s\n" "$$verb" "$$desc"; done

.PHONY: test
test:  ## Run all available tests
	@go test -v ./...

.PHONY: linters
linters:  fmt vet  ## Run all of the linters

.PHONY: fmt
fmt: ## Run go fmt
	@echo "checking formatting..."
	@go fmt ./...

.PHONY: vet
vet: ## Run go vet
	@echo "vetting..."
	@go vet ./...

.PHONY: docs
docs: ## exports documents from the source code
	@echo "creating documentation..."
	./make-public-docs > docs/endpoints.md
	@pandoc -f markdown docs/endpoints.md -c api.css -s --highlight-style monochrome --metadata pagetitle="API Documentation" -o docs/index.html
	@pandoc -f markdown docs/examples.md -c api.css -s --highlight-style monochrome --metadata pagetitle="API Examples" -o docs/examples.html
	@pandoc -f markdown docs/data_examples.md -c api.css -s --highlight-style monochrome --metadata pagetitle="API Data Examples" -o docs/data_examples.html
	@pandoc -f markdown docs/gitlab_examples.md -c api.css -s --highlight-style monochrome --metadata pagetitle="API Gitlab Examples" -o docs/gitlab_examples.html

.PHONY: deploy_docs
deploy_docs: ## deploys the docs to S3
	aws s3 sync ./docs/ s3://docs.ionchannel.io

.PHONY: get_schema
get_schema:
	rm schema.graphqls
	curl -L -s https://github.com/ion-channel/graphql-schema/releases/download/$(GRAPHQL_SCHEMA_VERSION)/consolidated_schema.graphqls \
		-o consolidated_schema.graphqls
	echo "## Schema version $(GRAPHQL_SCHEMA_VERSION)" > schema.graphqls
	cat consolidated_schema.graphqls >> schema.graphqls
	rm -f consolidated_schema.graphqls

.PHONY: generate
generate:
	go run github.com/99designs/gqlgen generate
