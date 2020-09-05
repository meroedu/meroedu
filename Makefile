BINARY=meroedu
DB_CONFIG_FILE ?= ./migrator/dbconf.yml
DB_DSN ?= $(shell sed -n 's/^dsn:[[:space:]]*"\(.*\)"/\1/p' $(DB_CONFIG_FILE))
##############################################################################
# Staging
##############################################################################

run:
	docker-compose -f docker-compose.yaml up --build -d
stop:
	docker-compose -f docker-compose.yaml down
##############################################################################
# Development
##############################################################################

run-dev:
	docker-compose -f docker-compose.dev.yaml up --build
stop-dev:
	docker-compose -f docker-compose.dev.yaml down

##############################################################################
# Lint
###############################################################################
lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

##############################################################################
# Test
###############################################################################

test-richgo: 
	richgo test -v -cover -covermode=atomic ./...

test: 
	go test -v -cover -covermode=atomic ./...

unittest:
	go test -short  ./...

test-coverage:
	mkdir -p ./out
	go test -coverprofile=./out/coverage.out ./...
	go tool cover -func=./out/coverage.out

sonar: test
	sonar-scanner -Dsonar.projectVersion="$(version)"

start-sonar:
	docker run --name sonarqube -p 9000:9000 sonarqube
#############################################################################
# Migration
#############################################################################
migrate-build:
	cd migrator/ && docker build -t migrator .

migrate: migrate-build
	docker run --network host migrator -path=/migrations/ -database "$(DB_DSN)" up

migrate-down:
	docker run --network host migrator -path=/migrations/ -database "$(DB_DSN)" down -all
	
#############################################################################
# Utility
#############################################################################
db-diagram:
	java -jar ~/Downloads/schema-gui/schemaspy-6.1.0.jar -dp ~/Downloads/mysql-connector-java-6.0.6.jar -t mysql -db course_api -host localhost -u root -p "root" -o ~/Downloads/schema-gui/course_api -s course_api
build-app: clean-app
	go build -o ${BINARY}

run-app: build-app
	./${BINARY}
clean-app:
	$(eval VALUE=$(shell sh -c "lsof -i:9090 -t"))
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	$(shell sh -c "if [ \"${VALUE}\" != \"\" ]  ; then kill ${VALUE} ; fi")
docker:
	docker build -t course_api .

swagger:
	go get github.com/swaggo/swag/cmd/swag
	$$(go env GOPATH)/bin/swag init -g meroedu.go --output ./api_docs
mock:
	cd internal/domain && mockery --all --keeptree
db-up:
	docker-compose up -d mysql
.PHONY: clean install unittest build docker run stop vendor lint-prepare lint