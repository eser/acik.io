# .RECIPEPREFIX := $(.RECIPEPREFIX)<space>
TESTCOVERAGE_THRESHOLD=0

.PHONY: init
init:
	brew install pre-commit
	brew install make
	brew install protobuf
	pre-commit install
	go install github.com/air-verse/air@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/jandelgado/gcov2lcov@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: dev
dev:
	air; if [ $$? -ne 0 ]; then go run ./cmd/acik-service/; fi

.PHONY: build
build:
	go build -o ./tmp/dist/acik-service ./cmd/acik-service/

.PHONY: generate
generate:
	go generate ./...

.PHONY: clean
clean:
	go clean

.PHONY: run
run: build
	./tmp/dist/acik-service

.PHONY: test-api
test-api:
	cd ./deployments/api/ && \
	bru run ./ --env development && \
	cd ../../

.PHONY: test
test:
	go test -failfast -count 1 ./...

.PHONY: test-cov
test-cov:
	go test -failfast -count 1 -coverpkg=./... -coverprofile=${TMPDIR}cov_profile.out ./...
	# `go env GOPATH`/bin/gcov2lcov -infile ${TMPDIR}cov_profile.out -outfile ./cov_profile.lcov

.PHONY: test-view-html
test-view-html:
	go tool cover -html ${TMPDIR}cov_profile.out -o ${TMPDIR}cov_profile.html
	open ${TMPDIR}cov_profile.html

.PHONY: test-ci
test-ci: test-cov
	$(eval ACTUAL_COVERAGE := $(shell go tool cover -func=${TMPDIR}cov_profile.out | grep total | grep -Eo '[0-9]+\.[0-9]+'))

	@echo "Quality Gate: checking test coverage is above threshold..."
	@echo "Threshold             : $(TESTCOVERAGE_THRESHOLD) %"
	@echo "Current test coverage : $(ACTUAL_COVERAGE) %"

	@if [ "$(shell echo "$(ACTUAL_COVERAGE) < $(TESTCOVERAGE_THRESHOLD)" | bc -l)" -eq 1 ]; then \
    echo "Current test coverage is below threshold. Please add more unit tests or adjust threshold to a lower value."; \
    echo "Failed"; \
    exit 1; \
  else \
    echo "OK"; \
  fi

.PHONY: dep
dep:
	go mod download
	go mod tidy

.PHONY: lint
lint:
	`go env GOPATH`/bin/golangci-lint run

.PHONY: container-start
container-start:
	docker compose --file ./deployments/compose.yml up --detach

.PHONY: container-rebuild
container-rebuild:
	docker compose --file ./deployments/compose.yml up --detach --build

.PHONY: container-restart
container-restart:
	docker compose --file ./deployments/compose.yml restart

.PHONY: container-stop
container-stop:
	docker compose --file ./deployments/compose.yml stop

.PHONY: container-destroy
container-destroy:
	docker compose --file ./deployments/compose.yml down

.PHONY: container-update
container-update:
	docker compose --file ./deployments/compose.yml pull

.PHONY: container-dev
container-dev:
	docker compose --file ./deployments/compose.yml watch

.PHONY: container-ps
container-ps:
	docker compose --file ./deployments/compose.yml ps --all

.PHONY: container-logs-all
container-logs-all:
	docker compose --file ./deployments/compose.yml logs

.PHONY: container-logs
container-logs:
	docker compose --file ./deployments/compose.yml logs acik-service

.PHONY: container-cli
container-cli:
	docker compose --file ./deployments/compose.yml exec acik-service bash

.PHONY: container-push
container-push:
ifdef VERSION
	docker build --platform=linux/amd64 -t acikyazilim.registry.cpln.io/acik-service:v$(VERSION) .
	docker push acikyazilim.registry.cpln.io/acik-service:v$(VERSION)
else
	@echo "VERSION is not set"
endif

.PHONY: generate-proto
# --ts_proto_opt="context=true,env=node,lowerCaseServiceMethods=true,outputServices=grpc-js,removeEnumPrefix=true,snakeToCamel=true,useAbortSignal=true,useAsyncIterable=true,useReadonlyTypes=true,comments=false,useNullAsOptional=true"
generate-proto:
	@{ \
	  for f in ./proto/*; do \
	    current_proto="$$(basename $$f)"; \
	    echo "Generating stubs for $$current_proto"; \
			\
			protoc --plugin=./web/node_modules/.bin/protoc-gen-ts_proto \
				--proto_path=./proto/ \
				--ts_proto_out=./web/proto/ \
				--ts_proto_opt="context=true,lowerCaseServiceMethods=true,outputServices=grpc-js,removeEnumPrefix=false,snakeToCamel=true,useReadonlyTypes=true,comments=false,useNullAsOptional=true" \
				"./proto/$$current_proto/$$current_proto.proto"; \
			protoc \
				--proto_path=./proto/ \
				--go_out=./pkg/proto/ --go_opt=paths=source_relative \
				--go-grpc_out=./pkg/proto/ --go-grpc_opt=paths=source_relative \
				"./proto/$$current_proto/$$current_proto.proto"; \
	  done \
	}
