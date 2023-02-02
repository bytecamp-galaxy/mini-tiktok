.PHONY: run test kill

run:
	go run ./cmd/api &
	go run ./cmd/comment &
	go run ./cmd/feed &
	go run ./cmd/publish &
	go run ./cmd/user &

test:
	go test ./test/base_api_test.go ./test/common.go -v

kill:
	# `$` will be escaped
	ps -ef | grep -E "(go run ./cmd)|(/tmp/go-build)" | grep -v "grep" | awk '{print $2}' | xargs kill -9
