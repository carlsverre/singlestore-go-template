ALL: run

run: build
	./main

build:
	go build main.go

.PHONY: init-singlestore
init-singlestore: check-env
	@docker inspect singlestore-db >/dev/null 2>&1 \
	|| docker run -it \
		--name singlestore-db \
		-e LICENSE_KEY=$(SINGLESTORE_LICENSE) \
		-p 3306:3306 -p 8080:8080 \
		memsql/cluster-in-a-box \
	|| docker rm singlestore-db

start-singlestore: init-singlestore
	@docker start singlestore-db >/dev/null \
	&& echo "SingleStore DB started"

stop-singlestore:
	@docker stop singlestore-db >/dev/null \
	&& echo "SingleStore DB stopped"

check-env:
ifndef SINGLESTORE_LICENSE
	$(error SINGLESTORE_LICENSE is undefined; you can get a free license here: https://portal.singlestore.com/)
endif

# .PHONY tells Makefile to ignore files named after these targets, will ensure
# that these targets always execute.
.PHONY: run
.PHONY: build
.PHONY: init-singlestore
.PHONY: start-singlestore
.PHONY: stop-singlestore
.PHONY: check-env
