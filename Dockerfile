FROM golang:latest

WORKDIR /go/src/github.com/wuttifresh/go-todo-by-user/todo

COPY . .

RUN go mod download

RUN go build ./todo/main.go

CMD ["/example/todolist"]