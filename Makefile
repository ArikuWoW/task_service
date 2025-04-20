APP_NAME := task-service
CMD_DIR := ./cmd
GO_FILES := $(shell find . -name '*.go' -not -path "./vendor/*")

.PHONY: all build run tidy fmt lint test clean

# Собрать бинарник
build:
	go build -o bin/$(APP_NAME) $(CMD_DIR)

# Запустить приложение
run:
	go run $(CMD_DIR)/main.go

# Установить зависимости и обновить go.mod/go.sum
tidy:
	go mod tidy

# Форматировать весь код
fmt:
	go fmt ./...

# Проверка стиля и линтинг
lint:
	golangci-lint run || echo "golangci-lint not installed (hint: make install-linter)"

# Установить линтер (если не установлен)
install-linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Запустить тесты
test:
	go test ./... -v -cover

# Удалить бинарник
clean:
	rm -rf bin/
