GO111MODULE	= on
export BUILD = CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s -w'

MODULE = $(shell go list -m)
MODULE_NAME = $(lastword $(subst /, ,$(MODULE)))
SERVICES = $(shell ls -d {api,fnc,srv,web}/*/ 2> /dev/null)
DIR = $(shell pwd)

.PHONY: lint build run api fnc srv web default

default: build run

mod_check:
	@if [ -f 'go.mod' ]; \
	then \
		exit 0; \
	else \
		echo "Go 1.11 modules required: use 'go mod init'"; \
		echo "More about modules: 'go help moodules'"; \
		exit 1; \
	fi

stat:
	@echo "Module: " $(MODULE) ", List of services: " $(SERVICES)

lint: stat
	@for s in $(SERVICES); do \
		echo "Verifying $(DIR)/$$s"; \
		pushd $(DIR)/$$s >/dev/null; \
		go fmt; \
		golint; \
		popd >/dev/null; \
	done

test:
	@for s in $(SERVICES); do \
		echo "Building $(DIR)/$$s"; \
		$(MAKE) -C $$s test; \
	done

build:
	@for s in $(SERVICES); do \
		echo "Building $(DIR)/$$s"; \
		$(MAKE) -C $$s build; \
	done

run:
	docker-compose up --build

api: mod_check
	@while [ -z "$$NAME" ]; do \
		read -r -p "API service name: " NAME; \
	done ; \
	micro new $(MODULE)/api/$$NAME --namespace=$(MODULE_NAME) --type=api

fnc: mod_check
	@while [ -z "$$NAME" ]; do \
		read -r -p "Function service name: " NAME; \
	done ; \
	micro new $(MODULE)/fnc/$$NAME --namespace=$(MODULE_NAME) --type=fnc

srv: mod_check
	@while [ -z "$$NAME" ]; do \
		read -r -p "Service name: " NAME; \
	done ; \
	micro new $(MODULE)/srv/$$NAME --namespace=$(MODULE_NAME) --type=srv

web: mod_check	
	@while [ -z "$$NAME" ]; do \
		read -r -p "Web service name: " NAME; \
	done ; \
	micro new $(MODULE)/web/$$NAME --namespace=$(MODULE_NAME) --type=web
	
user-over-api:
	@curl 'http://localhost:8082/rpc' \
		-H 'Content-type: application/json' \
		-d '{"service":"go.micro.srv.user","endpoint":"UserService.Create","request":{"name":"Ewan Valentine","email":"ewan.valentine'$$RANDOM'@gmail.com","password":"test123"}}'

user-over-grpc-client:
	docker-compose exec user-service /user-srv --run_client
