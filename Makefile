.PHONY: run test kill

REPLICA_CONFIG = ./configs/replica.config
include ${REPLICA_CONFIG}

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
	docker compose up -d --scale user-server=${USER_SERVER} \
			--scale feed-server=${FEED_SERVER} \
			--scale publish-server=${PUBLISH_SERVER} \
			--scale relation-server=${RELATION_SERVER} \
			--scale favorite-server=${FAVORITE_SERVER} \
			--scale comment-server=${COMMENT_SERVER}

image:
	docker build -f Dockerfile -t mini-tiktok .
