# .RECIPEPREFIX := $(.RECIPEPREFIX)<space>
TESTCOVERAGE_THRESHOLD=0

ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

.PHONY: init
init:
	command -v deno >/dev/null || curl -fsSL https://deno.land/install.sh | sh
	command -v pre-commit >/dev/null || brew install pre-commit
	command -v make >/dev/null || brew install make
	[ -f .git/hooks/pre-commit ] || pre-commit install

.PHONY: dev
dev:
	deno run --inspect --allow-all dev

.PHONY: container-start
container-start:
	docker compose --file ./ops/docker/compose.yml up --detach

.PHONY: container-rebuild
container-rebuild:
	docker compose --file ./ops/docker/compose.yml up --detach --build

.PHONY: container-restart
container-restart:
	docker compose --file ./ops/docker/compose.yml restart

.PHONY: container-stop
container-stop:
	docker compose --file ./ops/docker/compose.yml stop

.PHONY: container-destroy
container-destroy:
	docker compose --file ./ops/docker/compose.yml down

.PHONY: container-update
container-update:
	docker compose --file ./ops/docker/compose.yml pull

.PHONY: container-dev
container-dev:
	docker compose --file ./ops/docker/compose.yml watch

.PHONY: container-ps
container-ps:
	docker compose --file ./ops/docker/compose.yml ps --all

.PHONY: container-logs
container-logs:
	docker compose --file ./ops/docker/compose.yml logs

.PHONY: container-cli
container-cli:
	docker compose --file ./ops/docker/compose.yml exec web bash

.PHONY: container-push
container-push:
ifdef VERSION
	docker build --platform=linux/amd64 -t acikyazilim.registry.cpln.io/acikio-web:v$(VERSION) .
	docker push acikyazilim.registry.cpln.io/acikio-web:v$(VERSION)
else
	@echo "VERSION is not set"
endif

%:
	@:
