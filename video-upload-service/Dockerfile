FROM golang:1.16 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -o video-upload-service

FROM alpine:3.14
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/video-upload-service /video-upload-service

CMD ["/video-upload-service"]
