FROM golang:1.22

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -v -o ./expensestracker ./cmd/app GOOS=linux GOARCH=ard64
EXPOSE 8080
CMD ["./expensestracker"]