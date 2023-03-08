FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go mod download github.com/bytedance/sonic

RUN go build -o /unit-test

EXPOSE 8080

CMD [ "/unit-test" ]

