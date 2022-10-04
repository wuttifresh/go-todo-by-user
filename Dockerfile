FROM golang:latest

WORKDIR /go/src/github.com/wuttifresh/go-todo-by-user/todo

COPY . .

RUN go mod download

RUN go build -o ./main ./todo/main.go

CMD ["./main"]