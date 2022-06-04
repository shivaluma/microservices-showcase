FROM golang:1.18.3-alpine as builder


WORKDIR /build

COPY go.mod .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./main
RUN ["chmod", "777", "./main"]


FROM alpine:latest
COPY --from=builder /build/main ./app
RUN ["chmod", "+x", "./app"]

EXPOSE 1001

ENTRYPOINT ["./app"]


