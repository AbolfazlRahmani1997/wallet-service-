FROM golang:latest AS build
ENV GO111MODULE=on
WORKDIR /build
COPY . .
WORKDIR ./cmd
RUN CGO_ENABLED=0 GOOS=linux go build ./main.go

# Final stage
FROM alpine:latest

WORKDIR /app
COPY --from=build /build/cmd/main .
RUN chmod +x ./main
EXPOSE 9000
CMD ["./main"]