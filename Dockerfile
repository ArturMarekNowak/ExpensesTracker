FROM golang:1.22

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -C cmd/app -o /expensestracker
EXPOSE 8080

CMD ["/expensestracker"]