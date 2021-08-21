
run-question-3:
	go run ./question-3 .

run-question-4:
	go run ./question-4 .

test:
	go test -coverprofile fmtcoverage.out ./...

tidy:
	go mod tidy
	go mod vendor