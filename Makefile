.PHONY: dev
dev:
	docker-compose -f  tools/docker-compose.dev.yml up -d
	REDIS_ADDR=localhost:6379 \
	JWT_KEY=zUxMiIsInR5c \
	go run main.go

.PHONY: stop
stop:
	docker-compose -f  tools/docker-compose.dev.yml down
