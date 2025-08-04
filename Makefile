OUTPUT_DIR = bin/app
API_MAIN = cmd/api/main.go
SWAGGER_OUT = docs
.PHONY: all build clean run tidy
.PHONY: swagger

all: build

build:
	@echo "Building the application..."
	go build -o ${OUTPUT_DIR} $(API_MAIN)

run: build
	@echo "Running the application..."
	${OUTPUT_DIR}

clean:
	@echo "Cleaning up..."
	rm -rf ${OUTPUT_DIR}

tidy:
	@echo "Tidying up dependencies..."
	go mod tidy

migration:
	@if [ -z "$(name)" ]; then \
		echo "err: migration name is required name=<migration_name>"; \
		exit 1; \
	fi
	@echo "Creating migration: $(name)..."
	migrate create -ext sql -dir migrations -seq $(name)

migrate:
	@if [ -z "$(direction)" ]; then \
		echo "err: direction is required (direction=up|down)"; \
		exit 1; \
	fi; \
	if [ "$(direction)" != "up" ] && [ "$(direction)" != "down" ]; then \
		echo "err: invalid direction value (use up or down)"; \
		exit 1; \
	fi
	@echo "Running migrations $(direction)..."
	go run migrations/auto.go $(direction)

swagger:
	@echo "Generating Swagger docs..."
	swag init -g $(API_MAIN) -d .,./docs/response -o $(SWAGGER_OUT)
	@echo "Cleaning up unwanted lines in Swagger docs..."
	@sed -i.bak '/LeftDelim/d' $(SWAGGER_OUT)/docs.go && rm -f $(SWAGGER_OUT)/docs.go.bak
	@sed -i.bak '/RightDelim/d' $(SWAGGER_OUT)/docs.go && rm -f $(SWAGGER_OUT)/docs.go.bak

swagger-clean:
	@echo "Cleaning up Swagger docs..."
	rm -f $(SWAGGER_OUT)/docs.go $(SWAGGER_OUT)/swagger.json $(SWAGGER_OUT)/swagger.yaml

swagger-fmt:
	@echo "Formatting Swagger docs..."
	swag fmt