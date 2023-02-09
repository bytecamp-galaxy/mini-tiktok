.PHONY: run test kill

run:
	mkdir -p logs
	go run ./cmd/api &
	go run ./cmd/comment &
	go run ./cmd/feed &
	go run ./cmd/publish &
	go run ./cmd/user &
	go run ./cmd/favorite &
	go run ./cmd/relation &

test:
	go test ./test/api_test.go ./test/common.go -v

kill:
	# `$` will be escaped
	ps -ef | grep -E "(go run ./cmd)|(/tmp/go-build)" | grep -v "grep" | awk '{print $$2}' | xargs kill -9

service:
	docker compose up mysql redis etcd otel-collector jaeger-all-in-one victoriametrics grafana minio -d

docker:
	mkdir -p logs
	docker compose up

image:
	docker build -f Dockerfile -t mini-tiktok .
