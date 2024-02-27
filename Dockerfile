FROM golang:1.20-alpine as builder

WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./src ./src

RUN go build -o ./app ./src

FROM scratch
COPY --from=builder /app/app /app

ENTRYPOINT ["/app"]
