FROM golang:1.22

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=arm64 go build -v -o ./expensestracker ./cmd/app
EXPOSE 8080
CMD ["./expensestracker"]