.PHONY: test create-challenge

test:
	go test ./challenges/...

create-challenge:
	go run ./create-challenge/create-challenge.go -year $(year) -day $(day)