# If the first argument is "run"
ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

# If the first argument is "dev"
ifeq (dev,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "dev"
  DEV_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(DEV_ARGS):;@:)
endif

build:
	@go build -o bin/genomdb

run: build
	@./bin/genomdb $(RUN_ARGS)

dev:
	@go run . $(DEV_ARGS)

test:
	@go test -v ./...
