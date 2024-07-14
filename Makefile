build:
	cd cmd/ && go build -o gendiff
test:
	go test ./tests/...