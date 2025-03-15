FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin ./cmd/main.go

FROM scratch AS final

WORKDIR /goloom

COPY --from=builder /build/bin /bin

CMD ["/bin"]