FROM golang:1.16 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -o user-management-service

FROM alpine:3.14
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/user-management-service /user-management-service

CMD ["/user-management-service"]
