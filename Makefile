.PHONY: test create-challenge

test:
	go test ./challenges/...

challenge:
	go run main.go -year $(year) -day $(day)

create-challenge:
	go run ./create-challenge/create-challenge.go -year $(year) -day $(day)