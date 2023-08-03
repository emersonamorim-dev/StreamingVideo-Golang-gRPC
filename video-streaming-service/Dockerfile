FROM golang:1.16 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -o video-streaming-service

FROM alpine:3.14
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/video-streaming-service /video-streaming-service

CMD ["/video-streaming-service"]
