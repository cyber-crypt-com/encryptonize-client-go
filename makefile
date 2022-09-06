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

# Check that given variables are set and all have non-empty values,
# die with an error otherwise.
#
# Params:
#   1. Variable name(s) to test.
#   2. (optional) Error message to print.
check_defined = \
    $(strip $(foreach 1,$1, \
        $(call __check_defined,$1,$(strip $(value 2)))))
__check_defined = \
    $(if $(value $1),, \
      $(error Undefined $1$(if $2, ($2))))

##### Build targets #####
.PHONY: build
build: ## Build the client library
	go build -v ./...

.PHONY: lint
lint: ## Lint the codebase
	gofmt -l -w .
	go mod tidy
	golangci-lint run -E gosec,asciicheck,bodyclose,gocyclo,unconvert,gocognit,misspell,revive,whitespace --timeout 5m

##### Copy targets #####
.PHONY: copy-generic-client
copy-generic-client: ## Copy D1 Generic client source code into this repo
	$(call check_defined, VERSION, Usage: make copy-generic-client VERSION=<version>)
	./scripts/copy-client.sh generic ${VERSION}

.PHONY: copy-storage-client
copy-storage-client: ## Copy D1 Storage client source code into this repo
	$(call check_defined, VERSION, Usage: make copy-storage-client VERSION=<version>)
	./scripts/copy-client.sh storage ${VERSION}

.PHONY: copy-k1-client
copy-k1-client: ## Copy K1 client source code into this repo
	$(call check_defined, VERSION, Usage: make copy-k1-client VERSION=<version>)
	./scripts/copy-client.sh k1 ${VERSION}
