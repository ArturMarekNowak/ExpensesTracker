FROM golang:1.22

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build GOOS=linux GOARCH=ard64 -v -o ./expensestracker ./cmd/app
EXPOSE 8080
CMD ["./expensestracker"]