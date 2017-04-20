.PHONY: swagger run

all: swagger

run:
	go run main.go

swagger:
	go get github.com/go-swagger/go-swagger/cmd/swagger
	swagger generate spec -o swagger.json