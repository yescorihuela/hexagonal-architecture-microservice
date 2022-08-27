.PHONY: test
test:
	go clean -testcache && go test ./...

.PHONY: dcu
dcu:
	docker compose up --build

.PHONY: dcd
dcd:
	docker compose down
