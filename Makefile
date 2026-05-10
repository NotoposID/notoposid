.PHONY: dev run-api run-web build-api docker-up docker-down migrate-up migrate-down clean swagger

# Variabel
BINARY_NAME=notopos-api
API_DIR=apps/api
WEB_DIR=apps/web

# Menjalankan seluruh project (Frontend & Backend) menggunakan Turbo
dev:
	pnpm dev

# Menjalankan Backend API saja
run-api:
	cd $(API_DIR) && go run cmd/api/main.go

# Menjalankan Frontend Web saja
run-web:
	pnpm --filter web dev

# Build Backend API
build-api:
	cd $(API_DIR) && go build -o bin/$(BINARY_NAME) cmd/api/main.go

# Infrastruktur (Docker)
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

# Database Migrations (Jika menggunakan alat migrasi kedepannya)
migrate-up:
	@echo "Running migrations up..."
	# Tambahkan perintah migrasi di sini

# Membersihkan file binary
clean:
	rm -rf $(API_DIR)/bin
	rm -rf $(WEB_DIR)/.next

# Swagger Generation
swagger:
	cd $(API_DIR) && swag init -g cmd/api/main.go --parseDependency --parseInternal -o ./docs
