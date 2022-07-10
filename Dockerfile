FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/banking-api /app/cmd/banking-api/main.go

EXPOSE 8080

CMD [ "./out/banking-api" ]