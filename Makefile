
.PHONY: all
all: checks unit-test

.PHONY: checks
checks: lint

.PHONY: lint
lint:
	@scripts/check_lint.sh

.PHONY: unit-test
unit-test:
	@scripts/check_unit.sh
