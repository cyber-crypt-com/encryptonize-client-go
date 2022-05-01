# Copyright 2020-2022 CYBERCRYPT
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

##### Help message #####
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <target> \033[36m\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

##### Build targets #####
.PHONY: build
build: ## Build the Encryptonize client library
	go build -v .

.PHONY: lint
lint: ## Lint the codebase
	gofmt -l -w .
	go mod tidy
	golangci-lint run -E gosec,asciicheck,bodyclose,gocyclo,unconvert,gocognit,misspell,revive,whitespace --timeout 5m

##### Test targets #####
.PHONY: tests
tests: build ## Run tests against Encryptonize server
	@make docker-core-test
	@make docker-objects-test

.PHONY: docker-core-test
docker-core-test: docker-core-test-up ## Run EAAS tests
	USER_INFO=$$(docker exec encryptonize-eaas /eaas create-user rcudiom  | tail -n 1) && \
		export E2E_TEST_UID=$$(echo $$USER_INFO | jq -r ".user_id") && \
		export E2E_TEST_PASS=$$(echo $$USER_INFO | jq -r ".password") && \
		go test -v ./... -run ^TestCore && \
		go test -v ./... -run ^TestEncrypt
	@make docker-core-test-down

.PHONY: docker-core-test-up
docker-core-test-up: ## Start docker EAAS test environment
	cd test && \
		docker-compose --profile eaas up -d

.PHONY: docker-core-test-down
docker-core-test-down: ## Stop docker EAAS test environment
	docker-compose --profile eaas -f test/compose.yaml down

.PHONY: docker-objects-test
docker-objects-test: docker-objects-test-up ## Run objects tests
	USER_INFO=$$(docker exec encryptonize-objects /encryptonize-objects create-user rcudiom  | tail -n 1) && \
		export E2E_TEST_UID=$$(echo $$USER_INFO | jq -r ".user_id") && \
		export E2E_TEST_PASS=$$(echo $$USER_INFO | jq -r ".password") && \
		go test -v ./... -run ^TestCore && \
		go test -v ./... -run ^TestObjects
	@make docker-objects-test-down

.PHONY: docker-objects-test-up
docker-objects-test-up: ## Start docker Objects test environment
	cd test && \
		docker-compose --profile objects up -d

.PHONY: docker-objects-test-down
docker-objects-test-down: ## Stop docker Objects test environment
	docker-compose --profile objects -f test/compose.yaml down