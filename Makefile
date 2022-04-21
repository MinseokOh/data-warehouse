#!/usr/bin/make -f

BUILDDIR ?= $(CURDIR)/build
KAFKADIR ?= $(CURDIR)/kafka
CLI ?= data-warehouse

##########################################
#				protobuf	 			 #
##########################################
protocgen:
	@echo "===> generate .proto files..."
	@bash ./scripts/protocgen.sh

##########################################
#				build		 			 #
##########################################
build: clean
	@echo "===> build project..."
	@go build -o $(BUILDDIR)/$(CLI) ./cmd/main.go

clean:
	@echo "===> clean build files..."
	@find ./build -maxdepth 2 -name '*.log' | xargs rm -f >/dev/null || true
	@mkdir -p $(BUILDDIR)/transformer
	@mkdir -p $(BUILDDIR)/producer

##########################################
#				producer 	 			 #
##########################################
producer-start:
	@echo "===> running producers..."
	@bash ./scripts/start_producers.sh

producer-stop:
	@echo "===> stopping producers..."
	@bash ./scripts/stop_producers.sh 2>/dev/null; true


##########################################
#				transformer	 			 #
##########################################

# TODO check all process stopped
transformer-start:
	@echo "===> running transformers..."
	@sleep 1
	@bash ./scripts/start_transformer.sh

# TODO check all process stopped
transformer-stop:
	@echo "===> stopping transformers..."
	@sleep 1
	@bash ./scripts/stop_transformer.sh 2>/dev/null; true

##########################################
#				testing 	 			 #
##########################################
start-test:docker-compose-down docker-compose build transformer-start producer-start multitail

stop-test:transformer-stop producer-stop docker-compose-down

##########################################
#				test		 			 #
##########################################
test: test-unit test-benchmark

test-unit: test-schemas test-types

test-types:
	@echo "===> test types..."
	@go test -v ./types/...

test-schemas:
	@echo "===> test omni schemas..."
	@go test -v ./schema/

test-benchmark:
	@echo "===> test benchmark..."
	@go test -v -bench=. -benchmem ./bench/

##########################################
#			docker-compose				 #
##########################################
install-docker-compose:
	@sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

#TODO check kafka cluster running
docker-compose:
	@echo "===> docker-compose up -d ..."
	@docker-compose up -d
	@echo "===> wait for containers healthy..."
	@bash ./scripts/wait-for-healthy.sh mysql

docker-compose-stop:
	@echo "===> docker-compose stop ..."
	@docker-compose stop

docker-compose-down:
	@docker-compose down

##########################################
#			multitail					 #
##########################################
multitail:
	@bash ./scripts/multitail.sh