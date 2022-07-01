# Copyright 2022 CYBERCRYPT
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# 	http://www.apache.org/licenses/LICENSE-2.0
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
build: ## Build the client library
	go build -v ./...

.PHONY: lint
lint: ## Lint the codebase
	gofmt -l -w .
	go mod tidy
	golangci-lint run -E gosec,asciicheck,bodyclose,gocyclo,unconvert,gocognit,misspell,revive,whitespace --timeout 5m

##### Test targets #####
.PHONY: tests
tests: build ## Run tests against dockerized servers
	@make docker-generic-test
	@make docker-storage-test
	@make docker-k1-test

.PHONY: docker-generic-test
docker-generic-test: docker-generic-test-up ## Run D1 Generic tests
	USER_INFO=$$(docker exec d1-service-generic /d1-service-generic create-user rcudio  | tail -n 1) && \
		export E2E_TEST_UID=$$(echo $$USER_INFO | jq -r ".user_id") && \
		export E2E_TEST_PASS=$$(echo $$USER_INFO | jq -r ".password") && \
		go test -v ./d1 -count=1 -run ^TestBase && \
		go test -v ./d1 -count=1 -run ^TestGeneric
	@make docker-generic-test-down

.PHONY: docker-generic-test-up
docker-generic-test-up: ## Start docker D1 Generic test environment
	cd test/d1 && \
	docker-compose --profile generic up -d

.PHONY: docker-generic-test-down
docker-generic-test-down: ## Stop docker D1 Generic test environment
	docker-compose --profile generic -f test/d1/compose.yaml down -v

.PHONY: docker-storage-test
docker-storage-test: docker-storage-test-up ## Run D1 Storage tests
	USER_INFO=$$(docker exec d1-service-storage /d1-service-storage create-user rcudio  | tail -n 1) && \
		export E2E_TEST_UID=$$(echo $$USER_INFO | jq -r ".user_id") && \
		export E2E_TEST_PASS=$$(echo $$USER_INFO | jq -r ".password") && \
		go test -v ./d1 -count=1 -run ^TestBase && \
		go test -v ./d1 -count=1 -run ^TestStorage
	@make docker-storage-test-down

.PHONY: docker-storage-test-up
docker-storage-test-up: ## Start docker D1 Storage test environment
	cd test/d1 && \
	docker-compose --profile storage up -d

.PHONY: docker-storage-test-down
docker-storage-test-down: ## Stop docker D1 Storage test environment
	docker-compose --profile storage -f test/d1/compose.yaml down -v

.PHONY: docker-k1-test
docker-k1-test: docker-k1-test-up ## Run Key Server tests
	KS_RESPONSE=$$(docker exec key-server /k1 newKeySet 2> /dev/null) && \
		KS_ID=$$(echo $$KS_RESPONSE | jq -r ".KsID") && \
		KIK_RESPONSE=$$(docker exec key-server /k1 newKik --ksid=$$KS_ID 2> /dev/null) && \
		export E2E_TEST_KIK_ID=$$(echo $$KIK_RESPONSE | jq -r ".KikID") && \
		export E2E_TEST_KIK=$$(echo $$KIK_RESPONSE | jq -r ".Kik") && \
		go test -v ./k1 -count=1
	@make docker-k1-test-down

.PHONY: docker-k1-test-up
docker-k1-test-up: ## Start docker Key Server test environment
	docker-compose -f test/k1/compose.yaml up -d

.PHONY: docker-k1-test-down
docker-k1-test-down: ## Stop docker Key Server test environment
	docker-compose -f test/k1/compose.yaml down -v
