FROM golang:1.12-alpine3.9 as base
WORKDIR /tmp/vk-auth-service
COPY . .
RUN go build -mod vendor -o /tmp/service .

FROM alpine:3.9
WORKDIR /tmp
COPY --from=base /tmp/service ./service
ENTRYPOINT ["./service"]
EXPOSE 80
