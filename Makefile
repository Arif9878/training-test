
generate-grpc:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./question-2/transport/grpc/imdb.proto

run-question-2:
	go run ./question-2/cmd .

run-question-3:
	go run ./question-3 .

run-question-4:
	go run ./question-4 .

test:
	go test -coverprofile fmtcoverage.out ./...

tidy:
	go mod tidy
	go mod vendor