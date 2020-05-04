CURR_DIR = $(shell pwd)
MAIN_PATH = $(CURR_DIR)/api/main.go
PROJECT_NAME = "TODO"

init:
	cd $(CURR_DIR)/ui && npm install
	cd $(CURR_DIR)/api && go mod tidy

test:
	cd $(CURR_DIR)/api && go test ./...
	cd $(CURR_DIR)/ui && ng test

build:
	cd $(CURR_DIR)/ui && ng build --prod
	cd $(CURR_DIR) && docker build -t todo .

run:
	@docker run -d -p 8080:8080 todo
