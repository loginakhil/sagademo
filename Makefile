.PHONY: start
start:
	docker-compose up
	# http://localhost:8080/

.PHONY: trigger
trigger:
	go run cmd/trigger/main.go

.PHONY: start-deposit-worker
start-deposit-worker:
	go run cmd/deposit_worker/main.go

.PHONY: start-withdraw-worker
start-withdraw-worker:
	go run cmd/withdraw_worker/main.go
